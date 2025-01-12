package controllers

import (
	"forum/app/models"
	"net/http"
)

func LogoutController(w http.ResponseWriter, r *http.Request) {
	cookies, err := r.Cookie("token")
	if err != nil || cookies.Value == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	delete(cookies.Value)
	http.SetCookie(w, &http.Cookie{
		Value:  "",
		Name:   "token",
		MaxAge: -1,
	})
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
func delete(token string) {
	query := "DELETE FROM sessions WHERE token = ?"
	_, err := models.Database.Exec(query, token)
	if err != nil {
		return
	}
}
