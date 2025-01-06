package routes

import (
	"forum/app/http/controllers"
	"net/http"
)

func WebRouter() {
	http.HandleFunc("/", controllers.HomeController)
	http.HandleFunc("/posts", controllers.PostController)
	http.HandleFunc("/categories", controllers.CategoryController)
	http.HandleFunc("/comments", controllers.CommentController)
	http.HandleFunc("/reacts", controllers.ReactController)
	http.HandleFunc("/login", controllers.LoginController)
	http.HandleFunc("/register", controllers.RegisterController)
}
