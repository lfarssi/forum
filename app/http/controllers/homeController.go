package controllers

import (
	"database/sql"
	"fmt"
	"forum/app/models"
	"net/http"
)

func HomeController(w http.ResponseWriter, r *http.Request,db *sql.DB) {

	// if err := CheckSession(w, r, db); err != nil {
	// 	fmt.Println("error checking session: ", err)
	// 	return 
	// }

	categories, err := fetchCategories(w, r,db )
	if err != nil {
		fmt.Println("error fetching categories: ", err)
		return
	}

	posts, err := fetchPosts(w, r,db )
	if err != nil {
		fmt.Println("error fetching posts: ", err)
		return
	}

	data := struct {
		Posts      []models.Post
		Categories []models.Category
	}{
		Posts:      posts,
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
