package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"forum/app/models"
)

func PostByCategoriesController(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		ErrorController(w, r, http.StatusMethodNotAllowed, "")
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
			if _, exist:=postSet[post.ID] ; !exist {
				posts = append(posts, post)
				postSet[post.ID] = struct{}{} 

			}
        }
	}
	data := models.Data{
		Posts: posts,
	}
	ParseFileController(w, r, "users/index", data)
}

func CreatePosts(w http.ResponseWriter, r *http.Request) {
	title := r.PostFormValue("title")
	category := r.PostForm["categories"]
	content := r.PostFormValue("content")
	if title == "" || len(category) == 0 || content == "" {

		ErrorController(w, r, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	} else if len(content) > 10000 || len(title) > 500 {
		ErrorController(w, r, http.StatusBadRequest, "Content or Title should not exceed 10000 or 500 characters respectively")
		return
	}

	userId, err := models.GetUserId(r)
	if err != nil {
		ErrorController(w, r, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	idPost, err := models.CreatePost(title, content, category, userId)
	if err != nil {
		ErrorController(w, r, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
	for _, c := range category {
		catId, err := strconv.Atoi(c)
		if err != nil {
			fmt.Println("atoi")
			ErrorController(w, r, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
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
