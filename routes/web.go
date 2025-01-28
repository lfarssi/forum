package routes

import (
	"forum/app/http/controllers"
	"forum/app/http/middleware"
	"net/http"
)

func WebRouter() {
	http.HandleFunc("/", controllers.HomeController)
	http.HandleFunc("/PostByCategories", controllers.PostByCategoriesController)
	http.HandleFunc("/categories", controllers.CategoryController)
	http.HandleFunc("/comments", controllers.CommentController)
	http.HandleFunc("/myliked", controllers.LikedPostController)

	// http.HandleFunc("/reacts", controllers.ReactController)
	http.HandleFunc("/login", middleware.AlreadyLoggedIn(controllers.ParseLogin))
	http.HandleFunc("/register", middleware.AlreadyLoggedIn(controllers.ParseRegister))
}
