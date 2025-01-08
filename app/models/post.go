package models

import "time"

type Posts struct {
	ID         int       `json:"id"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	Categories []string  `json:"categories"`
	Likes      int       `json:"likes"`
	Dislikes   int       `json:"dislikes"`
	CreatedAt  time.Time `json:"created_at"`
	Comments   []Comment `json:"comments"`
	Username   string    `json:"username"`
}
