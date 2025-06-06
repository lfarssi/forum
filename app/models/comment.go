package models

import (
	"time"
)

type Comment struct {
	ID         int       `json:"id"`
	PostID     int       `json:"post_id"`
	UserID     int       `json:"user_id"`
	Content    string    `json:"content"`
	CreatedAt  time.Time `json:"created_at"`
	Username   string    `json:"username"`
	Likes      int       `json:"likes"`
	Dislikes   int       `json:"dislikes"`
	IsLiked    bool
	IsDisliked bool
}

func CreateComment(comment Comment) error {
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

func GetComments(postid int) ([]Comment, error) {
	qwiri := `
	SELECT c.id, c.content, c.user_id, c.date_creation,  c.post_id , u.username
	FROM comments c
	INNER JOIN  users u 
	ON u.id = c.user_id 
	WHERE post_id = ?
	 ORDER BY date_creation DESC;	
	`
	rows, err := Database.Query(qwiri, postid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var comments []Comment
	for rows.Next() {

		var c Comment
		if err := rows.Scan(&c.ID, &c.Content, &c.UserID, &c.CreatedAt, &c.PostID, &c.Username); err != nil {
			return nil, err
		}
		comments = append(comments, c)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return comments, nil
}

func CommentExists(id int) (bool, error) {
	var exists bool
	err := Database.QueryRow("SELECT EXISTS(SELECT 1 FROM comments WHERE id = ?)", id).Scan(&exists)
	return exists, err
}

func DeleteComment(id int) error {
	_, err := Database.Exec("DELETE FROM comments WHERE id = ?", id)
	return err
}
