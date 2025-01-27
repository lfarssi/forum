package controllers

import (
	"net/http"
	"strconv"
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
	var err error
	var comment models.Comment
	comment.Content = r.FormValue("content")
	comment.CreatedAt= time.Now()
	comment.PostID, err =strconv.Atoi( r.FormValue("post_id"))
	if err != nil{
		ErrorController(w,r , http.StatusBadRequest, "Post number Invalid")
		return 
	}
	comment.UserID = 1 // func get user after
	
	if err = models.CreatComment(comment) ; err != nil {
		ErrorController(w, r, http.StatusInternalServerError, "")
		return
	}
	http.Redirect(w,r , "/", http.StatusFound)
}
