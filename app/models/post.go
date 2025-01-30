package models

import (
	"time"
)

type Posts struct {
	ID            int `json:"id"`
	UserID        int
	Title         string    `json:"title"`
	Content       string    `json:"content"`
	Categories    []string  `json:"categories"`
	Likes         int       `json:"likes"`
	Dislikes      int       `json:"dislikes"`
	CreatedAt     string    `json:"created_at"`
	Comments      []Comment `json:"comments"`
	CommentsCount int
	Username      string `json:"username"`
	IsLiked       bool
	IsDisliked    bool
}

func CreatePost(title string, content string, categories []string, userId int) (int, error) {
	query := "INSERT INTO posts ( title, user_id, content, creat_at) VALUES ( ?, ?, ?, ?)"
	stm1, err := Database.Prepare(query)
	if err != nil {
		return 0, err
	}
	defer stm1.Close()

	res, err := stm1.Exec(title, userId, content, time.Now().Format("2006-01-02 15:04:05"))
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}
func LikedPost(userID int) ([]Posts, error) {
	query := `
	SELECT p.id , p.title,p.content,p.creat_at ,u.username , GROUP_CONCAT(DISTINCT c.name) AS categories
	FROM posts p 
	INNER JOIN users u ON u.id=p.user_id
	INNER JOIN reactPost r ON p.id = r.post_id
	  INNER JOIN post_categorie pc ON p.id = pc.post_id
    INNER JOIN categories c ON pc.categorie_id = c.id
	WHERE react_type='like' AND r.user_id=?
	GROUP BY p.id
	ORDER BY p.creat_at DESC
	`
	rows, err := Database.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var LikedPost []Posts
	for rows.Next() {
		var post Posts
		var categorie string
		var CreatedAt time.Time
		err = rows.Scan(&post.ID, &post.Title, &post.Content, &CreatedAt, &post.Username, &categorie)
		if err != nil {
			return nil, err
		}
		post.Categories = append(post.Categories, categorie)
		post.CreatedAt = CreatedAt.Format("2006-01-02 15:04:05")
		LikedPost = append(LikedPost, post)
	}
	return LikedPost, nil

}

func CreatedPost(iduser int) ([]Posts, error) {
	query := `
	SELECT p.id , p.title,p.content,p.creat_at ,u.username , GROUP_CONCAT(DISTINCT c.name) AS categories
	FROM posts p 
	INNER JOIN users u ON u.id=p.user_id
	  INNER JOIN post_categorie pc ON p.id = pc.post_id
    INNER JOIN categories c ON pc.categorie_id = c.id
	WHERE p.user_id=?
	GROUP BY p.id
	ORDER BY p.creat_at DESC
	`
	rows, err := Database.Query(query, iduser)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var createdPost []Posts
	for rows.Next() {
		var post Posts
		var categorie string
		var CreatedAt time.Time
		err = rows.Scan(&post.ID, &post.Title, &post.Content, &CreatedAt, &post.Username, &categorie)
		if err != nil {
			return nil, err
		}
		post.Categories = append(post.Categories, categorie)
		post.CreatedAt = CreatedAt.Format("2006-01-02 15:04:05")
		createdPost = append(createdPost, post)
	}
	return createdPost, nil
}

func GetPosts() ([]Posts, error) {
	query := `
    SELECT p.id,p.user_id, p.title, p.content, GROUP_CONCAT(c.name) AS categories, p.creat_at, u.username
    FROM posts p
    INNER JOIN users u ON p.user_id = u.id
    INNER JOIN post_categorie pc ON p.id = pc.post_id
    INNER JOIN categories c ON pc.categorie_id = c.id
    GROUP BY p.id
    ORDER BY p.creat_at DESC;
    `
	rows, err := Database.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []Posts
	for rows.Next() {
		var post Posts
		var CreatAt time.Time
		var categorie string
		err = rows.Scan(&post.ID, &post.UserID, &post.Title, &post.Content, &categorie, &CreatAt, &post.Username)
		if err != nil {
			return nil, err
		}
		post.Categories = append(post.Categories, categorie)
		post.CreatedAt = CreatAt.Format("2006-01-02 15:04:05")
		posts = append(posts, post)
	}
	return posts, nil
}

func GetPostsByCategory(idCategorie int) ([]Posts, error) {
	query := `
	SELECT   p.id, p.title, p.content, c.name, p.creat_at, u.username
	FROM posts p
	INNER JOIN users u ON p.user_id = u.id
	INNER JOIN post_categorie pc ON p.id = pc.post_id
	INNER JOIN categories c ON pc.categorie_id = c.id
	WHERE pc.categorie_id =?
	ORDER BY p.creat_at DESC;
	`
	rows, err := Database.Query(query, idCategorie)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var posts []Posts
	tempPosts := make(map[int]*Posts)
	for rows.Next() {
		var post Posts
		var categorie string
		err = rows.Scan(&post.ID, &post.Title, &post.Content, &categorie, &post.CreatedAt, &post.Username)
		if err != nil {
			return nil, err
		}
		if temposts, ok := tempPosts[post.ID]; ok {
			temposts.Categories = append(temposts.Categories, categorie)
		} else {
			post.Categories = CorrectCategories(post.ID)

			tempPosts[post.ID] = &post
		}
		
	}
	for _, post := range tempPosts {
		posts = append(posts, *post)
	}
	return posts, nil

}
