package controllers

import (
	"encoding/json"
	"html"
	"net/http"
	"strconv"
	"time"

	"forum/app/models"
	"forum/utils"
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
			"post": "Error: Id post not an integer",
		})
		return
	}
	if len(comment.Content) == 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"content": "Error: Content of the comment is empty",
		})
		return
	} else if len(comment.Content) >= 10000 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusRequestEntityTooLarge)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"content": "Error: Content of the comment too large",
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
			"error": "Error: Cannot  create comment",
		})
		return
	}
}

func DeleteCommentHandler(w http.ResponseWriter, r *http.Request) {
	var logedIn bool
	if r.Method != http.MethodPost {
		ErrorController(w, r, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	commentIDStr := r.FormValue("comment_id")
	if commentIDStr == "" {
		ErrorController(w, r, http.StatusBadRequest, "Comment ID is required")
		return
	}

	commentID, err := strconv.Atoi(commentIDStr)
	if err != nil {
		ErrorController(w, r, http.StatusBadRequest, "Invalid Comment ID")
		return
	}

	if !utils.IsLoggedIn(r) {
		logedIn = false
	} else {
		logedIn = true
	}
	var user models.User

	var iduser int
	
	if logedIn {
		iduser, err = models.GetUserId(r)
		if err != nil {
			LogoutController(w, r) // Log out if there is an error getting user ID
			return
		}
		user.Role, err = models.GetRoleUser(iduser)
		if err != nil {
			LogoutController(w, r) // Log out if there is an error getting user ID
			return
		}
	}

	// Get session user to check role
	
	if user.Role != "admin" {
		ErrorController(w, r, http.StatusForbidden, "You are not authorized to delete comments")
		return
	}

	// Check if the comment exists
	exists, err := models.CommentExists(commentID)
	if err != nil {
		ErrorController(w, r, http.StatusInternalServerError, "Server error while checking comment")
		return
	}
	if !exists {
		ErrorController(w, r, http.StatusNotFound, "Comment not found")
		return
	}

	// Delete the comment
	err = models.DeleteComment(commentID)
	if err != nil {
		ErrorController(w, r, http.StatusInternalServerError, "Failed to delete comment")
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
