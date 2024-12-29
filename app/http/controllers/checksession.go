package controllers

import (
	"database/sql"
	"net/http"
	"time"
)

func CheckSession(w http.ResponseWriter, r *http.Request,db *sql.DB) error {
	cookie, err := r.Cookie("session_token")
	if err != nil || cookie.Value == "" {
		return nil
	}

	var UserId int
	err = db.QueryRow("SELECT user_id FROM sessions WHERE token=? AND expired_at > ?", cookie.Value, time.Now()).Scan(&UserId)
	if err == nil {
		http.Redirect(w, r, "/home", http.StatusFound)
		return nil
	}

	return err
}
