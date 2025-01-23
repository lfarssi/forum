package models

import "time"

type Comment struct {
	ID        int       `json:"id"`
	PostID    int       `json:"post_id"`
	UserID    int       `json:"user_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	Username  string    `json:"username"`
	Likes     int       `json:"likes"`
	Dislikes  int       `json:"dislikes"`
}

func CreatComment(comment Comment) error {
	qwiri := `
		INSERT INTO comments
		(content , user_id, date_creation, post_id)
		VALUES (?,?,?,?);
	
	`
	_, err := Database.Exec(qwiri, comment.Content, comment.UserID, comment.CreatedAt, comment.PostID)
	if err != nil {
		return err
	}
	return nil
}
