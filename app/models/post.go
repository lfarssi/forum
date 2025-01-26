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
	CommentsCount   int 
	Username   string    `json:"username"`
}


func CreatePost(title string, content string, categories []string, userId int) (int, error){
	query := "INSERT INTO posts ( title, user_id, content, creat_at) VALUES ( ?, ?, ?, ?)"
	stm1, err := Database.Prepare(query)
	if err != nil {
		return 0, err
	}
	defer stm1.Close()

	res, err := stm1.Exec(title, userId, content, time.Now())
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err!= nil {
        return 0, err
    }
	return int(id), nil
}

func GetPosts() ([]Posts, error) {
	query := `
    SELECT p.id, p.title, p.content, GROUP_CONCAT(c.name) AS categories, p.creat_at, u.username
    FROM posts p
    INNER JOIN users u ON p.user_id = u.id
    INNER JOIN post_categorie pc ON p.id = pc.post_id
    INNER JOIN categories c ON pc.categorie_id = c.id
    GROUP BY p.id
    ORDER BY p.creat_at DESC;
    `
    rows, err := Database.Query(query)
    if err!= nil {
        return nil, err
    }
    defer rows.Close()

    var posts []Posts
	for rows.Next() {
		var post Posts
        var categories []string
		var categorie string
        err = rows.Scan(&post.ID, &post.Title, &post.Content, &categorie, &post.CreatedAt, &post.Username)
        if err!= nil {
            return nil, err
        }
		categories = append(categories, categorie)
        post.Categories = categories
        posts = append(posts, post)
	}
	return posts, nil
}

func GetPostsByCategory(idCategorie int) ([]Posts, error)  {
	query := `
	SELECT p.id, p.title, p.content, GROUP_CONCAT(c.name) AS categories, p.creat_at, u.username
	FROM posts p
	INNER JOIN users u ON p.user_id = u.id
	INNER JOIN post_categorie pc ON p.id = pc.post_id
	INNER JOIN categories c ON pc.categorie_id = c.id
	WHERE pc.categorie_id =?`
	rows, err := Database.Query(query, idCategorie)
	if err!= nil {
        return nil, err
    }
	defer rows.Close()
	var posts []Posts
	for rows.Next() {
		var post Posts
        var categories []string
        var categorie string
        err = rows.Scan(&post.ID, &post.Title, &post.Content, &categorie, &post.CreatedAt, &post.Username)
        if err!= nil {
            return nil, err
        }
        categories = append(categories, categorie)
        post.Categories = categories
        posts = append(posts, post)
	}
	return posts, nil

}