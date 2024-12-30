package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
)

func AddPostController(w http.ResponseWriter, r *http.Request, db *sql.DB, title string, content string, categories []string,user_id int) error {
	trans, err := db.Begin()
	if err != nil {
		fmt.Println("error in begin transaction: ", err)
		ErrorController(w, r, http.StatusInternalServerError)
		return err
	}

	var postID int
	err = trans.QueryRow("insert into posts (title,content,user_id)values (?,?,?)returning id", title, content, user_id).Scan(&postID)
	if err != nil {
		fmt.Println("error in inserting post: ", err)
		ErrorController(w, r, http.StatusInternalServerError)
		return err
	}
	for _, categoryID := range categories {
		if _, err = trans.Exec("INSERT INTO post_categories (post_id, category_id) VALUES (?, ?)", postID, categoryID); err != nil {
			trans.Rollback()
			ErrorController(w, r, http.StatusInternalServerError)
			return err
		}
	}
	if err = trans.Commit(); err != nil {
		fmt.Println("error in commit transaction: ", err)
		ErrorController(w, r, http.StatusInternalServerError)
		return err
	}

	return err
}
