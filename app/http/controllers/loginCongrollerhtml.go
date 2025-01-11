package controllers

import (
	"encoding/json"
	"fmt"
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
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		ErrorController(w, r, http.StatusInternalServerError, "")
		return
	}
	if user.UserName == "" || user.Password == "" {
		fmt.Println("&", r.FormValue("username"), "=>", user.Password)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"username": "Username field are required",
			"password": "Password field are required",
		})

		return
	} else if !utils.IsValidUsername(user.UserName) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Println("ùù", user.UserName, "=>", user.Password)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"username": "Invalid username",
		})
		return
	}
	// else if !utils.IsValidPassword(user.Password) {
	// 	w.Header().Set("Content-Type", "application/json")
	//	w.WriteHeader(http.StatusBadRequest)
	//	json.NewEncoder(w).Encode(map[string]interface{}{"password": "Weak password",})
	// 	return
	// }
	id, authErr := checkAuth(user.UserName, user.Password)
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
		ErrorController(w, r, http.StatusInternalServerError, "")
		return
	}
	err = CreateSession(id, token.String(), int(time.Now().Hour())*24)
	if err != nil {
		ErrorController(w, r, http.StatusInternalServerError, "")
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    token.String(),
		MaxAge:   int(time.Now().Hour()) * 24,
		Path:     "/home",
		HttpOnly: true,
	})
	println(time.Hour)
	http.Redirect(w, r, "/home", http.StatusSeeOther)

}
func checkAuth(userName, password string) (int, map[string]string) {
	query := "SELECT id, password FROM users WHERE username = ?"
	statement, err := models.Database.Prepare(query)
	if err != nil {
		return 0, map[string]string{"error": "database error"}
	}
	defer statement.Close()
	var id int
	var hashedPassword string
	err = statement.QueryRow(userName).Scan(&id, &hashedPassword)
	if err != nil {
		return 0, map[string]string{"username": "Username not found"}
	}
	if bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) != nil {
		return 0, map[string]string{"password": "Password Incorrect"}
	}
	return id, map[string]string{}
}
