package controllers

import (
	"fmt"
	"forum/app/models"
	"net/http"
	"time"
)

func PostController(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		ErrorController(w, r, http.StatusMethodNotAllowed, "")
		return
	}
	ParseFileController(w, r, "users/posts", nil)

}

func CreatePosts(w http.ResponseWriter, r *http.Request) {
	title := r.PostFormValue("title")
	category := r.PostForm["category"]
	content := r.PostFormValue("content")
	if title == "" || len(category) == 0 || content == "" {
		ErrorController(w, r, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}
	query := "INSERT INTO users ( title, content, user_id, creat_at) VALUES ( ?, ?, ?, ?)"
	stm, err := models.Database.Prepare(query)
	if err != nil {
		ErrorController(w, r, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
	defer stm.Close()
	res, err := stm.Exec( title,"user_id(get it from the the session)", content, time.Now())
	if err!= nil {
        ErrorController(w, r, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
        return
    }
	fmt.Println(res)
}
