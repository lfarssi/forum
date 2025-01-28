package controllers

import (
	"net/http"
	"strconv"

	"forum/app/models"
)

func ReactPostController(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		ErrorController(w, r, http.StatusMethodNotAllowed, "")
		return
	}
	var react models.React

	react.UserID = 1
	react.Status = r.FormValue("status")
	react.Sender = r.FormValue("sender")
	if react.Sender == "post" {
		postID, err := strconv.Atoi(r.FormValue("post_id"))
		if err != nil {
			ErrorController(w, r, http.StatusInternalServerError, "post id not an integer")
			return
		}
		react.PostID = postID

		err = models.InsertReactPost(react)
		if err != nil {
			ErrorController(w, r, http.StatusInternalServerError, "")
			return
		}
	} else if react.Sender == "comment" {
		commentID, err := strconv.Atoi(r.FormValue("comment_id"))
		if err != nil {
			ErrorController(w, r, http.StatusInternalServerError, "comment id not an integer")
			return
		}
		react.CommentID = commentID
		err = models.InsertReactComment(react)
		if err != nil {
			ErrorController(w, r, http.StatusInternalServerError, "")
			return
		}
	}
}
