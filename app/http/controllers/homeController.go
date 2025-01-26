package controllers

import (
	"fmt"
	"net/http"

	"forum/app/models"
	"forum/utils"
)

func HomeController(w http.ResponseWriter, r *http.Request) {
	var logedIn bool
	comments, err := models.GetComments(1)
	if err != nil {
		ErrorController(w, r, http.StatusInternalServerError, "")
		return
	}
	categories, err := models.GetCategories()
	if err!= nil {
        ErrorController(w, r, http.StatusInternalServerError, "")
        return
    }
	if !utils.IsLoggedIn(r) {
		logedIn = false
	} else {
		logedIn = true
	}
	posts, err := models.GetPosts()
	if err != nil {
		fmt.Println(err)
		ErrorController(w, r, http.StatusInternalServerError, "")
        return
	}	

	data := models.Data{
		IsLoggedIn: logedIn,
		Comment:    comments,
		Category: categories,
		Posts:    posts,
	}

	if r.Method == "GET" {
		if r.URL.Path != "/" {
			ErrorController(w, r, http.StatusNotFound, http.StatusText(http.StatusNotFound))
			return
		}
		ParseFileController(w, r, "users/index", data)
	} else {
		ErrorController(w, r, http.StatusMethodNotAllowed, "")
		return
	}
}
