package routes

import (
	"forum/app/http/controllers"
	"forum/app/http/middleware"
	"net/http"
)

func WebRouter() {
	http.HandleFunc("/", controllers.HomeController)
	http.HandleFunc("/PostByCategories", controllers.PostByCategoriesController)
	http.HandleFunc("/createdPost",middleware.AuthMiddleware(  controllers.CreatedPostController))
	http.HandleFunc("/myliked",middleware.AuthMiddleware( controllers.LikedPostController))
	http.HandleFunc("/login", middleware.AlreadyLoggedIn(controllers.ParseLogin))
	http.HandleFunc("/register", middleware.AlreadyLoggedIn(controllers.ParseRegister))
}
