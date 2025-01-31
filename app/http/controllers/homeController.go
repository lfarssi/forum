package controllers

import (
	"fmt"
	"net/http"

	"forum/app/models"
	"forum/utils"
)

// HomeController handles the request for the homepage
func HomeController(w http.ResponseWriter, r *http.Request) {
	var logedIn bool
	
	// Get categories from the database
	categories, err := models.GetCategories()
	if err != nil {
		fmt.Println(err)
		// Handle error if categories cannot be fetched
        ErrorController(w, r, http.StatusInternalServerError, "Cannot Fetch Category")
        return
    }

	// Check if the user is logged in
	if !utils.IsLoggedIn(r) {
		logedIn = false
	} else {
		logedIn = true
	}

	var iduser int
	// If the user is logged in, get their user ID
	if logedIn {
		iduser, err = models.GetUserId(r)
		if err != nil {
			LogoutController(w, r) // Log out if there is an error getting user ID
			return
		}
	}

	// Get posts from the database
	posts, err := models.GetPosts()
	if err != nil {
		// Handle error if posts cannot be fetched
		ErrorController(w, r, http.StatusInternalServerError, "Cannot Fetch Post")
        return
	}

	// Loop through each post to add reactions (likes/dislikes) and comments
	for i := range posts {
		// Get the comments associated with the current post
		comment, err := models.GetComments(posts[i].ID)
		if err != nil {
			// Handle error if comments cannot be fetched
			ErrorController(w, r, http.StatusInternalServerError, "Cannot Fetch Comment")
			return
		}

		// Get the likes for the current post
		likePost, err := models.GetReactionPost(posts[i].ID, "like")
		if err != nil {
			// Handle error if like reactions cannot be fetched
			ErrorController(w, r, http.StatusInternalServerError, "Cannot Fetch like posts")
			return
		}
		posts[i].Likes = len(likePost) // Set the like count for the post

		// Get the dislikes for the current post
		dislikePost, err := models.GetReactionPost(posts[i].ID, "dislike")
		if err != nil {
			// Handle error if dislike reactions cannot be fetched
			ErrorController(w, r, http.StatusInternalServerError, "Cannot Fetch dislike posts")
			return
		}
		posts[i].Dislikes = len(dislikePost) // Set the dislike count for the post

		// Check if the logged-in user has liked or disliked this post
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

		// Loop through each comment on the post
		for _, commentItem := range comment {
			// Get the dislikes for the current comment
			dislikecomment, err := models.GetReactionComment(commentItem.ID, "dislike")
			if err != nil {
				// Handle error if dislike reactions for the comment cannot be fetched
				ErrorController(w, r, http.StatusInternalServerError, "Cannot Fetch dislike comment")
				return
			}
			commentItem.Dislikes = len(dislikecomment) // Set the dislike count for the comment

			// Get the likes for the current comment
			likecomment, err := models.GetReactionComment(commentItem.ID, "like")
			if err != nil {
				// Handle error if like reactions for the comment cannot be fetched
				ErrorController(w, r, http.StatusInternalServerError, "Cannot Fetch like comment")
				return
			}
			commentItem.Likes = len(likecomment) // Set the like count for the comment

			// Check if the logged-in user has liked or disliked this comment
			for _, reaction := range likecomment {
				if reaction.UserID == iduser {
					commentItem.IsLiked = true
					break
				}
			}
			for _, reaction := range dislikecomment {
				if reaction.UserID == iduser {
					commentItem.IsDisliked = true
					break
				}
			}
		}
		
		// Set the comments and comment count for the post
		posts[i].Comments = comment
		posts[i].CommentsCount = len(comment)
	}

	// Prepare the data to be passed to the template
	data := models.Data{
		IsLoggedIn: logedIn,
		Category: categories,
		Posts:    posts,
	}

	// Check if the request method is GET and the URL path is the homepage
	if r.Method == "GET" {
		if r.URL.Path != "/" {
			// Handle 404 error if the URL path is not the homepage
			ErrorController(w, r, http.StatusNotFound, http.StatusText(http.StatusNotFound))
			return
		}
		// Parse and render the homepage template with the data
		ParseFileController(w, r, "users/index", data)
	} else {
		// Handle method not allowed error for non-GET requests
		ErrorController(w, r, http.StatusMethodNotAllowed, "")
		return
	}
}
