package routes

import (
	"forum/app/http/controllers"
	"net/http"
)

func ApiRouter() {

	http.HandleFunc("/singIn", controllers.LoginController)
	http.HandleFunc("/singUp", controllers.RegisterController)
	http.HandleFunc("/logout", controllers.LogoutController)
	http.HandleFunc("/create_post", controllers.CreatePosts)
	http.HandleFunc("/resources/", controllers.CssJsController)

}
