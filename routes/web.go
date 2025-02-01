package routes

import (
	"forum/app/http/controllers"
	"forum/app/http/middleware"
	"net/http"
)

func WebRouter() {
	http.HandleFunc("/",controllers.HomeController)
	http.HandleFunc("/PostByCategories", controllers.PostByCategoriesController)
	http.HandleFunc("/createdPost",middleware.AuthMiddleware(middleware.RateLimitMiddleware(  controllers.CreatedPostController)))
	http.HandleFunc("/myliked",middleware.AuthMiddleware(middleware.RateLimitMiddleware( controllers.LikedPostController)))
	http.HandleFunc("/login", middleware.AlreadyLoggedIn(middleware.RateLimitMiddleware(controllers.ParseLogin)))
	http.HandleFunc("/register", middleware.AlreadyLoggedIn(middleware.RateLimitMiddleware(controllers.ParseRegister)))
}
