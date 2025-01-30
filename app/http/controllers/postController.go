package controllers

import (
	"html"
	"net/http"
	"strconv"

	"forum/app/models"
	"forum/utils"
)

func PostByCategoriesController(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		ErrorController(w, r, http.StatusMethodNotAllowed, "")
		return
	}
	categoriess, err := models.GetCategories()
	if err != nil {
		ErrorController(w, r, http.StatusInternalServerError, "")
		return
	}
	r.ParseForm()
	categories := r.Form["categories"]
	var posts []models.Posts
	postSet := make(map[int]struct{})
	for _, category := range categories {
		idCategorie, err := strconv.Atoi(category)
		if err != nil {
			ErrorController(w, r, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
			return
		}
		postTemp, err := models.GetPostsByCategory(idCategorie)
		if err != nil {
			ErrorController(w, r, http.StatusInternalServerError, "")
			return
		}
		for _, post := range postTemp {
			if _, exist := postSet[post.ID]; !exist {
				posts = append(posts, post)
				postSet[post.ID] = struct{}{}

			}
		}
	}
	data := models.Data{
		Category: categoriess,
		Posts:    posts,
	}
	ParseFileController(w, r, "users/index", data)
}

func LikedPostController(w http.ResponseWriter, r *http.Request) {
	var logedIn bool

	if !utils.IsLoggedIn(r) {
		logedIn = false
	} else {
		logedIn = true
	}
	categories, err := models.GetCategories()
	if err != nil {
		ErrorController(w, r, http.StatusInternalServerError, "")
		return
	}
	userId, err := models.GetUserId(r)
	if err != nil {
		LogoutController(w,r)
		return
	}
	likedpost, err := models.LikedPost(userId)
	if err != nil {
		ErrorController(w, r, http.StatusInternalServerError, "")
		return
	}
	for i := range likedpost {
		comment, err := models.GetComments(likedpost[i].ID)
		if err != nil {

			ErrorController(w, r, http.StatusInternalServerError, "")
			return
		}
		userID, err := models.GetUserId(r)
		if err != nil {
			ErrorController(w, r, http.StatusInternalServerError, "")
			return
		}
		likePost, err := models.GetReactionPost(likedpost[i].ID, "like")
		if err != nil {

			ErrorController(w, r, http.StatusInternalServerError, "")
			return
		}

		likedpost[i].Likes = len(likePost)
		dislikePost, err := models.GetReactionPost(likedpost[i].ID, "dislike")
		if err != nil {

			ErrorController(w, r, http.StatusInternalServerError, "")
			return
		}
		likedpost[i].Dislikes = len(dislikePost)

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

		for i := range comment {
			dislikecomment, err := models.GetReactionComment(comment[i].ID, "dislike")
			if err != nil {

				ErrorController(w, r, http.StatusInternalServerError, "")
				return
			}
			comment[i].Dislikes = len(dislikecomment)
			likecomment, err := models.GetReactionComment(comment[i].ID, "like")
			if err != nil {
				ErrorController(w, r, http.StatusInternalServerError, "")
				return
			}
			comment[i].Likes = len(likecomment)
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

	data := models.Data{
		Category:   categories,
		Posts:      likedpost,
		IsLoggedIn: logedIn,
	}
	ParseFileController(w, r, "users/index", data)
}

func CreatePosts(w http.ResponseWriter, r *http.Request) {
	title := html.EscapeString(r.PostFormValue("title"))
	category := r.PostForm["categories"]
	content := html.EscapeString(r.PostFormValue("content"))
	if title == "" || len(category) == 0 || content == "" {

		ErrorController(w, r, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	} else if len(content) > 10000 || len(title) > 500 {
		ErrorController(w, r, http.StatusBadRequest, "Content or Title should not exceed 10000 or 500 characters respectively")
		return
	}

	userId, err := models.GetUserId(r)
	if err != nil {
		LogoutController(w,r)
		return
	}

	idPost, err := models.CreatePost(title, content, category, userId)
	if err != nil {
		ErrorController(w, r, http.StatusInternalServerError, "")
		return
	}
	for _, c := range category {
		catId, err := strconv.Atoi(c)
		if err != nil {
			ErrorController(w, r, http.StatusBadRequest,"")
			return
		}
		err = models.InsertIntoCategoryPost(int(idPost), catId)
		if err != nil {
			ErrorController(w, r, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
			return
		}
	}
	http.Redirect(w, r, "/", http.StatusFound)
}


func CreatedPostController(w http.ResponseWriter, r *http.Request)  {
	var logedIn bool

	if !utils.IsLoggedIn(r) {
		logedIn = false
	} else {
		logedIn = true
	}
	categories, err := models.GetCategories()
	if err != nil {
		ErrorController(w, r, http.StatusInternalServerError, "")
		return
	}
	userId, err := models.GetUserId(r)
	if err != nil {
		ErrorController(w, r, http.StatusInternalServerError, "")
		return
	}
	createdPost, err := models.CreatedPost(userId)
	if err != nil {
		ErrorController(w, r, http.StatusInternalServerError, "")
		return
	}
	for i := range createdPost {
		comment, err := models.GetComments(createdPost[i].ID)
		if err != nil {

			ErrorController(w, r, http.StatusInternalServerError, "")
			return
		}
		userID, err := models.GetUserId(r)
		if err != nil {
			LogoutController(w,r)
			return
		}
		likePost, err := models.GetReactionPost(createdPost[i].ID, "like")
		if err != nil {

			ErrorController(w, r, http.StatusInternalServerError, "")
			return
		}

		createdPost[i].Likes = len(likePost)
		dislikePost, err := models.GetReactionPost(createdPost[i].ID, "dislike")
		if err != nil {

			ErrorController(w, r, http.StatusInternalServerError, "")
			return
		}
		createdPost[i].Dislikes = len(dislikePost)

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

		for i := range comment {
			dislikecomment, err := models.GetReactionComment(comment[i].ID, "dislike")
			if err != nil {

				ErrorController(w, r, http.StatusInternalServerError, "")
				return
			}
			comment[i].Dislikes = len(dislikecomment)
			likecomment, err := models.GetReactionComment(comment[i].ID, "like")
			if err != nil {
				ErrorController(w, r, http.StatusInternalServerError, "")
				return
			}
			comment[i].Likes = len(likecomment)
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

	data := models.Data{
		Category:   categories,
		Posts:      createdPost,
		IsLoggedIn: logedIn,
	}
	ParseFileController(w, r, "users/index", data)
}