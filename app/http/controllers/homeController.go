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
	for i:= range posts {
		comment, err := models.GetComments(posts[i].ID)
		if err != nil {
			fmt.Println("get comment")

			ErrorController(w, r, http.StatusInternalServerError, "")
			return
		}
		userID,err:=models.GetUserId(r)
		if err!=nil{
			ErrorController(w, r, http.StatusInternalServerError, "")
			return
		}

		likePost , err := models.GetReactionPost(posts[i].ID, "like")
		if err != nil {
			fmt.Println("like post")

			ErrorController(w, r, http.StatusInternalServerError, "")
			return
		}
		fmt.Println("like:",likePost)
		posts[i].Likes = len(likePost)
		dislikePost , err := models.GetReactionPost(posts[i].ID, "dislike")
		if err != nil {
			fmt.Println("dislike post")

			ErrorController(w, r, http.StatusInternalServerError, "")
			return
		}
		posts[i].Dislikes = len(dislikePost)

		for _, reaction := range likePost {
			if reaction.UserID == userID {
				posts[i].IsLiked = true
				break
			}
		}
		for _, reaction := range dislikePost {
			if reaction.UserID == userID {
				posts[i].IsDisliked = true
				break
			}
		}


		for i:= range comment {
			dislikecomment , err := models.GetReactionComment(comment[i].ID, "dislike")
		if err != nil {
			fmt.Println("dislike comment")

			ErrorController(w, r, http.StatusInternalServerError, "")
			return
		}
		comment[i].Dislikes= len(dislikecomment)
		likecomment , err := models.GetReactionComment(comment[i].ID, "like")
		if err != nil {
			fmt.Println("like comment")
			ErrorController(w, r, http.StatusInternalServerError, "")
			return
		}
		comment[i].Likes= len(likecomment)
		for _, reaction := range likecomment {
			if reaction.UserID == userID {
				comment[i].IsLiked = true
				break
			}
		}
		for _, reaction := range dislikecomment {
			if reaction.UserID == userID {
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
