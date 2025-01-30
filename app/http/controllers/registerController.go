package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"forum/app/models"
	"forum/utils"

	"github.com/gofrs/uuid"
)

func ParseRegister(w http.ResponseWriter, r *http.Request) {
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
		fmt.Println("decoding",err)
		ErrorController(w, r, http.StatusInternalServerError, "")
		return
	}
	if user.UserName == "" || user.Email == "" || user.Password == "" || user.ConfirmationPassword == "" {
		fmt.Println("the fields empty")
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
		fmt.Println("invalid username")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"username": "Invalid username",
		})
		return
	} else if !utils.IsValidEmail(user.Email) {
		fmt.Println("invalid email")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"email": "Invalid email",
		})
		return
	} else if user.Password != user.ConfirmationPassword {
		fmt.Println("invalid confirmation")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{"confirmPassword": "password unmatched"})
		return
	} else if utils.IsValidPassword(user.Password) {
		fmt.Println("invalid password ")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{"password": "Weak password"})
		return
	}

	user.Password = utils.HashPassword(user.Password)
	id, errInsertion := models.Register(user)
	
	if len(errInsertion) > 0 {
		fmt.Println("errinsertion",errInsertion)
		fmt.Println("len errinsertion",len(errInsertion))
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
		ErrorController(w, r, http.StatusInternalServerError, "Cannot Create Sessions")
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

