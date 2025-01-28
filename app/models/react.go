package models

import (
	"database/sql"
)

type React struct {
	PostID    int
	CommentID int
	UserID    int
	Status    string
	Sender    string
}

func InsertReactPost(react React) error {
	react_type, err := ExistReact(react.UserID, react.PostID)
	if err == sql.ErrNoRows {
		query := `
					INSERT INTO reactPost (user_id, post_id, react_type)
				VALUES(?, ?, ?)
				`
		_, err = Database.Exec(query, react.UserID, react.PostID, react.Status)
		if err != nil {
			return err
		}
	} else if err == nil {
		if react_type == react.Status {
			query := `
					DELETE FROM reactPost 
					WHERE user_id=? AND post_id = ?
				`
			_, err := Database.Exec(query, react.UserID, react.PostID)
			if err != nil {
				return err
			}
		} else {
			query := `
				UPDATE reactPost
				SET react_type=? 
				WHERE user_id=? AND post_id=?
			`
			_, err := Database.Exec(query, react.Status, react.UserID, react.PostID)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func InsertReactComment(react React) error {
	react_type, err := ExistReactComment(react.UserID, react.CommentID)
	if err == sql.ErrNoRows {
		query := `
		INSERT INTO reactComment (user_id, comment_id, react_type)
		VALUES(?, ?, ?)
		`
		_, err = Database.Exec(query, react.UserID, react.CommentID, react.Status)
		if err != nil {
			return err
		}
	} else if err == nil {
		if react_type == react.Status {
			query := `
			DELETE FROM reactComment
			WHERE user_id=? AND comment_id=?
			`
			_, err := Database.Exec(query, react.UserID, react.CommentID)
			if err != nil {
				return err
			}
		} else {
			query := `
			UPDATE reactComment
			SET react_type=? 
			WHERE user_id=? AND comment_id=?
		`
			_, err := Database.Exec(query, react.Status, react.UserID, react.CommentID)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func ExistReact(userId, postId int) (string, error) {
	var react_type string
	query := `
		SELECT react_type 
		FROM reactPost
		WHERE user_id = ? AND post_id= ?
	`
	err := Database.QueryRow(query, userId, postId).Scan(&react_type)
	if err != nil {
		return "", err
	}
	return react_type, nil
}
func ExistReactComment(userId, commentId int) (string, error) {
	var react_type string
	query := `
		SELECT react_type 
		FROM reactComment
		WHERE user_id = ? AND comment_id= ?
	`
	err := Database.QueryRow(query, userId, commentId).Scan(&react_type)
	if err != nil {
		return "", err
	}
	return react_type, nil
}