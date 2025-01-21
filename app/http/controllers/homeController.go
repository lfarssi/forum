package controllers

import (
	"forum/app/models"
	"forum/utils"
	"net/http"
	"time"
)

func HomeController(w http.ResponseWriter, r *http.Request) {
	var logedIn bool
	if !utils.IsLoggedIn(r) {
		logedIn = false
	} else {
		logedIn = true
	}
	data := struct {
		IsLoggedIn bool
		Post       []models.Posts
		Comment    []models.Comment
		Category   []models.Category
	}{
		IsLoggedIn: logedIn,
	}
	if r.Method == "GET" {
		if r.URL.Path != "/" {
			ErrorController(w, r, http.StatusNotFound, http.StatusText(http.StatusNotFound))
			return
		}
		ParseFileController(w, r, "users/index", data)
	} else if r.Method == "POST" {
		title := r.PostFormValue("title")
		category := r.PostForm["category"]
		content := r.PostFormValue("content")
		if title == "" || len(category) == 0 || content == "" {
			ErrorController(w, r, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
			return
		}
		query := "INSERT INTO users (id, title, content, user_id, creat_at) VALUES (?, ?, ?, ?, ?)"
		stm, err := models.Database.Prepare(query)
		if err != nil {
			ErrorController(w, r, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
			return
		}
		defer stm.Close()
		res, err := stm.Exec(user.UserName, title, content,,time.Now())
	} else {
		ErrorController(w, r, http.StatusMethodNotAllowed, "")
		return
	}
}
