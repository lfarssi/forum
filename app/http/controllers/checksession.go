package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"
)

func CheckSession(w http.ResponseWriter, r *http.Request, db *sql.DB) error {
	cookie, err := r.Cookie("session_token")
	if err != nil || cookie.Value == "" {
		fmt.Println("no value in cookie")
		ErrorController(w,r,http.StatusForbidden)
		return err
	}

	var UserId int
	err = db.QueryRow("SELECT user_id FROM sessions WHERE token=? AND expired_at > ?", cookie.Value, time.Now()).Scan(&UserId)
	if err != nil {
		ErrorController(w,r,http.StatusForbidden)
		return err
	}

	return err
}
