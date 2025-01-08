package auth

import (
	"forum/app/http/controllers"
	"forum/app/models"
	"forum/utils"
	"net/http"
	"strings"
	"time"

	"github.com/gofrs/uuid"
)

func RegisterController(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		controllers.ErrorController(w, r, http.StatusMethodNotAllowed, "")
		return
	}
	user := models.User{}
	user.UserName = r.FormValue("username")
	user.Email = r.FormValue("email")
	user.Password = r.FormValue("password")
	user.ConfirmationPassword = r.FormValue("confirmationPassword")
	if user.UserName == "" || user.Email == "" || user.Password == "" {
		controllers.ErrorController(w, r, http.StatusBadRequest, "")
		return
	} else if !utils.IsValidUsername(user.UserName) {
		controllers.ErrorController(w, r, http.StatusBadRequest, "Invalid username")
		return
	} else if !utils.IsValidEmail(user.Email) {
		controllers.ErrorController(w, r, http.StatusBadRequest, "Invalid email")
		return
	} else if user.Password != user.ConfirmationPassword {
		controllers.ErrorController(w, r, http.StatusBadRequest, "Password and confirmation password do not match")
		return
	} else if utils.IsValidPassword(user.Password) {
		controllers.ErrorController(w, r, http.StatusBadRequest, "Password is weak")
		return
	}

	user.Password = utils.HashPassword(user.Password)
	id, errInsertion := insert(user)
	if errInsertion != "" {
		controllers.ErrorController(w, r, http.StatusInternalServerError, errInsertion)
		return
	}

	token, err := uuid.NewV4()
	if err != nil {
		controllers.ErrorController(w, r, http.StatusInternalServerError, "")
		return
	}
	err = createSession(id, token.String(), int(time.Hour)*24)

	if err != nil {
		controllers.ErrorController(w, r, http.StatusInternalServerError, "")
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    token.String(),
		Path:     "/",
		HttpOnly: true,
		MaxAge:   int(time.Hour) * 24,
	})
	http.Redirect(w, r, "/home", http.StatusSeeOther)

}

func insert(user models.User) (int, string) {
	query := "INSERT INTO users (username, email, password) VALUES (?, ?, ?)"
	stm, err := models.Database.Prepare(query)
	if err != nil {
		return 0, "error preparing statement"
	}
	defer stm.Close()
	res, err := stm.Exec(user.UserName, user.Email, user.Password)
	if err != nil {
		if strings.Contains(err.Error(), "username") {
			return 0, "username already taken"
		} else if strings.Contains(err.Error(), "email") {
			return 0, "email already exists"
		}
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, "error getting last insert id"
	}
	return int(id), ""
}

func createSession(id int, token string, expired int) error {
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
