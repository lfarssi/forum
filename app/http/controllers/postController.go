package controllers

import (
	"encoding/json"
	"html"
	"net/http"
	"strconv"
	"strings"

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
	var user models.User
	var userID int
	var err error

	if utils.IsLoggedIn(r) {
		logedIn = true
		userID, err = models.GetUserId(r)
		if err != nil {
			LogoutController(w, r)
			return
		}
		user.Role, err = models.GetRoleUser(userID)
		if err != nil {
			LogoutController(w, r)
			return
		}
	} else {
		logedIn = false
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

	if user.Role == "user" {
		reqmod, erro := models.GetRequestInfo(userID)
		if erro != nil {
			ErrorController(w, r, http.StatusInternalServerError, "Cannot Fetch the Request info")
			return
		}

		categorie_report, err := models.GetCategorieReport()
		if err != nil {
			ErrorController(w, r, http.StatusInternalServerError, "Cannot Fetch the Categorie Report")
			return
		}
		datas := models.Data{
			IsLoggedIn:     logedIn,
			Category:       categoriess,
			Posts:          posts,
			Role:           user.Role,
			StatusReq:      reqmod,
			CategoryReport: categorie_report,
		}
		ParseFileController(w, r, "users/index", datas)

	} else if user.Role == "moderator" {
		categorie_report, err := models.GetCategorieReport()
		if err != nil {
			ErrorController(w, r, http.StatusInternalServerError, "Cannot Fetch the Categorie Report")
			return
		}

		data := models.Data{
			IsLoggedIn:     logedIn,
			Category:       categoriess,
			Posts:          posts,
			Role:           user.Role,
			CategoryReport: categorie_report,
		}
		ParseFileController(w, r, "moderator/index", data)

	} else if user.Role == "admin" {
		// You can fetch and pass additional admin data here

		categorie_report, err := models.GetCategorieReport()
		if err != nil {
			ErrorController(w, r, http.StatusInternalServerError, "Cannot Fetch the Categorie Report")
			return
		}
		data := models.Data{
			IsLoggedIn:     logedIn,
			Category:       categoriess,
			Posts:          posts,
			Role:           user.Role,
			CategoryReport: categorie_report,
		}
		ParseFileController(w, r, "admin/index", data)

	} else {

		// Guest view
		data := models.Data{
			IsLoggedIn: logedIn,
			Category:   categoriess,
			Posts:      posts,
			Role:       "guest",
		}
		ParseFileController(w, r, "guests/index", data)
	}
}

// LikedPostController handles the request for posts liked by the logged-in user
func LikedPostController(w http.ResponseWriter, r *http.Request) {
	var logedIn bool
	var user models.User
	var userID int
	var err error

	if utils.IsLoggedIn(r) {
		logedIn = true
		userID, err = models.GetUserId(r)
		if err != nil {
			LogoutController(w, r)
			return
		}
		user.Role, err = models.GetRoleUser(userID)
		if err != nil {
			LogoutController(w, r)
			return
		}
	} else {
		logedIn = false
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

	if user.Role == "user" {
		reqmod, erro := models.GetRequestInfo(userID)
		if erro != nil {
			ErrorController(w, r, http.StatusInternalServerError, "Cannot Fetch the Request info")
			return
		}
		categorie_report, err := models.GetCategorieReport()
		if err != nil {
			ErrorController(w, r, http.StatusInternalServerError, "Cannot Fetch the Categorie Report")
			return
		}

		datas := models.Data{
			IsLoggedIn:     logedIn,
			Category:       categories,
			Posts:          likedpost,
			Role:           user.Role,
			StatusReq:      reqmod,
			CategoryReport: categorie_report,
		}
		ParseFileController(w, r, "users/index", datas)

	} else if user.Role == "moderator" {
		categorie_report, err := models.GetCategorieReport()
		if err != nil {
			ErrorController(w, r, http.StatusInternalServerError, "Cannot Fetch the Categorie Report")
			return
		}
		data := models.Data{
			IsLoggedIn:     logedIn,
			Category:       categories,
			Posts:          likedpost,
			Role:           user.Role,
			CategoryReport: categorie_report,
		}
		ParseFileController(w, r, "moderator/index", data)

	} else if user.Role == "admin" {

		modRequests, err := models.GetAllModRequests()
		if err != nil {
			ErrorController(w, r, http.StatusInternalServerError, "Cannot Fetch moderator requests")
			return
		}
		categorie_report, err := models.GetCategorieReport()
		if err != nil {
			ErrorController(w, r, http.StatusInternalServerError, "Cannot Fetch the Categorie Report")
			return
		}

		data := models.Data{
			IsLoggedIn:     logedIn,
			Category:       categories,
			Posts:          likedpost,
			Role:           user.Role,
			ModRequests:    modRequests,
			CategoryReport: categorie_report,
		}
		ParseFileController(w, r, "admin/index", data)

	} else {
		// Guest view
		data := models.Data{
			IsLoggedIn: logedIn,
			Category:   categories,
			Posts:      likedpost,
			Role:       "guest",
		}
		ParseFileController(w, r, "guests/index", data)
	}
}

// CreatePosts handles the creation of a new post
func CreatePosts(w http.ResponseWriter, r *http.Request) {

	// Get the post data from the form
	title := html.EscapeString(r.PostFormValue("title"))
	category := r.PostForm["categories"]
	content := html.EscapeString(r.PostFormValue("content"))

	// Validate the input fields
	if strings.TrimSpace(title) == "" || len(category) == 0 || strings.TrimSpace(content) == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Error:Title or Content field's  empty ")
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
	idPost, err := models.CreatePost(title, content, category, userId)
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
	var user models.User
	var userID int
	var err error

	if utils.IsLoggedIn(r) {
		logedIn = true
		userID, err = models.GetUserId(r)
		if err != nil {
			LogoutController(w, r)
			return
		}
		user.Role, err = models.GetRoleUser(userID)
		if err != nil {
			LogoutController(w, r)
			return
		}
	} else {
		logedIn = false
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

	if user.Role == "user" {
		reqmod, erro := models.GetRequestInfo(userID)
		if erro != nil {
			ErrorController(w, r, http.StatusInternalServerError, "Cannot Fetch the Request info")
			return
		}
		categorie_report, err := models.GetCategorieReport()
		if err != nil {
			ErrorController(w, r, http.StatusInternalServerError, "Cannot Fetch the Categorie Report")
			return
		}
		datas := models.Data{
			IsLoggedIn:     logedIn,
			Category:       categories,
			Posts:          createdPost,
			Role:           user.Role,
			StatusReq:      reqmod,
			CategoryReport: categorie_report,
		}
		ParseFileController(w, r, "users/index", datas)

	} else if user.Role == "moderator" {
		categorie_report, err := models.GetCategorieReport()
		if err != nil {
			ErrorController(w, r, http.StatusInternalServerError, "Cannot Fetch the Categorie Report")
			return
		}

		data := models.Data{
			IsLoggedIn:     logedIn,
			Category:       categories,
			Posts:          createdPost,
			Role:           user.Role,
			CategoryReport: categorie_report,
		}
		ParseFileController(w, r, "moderator/index", data)

	} else if user.Role == "admin" {

		modRequests, err := models.GetAllModRequests()
		if err != nil {
			ErrorController(w, r, http.StatusInternalServerError, "Cannot Fetch moderator requests")
			return
		}
		categorie_report, err := models.GetCategorieReport()
		if err != nil {
			ErrorController(w, r, http.StatusInternalServerError, "Cannot Fetch the Categorie Report")
			return
		}
		data := models.Data{
			IsLoggedIn:     logedIn,
			Category:       categories,
			Posts:          createdPost,
			Role:           user.Role,
			ModRequests:    modRequests,
			CategoryReport: categorie_report,
		}
		ParseFileController(w, r, "admin/index", data)

	} else {
		// Guest view
		data := models.Data{
			IsLoggedIn: logedIn,
			Category:   categories,
			Role:       "guest",
		}
		ParseFileController(w, r, "guests/index", data)
	}
}

func ReportPostController(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		ErrorController(w, r, http.StatusMethodNotAllowed, "Method Not Allowed")
		return
	}

	postIDStr := r.FormValue("post_id")
	categoryIDStr := r.FormValue("category_report_id")

	if postIDStr == "" || categoryIDStr == "" {
		ErrorController(w, r, http.StatusBadRequest, "Post ID and Category are required")
		return
	}

	postID, err := strconv.Atoi(postIDStr)
	if err != nil {
		ErrorController(w, r, http.StatusBadRequest, "Invalid Post ID")
		return
	}

	categoryID, err := strconv.Atoi(categoryIDStr)
	if err != nil {
		ErrorController(w, r, http.StatusBadRequest, "Invalid Category ID")
		return
	}

	userID, err := models.GetUserId(r)
	if err != nil {
		ErrorController(w, r, http.StatusUnauthorized, "Unauthorized")
		return
	}

	err = models.ReportPost(postID, userID, categoryID)
	if err != nil {
		ErrorController(w, r, http.StatusInternalServerError, "Failed to report post")
		return
	}

	// If it's a fetch (JS), return OK. Otherwise redirect.
	if r.Header.Get("X-Requested-With") == "XMLHttpRequest" {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Reported successfully"))
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
