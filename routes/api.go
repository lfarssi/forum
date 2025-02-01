package routes

import (
	"forum/app/http/controllers"
	"forum/app/http/middleware"
	"net/http"
)

func ApiRouter() {

	http.HandleFunc("/signIn",middleware.RateLimitMiddleware( controllers.LoginController))
	http.HandleFunc("/signUp",middleware.RateLimitMiddleware( controllers.RegisterController))
	http.HandleFunc("/logout",middleware.RateLimitMiddleware( controllers.LogoutController))
	// http.HandleFunc("/delete_post", controllers.DeleteController)
	http.HandleFunc("/create_post", middleware.AuthMiddleware(middleware.RateLimitMiddleware( controllers.CreatePosts)))
	http.HandleFunc("/react",middleware.AuthMiddleware(middleware.RateLimitMiddleware( controllers.ReactPostController)))
	http.HandleFunc("/create_comment",middleware.AuthMiddleware(middleware.RateLimitMiddleware( controllers.CreateCommentController)))
	http.HandleFunc("/resources/", controllers.CssJsController)
	

}
