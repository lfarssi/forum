package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"forum/app/models"
	"forum/utils"

	"github.com/gofrs/uuid"
)

func ParseRegister(w http.ResponseWriter, r *http.Request) {
	if utils.IsLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	if r.Method != "GET" {
		ErrorController(w, r, http.StatusMethodNotAllowed, "")
		return
	}
	ParseFileController(w, r, "auth/register", "")
}

func RegisterController(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		ErrorController(w, r, http.StatusMethodNotAllowed, "")
		return
	}
	user := models.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		ErrorController(w, r, http.StatusInternalServerError, "")
		return
	}
	// fmt.Println("user", user)
	if user.UserName == "" || user.Email == "" || user.Password == "" || user.ConfirmationPassword == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"username":        "Username field are required",
			"email":           "Email field are required",
			"password":        "Password field are required",
			"confirmPassword": "Confirmation Password field are required",
		})
		return
	} else if !utils.IsValidUsername(user.UserName) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"username": "Invalid username",
		})
		return
	} else if !utils.IsValidEmail(user.Email) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"email": "Invalid email",
		})
		return
	} else if user.Password != user.ConfirmationPassword {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{"confirmPassword": "password unmatched"})
		return
	} else if utils.IsValidPassword(user.Password) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{"password": "Weak password"})
		return
	}

	user.Password = utils.HashPassword(user.Password)
	id, errInsertion := models.Register(user)
	if len(errInsertion) > 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error":    errInsertion["error"],
			"username": errInsertion["username"],
			"email":    errInsertion["email"],
		})
		return
	}

	token, err := uuid.NewV4()
	if err != nil {
		ErrorController(w, r, http.StatusInternalServerError, "")
		return
	}
	err = models.CreateSession(id, token.String(), time.Now().Add((24 * time.Hour)))
	if err != nil {
		ErrorController(w, r, http.StatusInternalServerError, "")
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    token.String(),
		HttpOnly: true,
		Expires:  time.Now().Add((24 * time.Hour)),
	})
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

