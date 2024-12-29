package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"
)

func CheckSession(w http.ResponseWriter, r *http.Request) (*sql.DB, error) {

	db, err := sql.Open("sqlite3", "/database/database.db")
	if err != nil {
		fmt.Println("error opening database: ", err)
		ErrorController(w, r, http.StatusInternalServerError)
		return nil, err
	}

	defer func() {
		if err2 := db.Close(); err2 != nil {
			fmt.Println("error closing database: ", err2)
			ErrorController(w, r, http.StatusInternalServerError)
		}
	}()

	cookie, err := r.Cookie("session_token")
	if err != nil || cookie.Value == "" {
		return db, nil
	}

	var UserId int
	err = db.QueryRow("SELECT user_id FROM sessions WHERE token=? AND expired_at > ?", cookie.Value, time.Now()).Scan(&UserId)
	if err == nil {
		http.Redirect(w, r, "/home", http.StatusFound)
		return nil, nil
	}

	return db, nil
}
