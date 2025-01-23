package controllers

import (
	"net/http"
	"time"

	"forum/app/models"
)

func CommentController(w http.ResponseWriter, r *http.Request) {
}

func CreatCommentController(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		ErrorController(w, r, http.StatusMethodNotAllowed, "")
		return
	}
	r.ParseForm()
	var comment models.Comment
	comment.Content = r.FormValue("content")
	comment.CreatedAt= time.Now()
	comment.PostID= r.FormValue("post_id")
	comment.UserID= 1 // func get user after
	error := models.CreatComment(comment)
}
