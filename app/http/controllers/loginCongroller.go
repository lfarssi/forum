package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"forum/app/models"
	"forum/utils"

	"github.com/gofrs/uuid"
)

func ParseLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	if utils.IsLoggedIn(r) {

		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	ParseFileController(w, r, "auth/login", "")
}

func LoginController(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		ErrorController(w, r, http.StatusMethodNotAllowed, "")
		return
	}
	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		ErrorController(w, r, http.StatusInternalServerError, "")
		return
	}
	if user.UserName == "" || user.Password == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"username": "Username field are required",
			"password": "Password field are required",
		})

		return
	} else if !utils.IsValidUsername(user.UserName) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"username": "Invalid username",
		})
		return
	}
	id, authErr := models.Login(user.UserName, user.Password)
	if len(authErr) > 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"username": authErr["username"],
			"password": authErr["password"],
		})
		return
	}
	token, err := uuid.NewV4()
	if err != nil {
		ErrorController(w, r, http.StatusInternalServerError, "Cannot Generate token")
		return
	}
	err = models.CreateSession(id, token.String(), time.Now().Add((24 * time.Hour)))
	if err != nil {
		ErrorController(w, r, http.StatusInternalServerError, "Cannot Create Sessions")
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    token.String(),
		Expires:  time.Now().Add((24 * time.Hour)),
		HttpOnly: true,
	})
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
