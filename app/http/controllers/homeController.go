package controllers

import (
	"database/sql"
	"fmt"
	"forum/app/models"
	"net/http"
	"time"
)

func HomeController(w http.ResponseWriter, r *http.Request) {

	db, err := sql.Open("sqlite3", "/database/database.db")
	if err != nil {
		fmt.Println("error opening database: ", err)
		ErrorController(w, r, http.StatusInternalServerError)
		return
	}
	defer func() {
		err2 := db.Close()
		if err2 != nil {
			fmt.Println("error closing database: ", err2)
			ErrorController(w, r, http.StatusInternalServerError)

		}
	}()
	cookie, err := r.Cookie("session_token")
	if err != nil || cookie.Value == "" {
		var UserId int
		err := db.QueryRow("select user_id from sessionss where token=? and expired_at > ? ", cookie.Value, time.Now()).Scan(&UserId)
		if err == nil {
			http.Redirect(w, r, "/home", http.StatusFound)
			return
		}
	}

	categories := []models.Category{}
	rows,err:=db.Query("select id,name from categories")
	if err!=nil {
        fmt.Println("error querying categories: ", err)
        ErrorController(w, r, http.StatusInternalServerError)
        return
    }
	defer rows.Close()

	for rows.Next(){
		var category models.Category
       if err3:=rows.Scan(&category.ID,&category.Name);err3!=nil{
		    fmt.Println("error scanning categories: ", err3)
			ErrorController(w, r, http.StatusInternalServerError)
			return
	   }
	   categories = append(categories, category)
	}
	posts :=[]models.Post{}
	query:= `
	selct 
		p.id
		p.title
		p.content
		c.name AS category
		coalesce(sum(case when pr.react_type='like' then 1 else 0 end ),0)as likes
		coalesce(sum(case when pr.react_type='dislike' then 1 else 0 end ),0)as dislikes
		p.created_at
		from posts p
		left join post_categories pc on p.id =pc.post_id
		left join categories c on pc.category_id = c.id
		left join reactPost pr on p.id = pr.post_id
		left join users u on p.user_id=u.id
		group by p.id
		order by p.created_at DESC
	`


	rows2,err:=db.Query(query)
	if err!=nil {
        fmt.Println("error querying posts: ", err)
        ErrorController(w, r, http.StatusInternalServerError)
        return
    }
	defer rows2.Close()
	for rows2.Next(){
		var post models.Post
		if err:=rows2.Scan(&post.ID,&post.Title,&post.Category,&post.Username,&post.Likes,&post.Dislikes,&post.CreatedAt);err!=nil{
			fmt.Println("error scanning posts: ", err)
            ErrorController(w, r, http.StatusInternalServerError)
            return
		}
	

	comments := []models.Comment{}
	commentQuery:=`
		slect 
		id,
		content,
		created_at ,
		user_id,
		(select username form users where id=user_id)as username
		from comments 
		where post_id =?
		order by created_at asc
	`

	commentRow,err:=db.Query(commentQuery,post.ID)
	if err!=nil {
        fmt.Println("error querying comments: ", err)
        ErrorController(w, r, http.StatusInternalServerError)
        return
    }
	defer commentRow.Close()
	for commentRow.Next(){
		var comment models.Comment
		if err:=commentRow.Scan(&comment.ID,&comment.Content,&comment.CreatedAt,&comment.UserID,&comment.Username);err!=nil{
            fmt.Println("error scanning comments: ", err)
            ErrorController(w, r, http.StatusInternalServerError)
            return
        }
		var username string 
		if err := db.QueryRow("SELECT username FROM users WHERE id = ?", comment.UserID).Scan(&username); err == nil {
			comment.Username = username // Set the username for each comment
		}
		comments=append(comments, comment)

	}
	
	post.Comments = comments
	posts = append(posts, post)
}

data := struct {
	Posts     []models.Post
	Categories []models.Category
}{
	Posts:     posts,
	Categories: categories,
}
	if r.URL.Path == "/" {
		ParseFileController(w, r, "guest/index", data)
		return
	} else if r.URL.Path == "/home" {
		ParseFileController(w, r, "user/index", data)
		return
	} else {
		ErrorController(w, r, http.StatusNotFound)
		return
	}
}
