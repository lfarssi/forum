package models

import (
	"database/sql"

	"strings"
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
	ReportID      int    `json:"report_id"`
	ReportReason  string `json:"report_reason"`
	Status        string `json:"status"`
	ReportDate    string `json:"report_date"`
	Category      string `json:"category"`
}

type ReportedPost struct {
	ReportID     int      `json:"report_id"`
	PostID       int      `json:"post_id"`
	Title        string   `json:"title"`
	Content      string   `json:"content"`
	Username     string   `json:"username"`
	ReportReason string   `json:"report_reason"`
	Categories   []string `json:"categories"`
	Category     string   `json:"category"` // raw category string (optional)
	ReportDate   string   `json:"report_date"`
	Status       string   `json:"status"`
	CreatedAt    string   `json:"created_at"`
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

func ReportPost(postID, userID, categoryID int) error {
	_, err := Database.Exec(`
		INSERT INTO report (post_id, user_id, report_category_id, comment_id)
		VALUES (?, ?, ?, NULL)
	`, postID, userID, categoryID)
	return err
}

func GetReportedPosts() ([]ReportedPost, error) {
	rows, err := Database.Query(`
		SELECT 
			r.id AS report_id,
			r.post_id,
			p.title,
			p.content,
			u.username,
			cr.name AS report_reason,
			GROUP_CONCAT(c.name) AS categories,
			r.report_date,
			r.status,
			p.creat_at
		FROM report r
		JOIN posts p ON r.post_id = p.id
		JOIN users u ON p.user_id = u.id
		JOIN categorie_report cr ON r.report_category_id = cr.id
		LEFT JOIN post_categorie pc ON p.id = pc.post_id
		LEFT JOIN categories c ON pc.categorie_id = c.id
		WHERE r.comment_id IS NULL
		GROUP BY r.id
		ORDER BY r.report_date DESC;
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reports []ReportedPost

	for rows.Next() {
		var r ReportedPost
		var categories sql.NullString

		err := rows.Scan(
			&r.ReportID,
			&r.PostID,
			&r.Title,
			&r.Content,
			&r.Username,
			&r.ReportReason,
			&categories,
			&r.ReportDate,
			&r.Status,
			&r.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		// Parse categories
		if categories.Valid && categories.String != "" {
			r.Category = categories.String
			cats := strings.Split(categories.String, ",")
			r.Categories = make([]string, len(cats))
			for i, c := range cats {
				r.Categories[i] = strings.TrimSpace(c)
			}
		} else {
			r.Category = "Uncategorized"
			r.Categories = []string{"Uncategorized"}
		}

		reports = append(reports, r)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return reports, nil
}
func DeleteReport(reportID int) error {
	_, err := Database.Exec(`
		DELETE FROM report WHERE id = ?
	`, reportID)
	return err
}

func DeletePost(postID int) error {
	// Start a transaction
	tx, err := Database.Begin()
	if err != nil {
		return err
	}

	// Use defer rollback in case of error before commit
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	// Delete reports for the post
	_, err = tx.Exec(`DELETE FROM report WHERE post_id = ?`, postID)
	if err != nil {
		return err
	}

	// Delete reactions to the post
	_, err = tx.Exec(`DELETE FROM reactPost WHERE post_id = ?`, postID)
	if err != nil {
		return err
	}

	// Delete comment reactions for comments on this post
	_, err = tx.Exec(`
        DELETE FROM reactComment WHERE comment_id IN (
            SELECT id FROM comments WHERE post_id = ?
        )
    `, postID)
	if err != nil {
		return err
	}

	// Delete reports for comments on this post
	_, err = tx.Exec(`
        DELETE FROM report WHERE comment_id IN (
            SELECT id FROM comments WHERE post_id = ?
        )
    `, postID)
	if err != nil {
		return err
	}

	// Delete comments for the post
	_, err = tx.Exec(`DELETE FROM comments WHERE post_id = ?`, postID)
	if err != nil {
		return err
	}

	// Delete post categories relationships
	_, err = tx.Exec(`DELETE FROM post_categorie WHERE post_id = ?`, postID)
	if err != nil {
		return err
	}

	// Finally, delete the post itself
	_, err = tx.Exec(`DELETE FROM posts WHERE id = ?`, postID)
	if err != nil {
		return err
	}

	// Commit the transaction
	err = tx.Commit()
	return err
}
