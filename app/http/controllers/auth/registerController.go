package auth

import (
	"encoding/json"
	"fmt"
	"forum/app/http/controllers"
	"forum/app/models"
	"forum/utils"
	"net/http"
)

func RegisterController(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		controllers.ErrorController(w, r, http.StatusMethodNotAllowed)
		return
	}
	user := models.User{}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		controllers.ErrorController(w, r, http.StatusBadRequest)
		return
	}
	// username := r.URL.Query().Get("username")
	// email := r.URL.Query().Get("email")
	// password := r.URL.Query().Get("password")
	// confimationPassword := r.URL.Query().Get("confirmationPassword")
	if user.UserName == "" || user.Email == "" || user.Password == "" || !utils.IsValidEmail(user.Email) || !utils.IsValidUsername(user.UserName) || user.Password != user.ConfirmationPassword {
		controllers.ErrorController(w, r, http.StatusBadRequest)
		return
	}
	query := "INSERT INTO users (username, email, password) VALUES (?, ?, ?)"
	fmt.Println(query)
}
