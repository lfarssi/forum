package middleware

import (
	"fmt"
	"forum/app/http/controllers"
	"forum/utils"
	"net/http"
)

func AuthMiddleware(auth http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		defer func() {
			if err := recover(); err != nil {
				controllers.ErrorController(w, r, http.StatusInternalServerError, "")
				return
			}
		}()
		token, err := r.Cookie("token")
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		fmt.Println(token)
		auth(w, r)
	}
}
func AlreadyLoggedIn(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if utils.IsLoggedIn(r) {
			http.Redirect(w, r, "/", http.StatusSeeOther)

			return
		}
		next(w, r)
	}
}
