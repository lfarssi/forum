package controllers

import (
	"database/sql"
	"fmt"
	"forum/app/models"
	"html"
	"net/http"
)

func PostController(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method == "GET" && r.URL.Path == "/add-post" {
		// if err := CheckSession(w, r, db); err != nil {
		// 		fmt.Println("error checking session: ", err)
		// 	return 
		// }
		categories, err := fetchCategories(w, r, db)
		if err != nil {
			fmt.Println("error fetching categories: ", err)
			return
		}
		data := struct {
			Categories []models.Category
			Error      string 
		}{
			Categories: categories,
			Error: "",
		}
			// fmt.Println("categories : ", data)
		ParseFileController(w, r, "user/addPost", data)
		return
	}else if r.Method == "POST" && r.URL.Path == "/submit-add-post"{
		// if err := CheckSession(w, r, db); err != nil {
		// 	fmt.Println("error checking session: ", err)
		// 	ErrorController(w,r,http.StatusForbidden)
		// 	return 
		// }
			var errors []string
		err:=r.ParseForm()
		if err!=nil{
            fmt.Println("error parsing form : ", err)
            ErrorController(w, r, http.StatusInternalServerError)
            return
        }
		title := html.EscapeString(r.FormValue("title"))
		content := html.EscapeString(r.FormValue("content"))
		categories := r.Form["categories"]
		if title == "" {
			errors = append(errors, "Title is required.")
		}
		if content == "" {
			errors = append(errors, "Content is required.")
		}

		fmt.Println("title : ", title)
		fmt.Println("content : ", content)
		fmt.Println("categories : ", categories)
		
		if len(errors) > 0 {
			categories, err := fetchCategories(w, r, db)
			if err != nil {
				fmt.Println("error fetching categories: ", err)
				return
			}
			data := struct {
				Categories []models.Category
				Error      string 
			}{
				Categories: categories,
				Error:      fmt.Sprintf("%v", errors), 
			}
			ParseFileController(w, r, "user/addPost", data)
			return
		}

		err1:=AddPostController(w,r,db,title,content,categories,1)
		if err1!=nil{ 
			fmt.Println("error adding post : ", err1)
            ErrorController(w, r, http.StatusInternalServerError)
            return
		}
		fmt.Println("redirect to home ")
			http.Redirect(w ,r,"/add-post",http.StatusSeeOther)
			return
		
	}

}
