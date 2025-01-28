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
	http.HandleFunc("/create_post", middleware.AuthMiddleware( controllers.CreatePosts))
	http.HandleFunc("/likepost", controllers.ReactPostController)
	http.HandleFunc("/likecomment", controllers.ReactPostController)
	http.HandleFunc("/create_comment",middleware.AuthMiddleware( controllers.CreatCommentController))
	http.HandleFunc("/resources/", controllers.CssJsController)

}
