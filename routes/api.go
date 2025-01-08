package routes

import (
	"forum/app/http/controllers/auth"
	"forum/app/http/controllers/static"
	"net/http"
)

func ApiRouter() {

	 http.HandleFunc("/singIn", auth.LoginController)
	 http.HandleFunc("/singUp", auth.RegisterController)
	http.HandleFunc("/logout", auth.LogoutController)
	http.HandleFunc("/resources/", static.CssJsController)

}
