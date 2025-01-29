package routes

import (
	"forum/app/http/controllers"
	"forum/app/http/middleware"
	"net/http"
)

func ApiRouter() {

	http.HandleFunc("/singIn", controllers.LoginController)
	http.HandleFunc("/singUp", controllers.RegisterController)
	http.HandleFunc("/logout", controllers.LogoutController)
	http.HandleFunc("/create_post", middleware.AuthMiddleware(middleware.RateLimitMiddleware( controllers.CreatePosts)))
	http.HandleFunc("/react",middleware.AuthMiddleware(middleware.RateLimitMiddleware( controllers.ReactPostController)))
	http.HandleFunc("/create_comment",middleware.AuthMiddleware(middleware.RateLimitMiddleware( controllers.CreatCommentController)))
	http.HandleFunc("/resources/", controllers.CssJsController)
	

}
