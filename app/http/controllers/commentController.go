package controllers

import (
	"encoding/json"
	"html"
	"net/http"
	"strconv"
	"time"

	"forum/app/models"
)

func CreateCommentController(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		ErrorController(w, r, http.StatusMethodNotAllowed, "")
		return
	}
	var err error
	var comment models.Comment
	comment.Content = html.EscapeString(r.FormValue("content"))
	comment.CreatedAt = time.Now()
	comment.PostID, err = strconv.Atoi(r.FormValue("post_id"))
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"post": "id not an integer",
		})
		return
	}
	if len(comment.Content) == 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
		"content": "content is empty",
		})			
		return
	} else if len(comment.Content) >= 10000 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusRequestEntityTooLarge)
		json.NewEncoder(w).Encode(map[string]interface{}{
		"content": "content too large",
		})			
		return
	}
	userId, err := models.GetUserId(r)
	if err != nil {
		LogoutController(w, r)
		return
	}
	comment.UserID = userId

	if err = models.CreateComment(comment); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
		"error": "cannot  create comment",
		})			
		return
	}
}
