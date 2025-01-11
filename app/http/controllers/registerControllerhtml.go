package controllers

import (
	"encoding/json"
	"forum/app/models"
	"forum/utils"
	"net/http"
	"strings"
	"time"

	"github.com/gofrs/uuid"
)

func ParseRegister(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/register" {
		ErrorController(w, r, http.StatusNotFound, "")
		return
	} else if r.Method != "GET" {
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
	//fmt.Println("user", user)
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
	id, errInsertion := insert(user)
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
	err = CreateSession(id, token.String(), int(time.Now().Hour())*24)

	if err != nil {
		ErrorController(w, r, http.StatusInternalServerError, "")
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    token.String(),
		Path:     "/home",
		HttpOnly: true,
		MaxAge:   int(time.Now().Hour()) * 24,
	})
	http.Redirect(w, r, "/home", http.StatusSeeOther)

}

func insert(user models.User) (int, map[string]string) {
	query := "INSERT INTO users (username, email, password) VALUES (?, ?, ?)"
	stm, err := models.Database.Prepare(query)
	if err != nil {
		return 0, map[string]string{"error": "error preparing query"}
	}
	defer stm.Close()
	res, err := stm.Exec(user.UserName, user.Email, user.Password)
	if err != nil {
		if strings.Contains(err.Error(), "username") {
			return 0, map[string]string{"username": "username already exists"}
		} else if strings.Contains(err.Error(), "email") {
			return 0, map[string]string{"email": "email already exists"}
		}
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, map[string]string{"error": "error getting last id"}
	}
	return int(id), map[string]string{}
}

func CreateSession(id int, token string, expired int) error {
	query := `
	INSERT INTO sessionss (user_id, token, expired_at) 
	VALUES (?, ?, ?) 
	ON CONFLICT DO UPDATE SET token = EXCLUDED.token , date = CURRENT_TIMESTAMP
	`
	stm, err := models.Database.Prepare(query)
	if err != nil {
		return err
	}
	defer stm.Close()
	_, err = stm.Exec(id, token, expired)
	if err != nil {
		return err
	}
	return nil
}
