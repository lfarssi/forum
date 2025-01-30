package middleware

import (
	"net/http"

	"forum/app/http/controllers"
	"forum/app/models"
	"forum/utils"
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
		sessionId := token.Value
		_, err = models.GetSession(sessionId)

		if err != nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
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
