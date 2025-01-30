package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"forum/app/models"
	"forum/utils"
)

func ReactPostController(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		ErrorController(w, r, http.StatusMethodNotAllowed, "")
		return
	}
	var react models.React

	var logedIn bool

	if !utils.IsLoggedIn(r) {
		logedIn = false
	} else {
		logedIn = true
	}
	var iduser int
	var err error
	if logedIn {
		iduser, err = models.GetUserId(r)
		if err != nil {
			LogoutController(w,r)
			return
		}
	}
	react.UserID = iduser
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
			ErrorController(w, r, http.StatusInternalServerError, "Cannot Insert React Posts")
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
			ErrorController(w, r, http.StatusInternalServerError, "Cannot Insert React Comments")
			return
		}

	}
	posts, err := models.GetPosts()
	if err != nil {
		ErrorController(w, r, http.StatusInternalServerError, "comment id not an integer")
		return
	}
	for i := range posts {
		likePost, err := models.GetReactionPost(posts[i].ID, "like")
		if err != nil {
			ErrorController(w, r, http.StatusInternalServerError, "comment id not an integer")
			return
		}
		posts[i].Likes = len(likePost)

		for _, reaction := range likePost {
			if reaction.UserID == iduser {
				posts[i].IsLiked = true
				break
			}
		}
	
		dislikePost, err := models.GetReactionPost(posts[i].ID, "dislike")
		if err != nil {
			ErrorController(w, r, http.StatusInternalServerError, "comment id not an integer")
			return
		}
		posts[i].Dislikes = len(dislikePost)
		for _, reaction := range dislikePost {
			if reaction.UserID == iduser {
				posts[i].IsDisliked = true
				break
			}
		}
		comment, err := models.GetComments(posts[i].ID)
		if err != nil {
			ErrorController(w, r, http.StatusInternalServerError, "comment id not an integer")
			return
		}
		for i := range comment {
			likeComment, err := models.GetReactionComment(comment[i].ID, "like")
			if err != nil {
				ErrorController(w, r, http.StatusInternalServerError, "comment id not an integer")
				return
			}
			comment[i].Likes = len(likeComment)
			
			dislikeComment, err := models.GetReactionComment(comment[i].ID, "dislike")
			if err != nil {
				ErrorController(w, r, http.StatusInternalServerError, "comment id not an integer")
				return
			}
			comment[i].Likes = len(likeComment)
			for _, reaction := range likeComment {
				if reaction.UserID == iduser {
					comment[i].IsLiked = true
					break
				}
			}
			for _, reaction := range dislikeComment {
				if reaction.UserID == iduser {
					comment[i].IsDisliked = true
					break
				}
			}
		
			comment[i].Dislikes = len(dislikeComment)

		}
		posts[i].Comments = comment

	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(posts)
}
