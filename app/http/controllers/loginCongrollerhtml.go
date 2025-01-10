package controllers

import (
	"forum/app/models"
	"forum/utils"
	"net/http"
	"time"

	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
)

func ParseLogin(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/login" {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	} else if r.Method != "GET" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	ParseFileController(w, r, "auth/login", "")
}
func LoginController(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		ErrorController(w, r, http.StatusMethodNotAllowed, "")
		return
	}
	user := models.User{}
	user.UserName = r.FormValue("username")
	user.Password = r.FormValue("password")
	if user.UserName == "" || user.Password == "" { 
		ErrorController(w, r, http.StatusBadRequest, "")
        return
	} else if !utils.IsValidUsername(user.UserName) {
		ErrorController(w, r, http.StatusBadRequest, "Invalid username")
		return
	} 
	// else if !utils.IsValidPassword(user.Password) {
	// 	ErrorController(w, r, http.StatusBadRequest, "Password is weak")
	// 	return
	// }
	id, authErr:= checkAuth(user.UserName, user.Password)
	if authErr != "" {
		ErrorController(w, r, http.StatusInternalServerError, authErr)
		return
	}
	token, err := uuid.NewV4()
	if err != nil {
		ErrorController(w, r, http.StatusInternalServerError, "")
		return
	}
	err = CreateSession(id, token.String(), int(time.Hour)*24)
	if err != nil {
		ErrorController(w, r, http.StatusInternalServerError, "")
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
        Value:    token.String(),
        MaxAge:   int(time.Hour) * 24,
		Path:     "/home",
		HttpOnly: true,
    })
	println(time.Hour)
	http.Redirect(w, r, "/home", http.StatusSeeOther)

}
func checkAuth(userName, password string) (int, string) {
	query := "SELECT id, password FROM users WHERE username = ?"
	statement, err := models.Database.Prepare(query)
	if err!= nil {
        return 0, "Error in the database"
    }
	defer statement.Close()
	var id int
	var hashedPassword string
	err = statement.QueryRow(userName).Scan(&id, &hashedPassword)
	if err!= nil {
        return 0, "User not found"
    }
	if bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))!= nil {
		return 0, "Password Incorrect"
	}
	return id, ""
}