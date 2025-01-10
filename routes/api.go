package routes

import (
	"forum/app/http/controllers"
	"forum/app/http/controllers/static"
	"net/http"
)

func ApiRouter() {

	http.HandleFunc("/singIn", controllers.LoginController)
	http.HandleFunc("/singUp", controllers.RegisterController)
	http.HandleFunc("/logout", controllers.LogoutController)
	http.HandleFunc("/resources/", static.CssJsController)

}
