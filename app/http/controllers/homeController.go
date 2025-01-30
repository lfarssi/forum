package controllers

import (
	"fmt"
	"net/http"

	"forum/app/models"
	"forum/utils"
)

func HomeController(w http.ResponseWriter, r *http.Request) {
	var logedIn bool
	
	categories, err := models.GetCategories()
	if err!= nil {
		fmt.Println(err)
        ErrorController(w, r, http.StatusInternalServerError, "Cannot Fetch Category")
        return
    }
	if !utils.IsLoggedIn(r) {
		logedIn = false
	} else {
		logedIn = true
	}
	var iduser int
	if logedIn {
		iduser, err = models.GetUserId(r)
		if err != nil {
			LogoutController(w,r)
			return
		}
	}
	posts, err := models.GetPosts()
	if err != nil {
		
		ErrorController(w, r, http.StatusInternalServerError, "Cannot Fetch Post")
        return
	}	
	for i:= range posts {
		comment, err := models.GetComments(posts[i].ID)
		if err != nil {

			ErrorController(w, r, http.StatusInternalServerError, "Cannot Fetch Comment")
			return
		}
		

		likePost , err := models.GetReactionPost(posts[i].ID, "like")
		if err != nil {

			ErrorController(w, r, http.StatusInternalServerError, "Cannot Fetch like posts")
			return
		}
		posts[i].Likes = len(likePost)
		dislikePost , err := models.GetReactionPost(posts[i].ID, "dislike")
		if err != nil {

			ErrorController(w, r, http.StatusInternalServerError, "Cannot Fetch dislike posts")
			return
		}
		posts[i].Dislikes = len(dislikePost)
		
		for _, reaction := range likePost {
			if reaction.UserID == iduser {
				posts[i].IsLiked = true
				break
			}
		}
		for _, reaction := range dislikePost {
			if reaction.UserID == iduser {
				posts[i].IsDisliked = true
				break
			}
		}
		for i:= range comment {
			dislikecomment , err := models.GetReactionComment(comment[i].ID, "dislike")
		if err != nil {

			ErrorController(w, r, http.StatusInternalServerError, "Cannot Fetch dislike comment")
			return
		}
		comment[i].Dislikes= len(dislikecomment)
		likecomment , err := models.GetReactionComment(comment[i].ID, "like")
		if err != nil {
			ErrorController(w, r, http.StatusInternalServerError, "Cannot Fetch like comment")
			return
		}
		comment[i].Likes= len(likecomment)
		
		comment[i].Likes = len(likecomment)
			for _, reaction := range likecomment {
				if reaction.UserID == iduser {
					comment[i].IsLiked = true
					break
				}
			}
			for _, reaction := range dislikecomment {
				if reaction.UserID == iduser {
					comment[i].IsDisliked = true
					break
				}
			}
		}
		
		posts[i].Comments = comment
		posts[i].CommentsCount = len(comment)
	}
	data := models.Data{
		IsLoggedIn: logedIn,
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
