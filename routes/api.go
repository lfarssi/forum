package routes

import (
	"forum/app/http/controllers/auth"
	"forum/app/http/controllers/static"
	"net/http"
)

func ApiRouter() {

	// http.HandleFunc("/login", auth.LoginController)
	// http.HandleFunc("/register", auth.RegisterController)
	http.HandleFunc("/logout", auth.LogoutController)
	http.HandleFunc("/resources/", static.CssJsController)

}
