package controllers

import (
	"database/sql"
	"fmt"
	"forum/app/models"
	"net/http"
	"strings"
)

func FilterPostsController(w http.ResponseWriter, r *http.Request, db *sql.DB, category string) {
	posts := []models.Post{}

	var query string
	if category == "All" || category == "" {
		query = `
        SELECT 
            p.id,
            p.title,
            p.content,
            GROUP_CONCAT(c.name) AS categories,
            COALESCE(SUM(CASE WHEN pr.react_type='like' THEN 1 ELSE 0 END), 0) AS likes,
            COALESCE(SUM(CASE WHEN pr.react_type='dislike' THEN 1 ELSE 0 END), 0) AS dislikes,
            p.created_at,
            u.username 
        FROM posts p
        LEFT JOIN post_categories pc ON p.id = pc.post_id
        LEFT JOIN categories c ON pc.category_id = c.id
        LEFT JOIN reactPost pr ON p.id = pr.post_id
        LEFT JOIN users u ON p.user_id = u.id
        GROUP BY p.id
        ORDER BY p.created_at DESC
        `
	} else {
		query = `
        SELECT 
            p.id,
            p.title,
            p.content,
            GROUP_CONCAT(c.name) AS categories,
            COALESCE(SUM(CASE WHEN pr.react_type='like' THEN 1 ELSE 0 END), 0) AS likes,
            COALESCE(SUM(CASE WHEN pr.react_type='dislike' THEN 1 ELSE 0 END), 0) AS dislikes,
            p.created_at,
            u.username 
        FROM posts p
        INNER JOIN post_categories pc ON p.id = pc.post_id
        INNER JOIN categories c ON pc.category_id = c.id AND c.name = ?
        LEFT JOIN reactPost pr ON p.id = pr.post_id
        LEFT JOIN users u ON p.user_id = u.id
        GROUP BY p.id
        ORDER BY p.created_at DESC
        `
	}

	var rows *sql.Rows
	var err error

	if category == "All" || category == "" {
		rows, err = db.Query(query)
	} else {
		rows, err = db.Query(query, category)
	}

	if err != nil {
		fmt.Println("error querying posts: ", err)
		ErrorController(w, r, http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var post models.Post
		var categoriesString string // Temporary variable for categories

		if err := rows.Scan(&post.ID, &post.Title, &post.Content, &categoriesString,
			&post.Likes, &post.Dislikes, &post.CreatedAt, &post.Username); err != nil {
			fmt.Println("error scanning posts: ", err)
			ErrorController(w, r, http.StatusInternalServerError)
			return
		}

		post.Categories = strings.Split(categoriesString, ",")
		posts = append(posts, post)
	}

	categories, err := fetchCategories(w, r, db)
	if err != nil {
		fmt.Println("error fetching categories: ", err)
		return
	}

	data := struct {
		Posts      []models.Post
		Categories []models.Category
	}{
		Posts:      posts,
		Categories: categories,
	}

    // Render the filtered posts in a separate template
    ParseFileController(w, r, "user/filterdposts", data)
}
