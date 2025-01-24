package models

import (
	"fmt"
	"time"
)

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
		fmt.Println("awal error")
		return nil, err
	}
	defer rows.Close()
	var comments []Comment
	for rows.Next() {

		var c Comment
		if err := rows.Scan(&c.ID,&c.Content, &c.UserID,  &c.CreatedAt, &c.PostID, &c.Username); err != nil {
			fmt.Println("error 2")
			return nil, err
		}
		comments = append(comments, c)
	}
	if err := rows.Err(); err != nil {
		fmt.Println("eroor 3")
		return nil, err
	}
	return comments, nil
}
