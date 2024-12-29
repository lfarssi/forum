package models



type Comment struct {
	ID        int    `json:"id"`
	PostID    int    `json:"post_id"`
	UserID    int    `json:"user_id"` 
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
	Username  string `json:"username"` 
	Reactions   []string `json:"reactions"` 
}
type Post struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Category  string    `json:"category"`
	Likes     int       `json:"likes"`
	Dislikes  int       `json:"dislikes"`
	CreatedAt string    `json:"created_at"`
	Comments  []Comment `json:"comments"`
	Username  string    `json:"username"` 
}
type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}