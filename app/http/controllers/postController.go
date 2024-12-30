package controllers

import (
	"database/sql"
	"fmt"
	"forum/app/models"
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
			}{
				Categories: categories,
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
		err:=r.ParseForm()
		if err!=nil{
            fmt.Println("error parsing form : ", err)
            ErrorController(w, r, http.StatusInternalServerError)
            return
        }
		title := r.FormValue("title")
		content := r.FormValue("content")
		categories := r.Form["categories"]
		fmt.Println("title : ", title)
		fmt.Println("content : ", content)
		fmt.Println("categories : ", categories)
		err1:=AddPostController(w,r,db,title,content,categories,1)
		if err1!=nil{ 
			fmt.Println("error adding post : ", err)
            ErrorController(w, r, http.StatusInternalServerError)
            return
		}
		http.Redirect(w ,r,"/home",http.StatusCreated)
		return
	}

}
