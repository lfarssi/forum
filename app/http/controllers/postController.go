package controllers

import (
	"encoding/json"
	"html"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"forum/app/models"
	"forum/utils"
)

// PostByCategoriesController handles the request for posts filtered by categories
func PostByCategoriesController(w http.ResponseWriter, r *http.Request) {
	// Ensure the method is GET
	if r.Method != "GET" {
		ErrorController(w, r, http.StatusMethodNotAllowed, "")
		return
	}

	var logedIn bool
	// Check if the user is logged in
	if !utils.IsLoggedIn(r) {
		logedIn = false
	} else {
		logedIn = true
	}

	// Fetch categories from the database
	categoriess, err := models.GetCategories()
	if err != nil {
		ErrorController(w, r, http.StatusInternalServerError, "Cannot Fetch Category")
		return
	}

	// Parse the form data to get selected categories
	r.ParseForm()
	categories := r.Form["categories"]
	var posts []models.Posts
	postSet := make(map[int]struct{})

	// Loop through selected categories and fetch posts for each
	for _, category := range categories {
		idCategorie, err := strconv.Atoi(category)
		if err != nil {
			ErrorController(w, r, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
			return
		}
		postTemp, err := models.GetPostsByCategory(idCategorie)
		if err != nil {
			ErrorController(w, r, http.StatusInternalServerError, "Cannot Fetch Post by category")
			return
		}

		// Avoid duplicate posts by using a map
		for _, post := range postTemp {
			if _, exist := postSet[post.ID]; !exist {
				posts = append(posts, post)
				postSet[post.ID] = struct{}{}
			}
		}
	}
	for i := range posts {
		// Fetch comments for each post
		comment, err := models.GetComments(posts[i].ID)
		if err != nil {
			ErrorController(w, r, http.StatusInternalServerError, "Cannot Fetch comment")
			return
		}

		// Get the user's reactions (likes/dislikes) for the post and comments
		userID, err := models.GetUserId(r)
		if err != nil {
			LogoutController(w, r)
			return
		}
		likePost, err := models.GetReactionPost(posts[i].ID, "like")
		if err != nil {
			ErrorController(w, r, http.StatusInternalServerError, "Cannot get posts likes")
			return
		}
		posts[i].Likes = len(likePost)
		dislikePost, err := models.GetReactionPost(posts[i].ID, "dislike")
		if err != nil {
			ErrorController(w, r, http.StatusInternalServerError, "Cannot get posts dislikes")
			return
		}
		posts[i].Dislikes = len(dislikePost)

		// Check if the user liked or disliked the post
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

		// Process each comment on the liked post
		for i := range comment {
			dislikecomment, err := models.GetReactionComment(comment[i].ID, "dislike")
			if err != nil {
				ErrorController(w, r, http.StatusInternalServerError, "Cannot comment dislikes")
				return
			}
			comment[i].Dislikes = len(dislikecomment)
			likecomment, err := models.GetReactionComment(comment[i].ID, "like")
			if err != nil {
				ErrorController(w, r, http.StatusInternalServerError, "Cannot get comment likes")
				return
			}
			comment[i].Likes = len(likecomment)
			// Check if the user liked or disliked the comment
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
	// Prepare data for rendering the page
	data := models.Data{
		Category:   categoriess,
		Posts:      posts,
		IsLoggedIn: logedIn,
	}

	// Render the posts filtered by categories
	ParseFileController(w, r, "users/index", data)
}

// LikedPostController handles the request for posts liked by the logged-in user
func LikedPostController(w http.ResponseWriter, r *http.Request) {
	var logedIn bool

	// Check if the user is logged in
	if !utils.IsLoggedIn(r) {
		logedIn = false
	} else {
		logedIn = true
	}

	// Fetch categories from the database
	categories, err := models.GetCategories()
	if err != nil {
		ErrorController(w, r, http.StatusInternalServerError, "Cannot Fetch Category")
		return
	}

	// Get the user ID
	userId, err := models.GetUserId(r)
	if err != nil {
		LogoutController(w, r)
		return
	}

	// Get the posts liked by the user
	likedpost, err := models.LikedPost(userId)
	if err != nil {
		ErrorController(w, r, http.StatusInternalServerError, "Cannot liked posts")
		return
	}

	// Process each liked post
	for i := range likedpost {
		// Fetch comments for each post
		comment, err := models.GetComments(likedpost[i].ID)
		if err != nil {
			ErrorController(w, r, http.StatusInternalServerError, "Cannot Fetch comment")
			return
		}

		// Get the user's reactions (likes/dislikes) for the post and comments
		userID, err := models.GetUserId(r)
		if err != nil {
			LogoutController(w, r)
			return
		}
		likePost, err := models.GetReactionPost(likedpost[i].ID, "like")
		if err != nil {
			ErrorController(w, r, http.StatusInternalServerError, "Cannot get posts likes")
			return
		}
		likedpost[i].Likes = len(likePost)
		dislikePost, err := models.GetReactionPost(likedpost[i].ID, "dislike")
		if err != nil {
			ErrorController(w, r, http.StatusInternalServerError, "Cannot get posts dislikes")
			return
		}
		likedpost[i].Dislikes = len(dislikePost)

		// Check if the user liked or disliked the post
		for _, reaction := range likePost {
			if reaction.UserID == userID {
				likedpost[i].IsLiked = true
				break
			}
		}
		for _, reaction := range dislikePost {
			if reaction.UserID == userID {
				likedpost[i].IsDisliked = true
				break
			}
		}

		// Process each comment on the liked post
		for i := range comment {
			dislikecomment, err := models.GetReactionComment(comment[i].ID, "dislike")
			if err != nil {
				ErrorController(w, r, http.StatusInternalServerError, "Cannot comment dislikes")
				return
			}
			comment[i].Dislikes = len(dislikecomment)
			likecomment, err := models.GetReactionComment(comment[i].ID, "like")
			if err != nil {
				ErrorController(w, r, http.StatusInternalServerError, "Cannot get comment likes")
				return
			}
			comment[i].Likes = len(likecomment)
			// Check if the user liked or disliked the comment
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
		likedpost[i].Comments = comment
		likedpost[i].CommentsCount = len(comment)
	}

	// Prepare data for rendering the liked posts page
	data := models.Data{
		Category:   categories,
		Posts:      likedpost,
		IsLoggedIn: logedIn,
	}

	// Render the liked posts page
	ParseFileController(w, r, "users/index", data)
}

// CreatePosts handles the creation of a new post
func CreatePosts(w http.ResponseWriter, r *http.Request) {
	// Get the post data from the form
	err := r.ParseMultipartForm(20 >> 20) // 20mb limit
	if err != nil {
		ErrorController(w, r, http.StatusBadRequest, "image to big")
		return
	}
	
	title := html.EscapeString(r.PostFormValue("title"))
	

	category := r.PostForm["categories"]
	content := html.EscapeString(r.PostFormValue("content"))
	file, header, err := r.FormFile("image")
	
	var filePath string
	if file != nil{
		if float64( header.Size)/ (1024 * 1024)  > 20 {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode("Error:Image bigger than 20 mb")		
			return
		}
		mimeType := header.Header.Get("Content-Type")

		if !strings.HasPrefix(mimeType, "image/") {
			w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Error:Only image allowed")
			return
		}
		
		if err != nil {
	
			w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Error:Cannot get the image")
			return
		}
		defer file.Close()
		filePath = filepath.Join("./resources/storage", time.Now().Format("2006-01-02_15-04-05") + "_" + header.Filename)
		if err := os.MkdirAll("./resources/storage", os.ModePerm); err != nil {
	
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode("Error:Internal server error ")			
			return
		}
		dst, err := os.Create(filePath)
		if err != nil {
	
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode("Error:Internal server error ")			
			return
		}
		defer dst.Close()
		_, err = io.Copy(dst, file)
		if err != nil {
	
	
	
			w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Error:Internal serever error")
			return
		}
	}
	

	// Validate the input fields
	if strings.TrimSpace(title) == "" || len(category) == 0 || strings.TrimSpace(content) == "" || len(category) == 0  {
		
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Error:Title or Content or Categorie field's  empty ")
		return
	} else if len(content) > 10000 || len(title) > 255 {

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusRequestEntityTooLarge)
		json.NewEncoder(w).Encode("Error: Title or Content field's too large")
		return
	}

	// Get the user ID
	userId, err := models.GetUserId(r)
	if err != nil {

		LogoutController(w, r)
		return
	}

	// Create the new post
	idPost, err := models.CreatePost(title, content, filePath, category, userId)
	if err != nil {

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Error: Cannot create post")
		return
	}

	// Assign the post to the selected categories
	for _, c := range category {
		catId, err := strconv.Atoi(c)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode("error: cannot loop on category")
			return
		}
		err = models.InsertIntoCategoryPost(int(idPost), catId)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode("error: cannot insert into category")
			return
		}
	}

	// Respond with success
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func DeleteController(w http.ResponseWriter, r *http.Request) {
	query := `DELETE FROM posts`
	_, err := models.Database.Exec(query)
	if err != nil {
		return
	}
}

// CreatedPostController handles the display of posts created by the logged-in user
func CreatedPostController(w http.ResponseWriter, r *http.Request) {
	var logedIn bool

	// Check if the user is logged in
	if !utils.IsLoggedIn(r) {
		logedIn = false
	} else {
		logedIn = true
	}

	// Fetch categories from the database
	categories, err := models.GetCategories()
	if err != nil {
		ErrorController(w, r, http.StatusInternalServerError, "Cannot Fetch Category")
		return
	}

	// Get the user ID
	userId, err := models.GetUserId(r)
	if err != nil {
		LogoutController(w, r)
		return
	}

	// Get the posts created by the user
	createdPost, err := models.CreatedPost(userId)
	if err != nil {
		ErrorController(w, r, http.StatusInternalServerError, "Cannot Create Posts")
		return
	}

	// Process each created post
	for i := range createdPost {
		comment, err := models.GetComments(createdPost[i].ID)
		if err != nil {
			ErrorController(w, r, http.StatusInternalServerError, "Cannot Get comments")
			return
		}

		// Get reactions (likes/dislikes) for the post
		userID, err := models.GetUserId(r)
		if err != nil {
			LogoutController(w, r)
			return
		}
		likePost, err := models.GetReactionPost(createdPost[i].ID, "like")
		if err != nil {
			ErrorController(w, r, http.StatusInternalServerError, "Cannot like posts")
			return
		}
		createdPost[i].Likes = len(likePost)
		dislikePost, err := models.GetReactionPost(createdPost[i].ID, "dislike")
		if err != nil {
			ErrorController(w, r, http.StatusInternalServerError, "Cannot get dislikes posts")
			return
		}
		createdPost[i].Dislikes = len(dislikePost)

		// Check if the user liked or disliked the post
		for _, reaction := range likePost {
			if reaction.UserID == userID {
				createdPost[i].IsLiked = true
				break
			}
		}
		for _, reaction := range dislikePost {
			if reaction.UserID == userID {
				createdPost[i].IsDisliked = true
				break
			}
		}

		// Process each comment on the created post
		for i := range comment {
			dislikecomment, err := models.GetReactionComment(comment[i].ID, "dislike")
			if err != nil {
				ErrorController(w, r, http.StatusInternalServerError, "Cannot get dislikes comment")
				return
			}
			comment[i].Dislikes = len(dislikecomment)
			likecomment, err := models.GetReactionComment(comment[i].ID, "like")
			if err != nil {
				ErrorController(w, r, http.StatusInternalServerError, "Cannot get likes comments")
				return
			}
			comment[i].Likes = len(likecomment)
			// Check if the user liked or disliked the comment
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
		createdPost[i].Comments = comment
		createdPost[i].CommentsCount = len(comment)
	}

	// Prepare data for rendering the created posts page
	data := models.Data{
		Category:   categories,
		Posts:      createdPost,
		IsLoggedIn: logedIn,
	}

	// Render the created posts page
	ParseFileController(w, r, "users/index", data)
}
