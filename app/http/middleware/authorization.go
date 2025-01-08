package middleware

import (
	"fmt"
	"forum/app/http/controllers"
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
