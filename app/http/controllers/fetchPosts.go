package controllers
import (
	"database/sql"
	"fmt"
	"forum/app/models"
	"net/http"
)
func fetchPosts(w http.ResponseWriter, r *http.Request, db *sql.DB) ([]models.Post, error) {
	posts := []models.Post{}
	query := `
	SELECT 
	    p.id,
	    p.title,
	    p.content,
	    c.name AS category,
	    COALESCE(SUM(CASE WHEN pr.react_type='like' THEN 1 ELSE 0 END), 0) AS likes,
	    COALESCE(SUM(CASE WHEN pr.react_type='dislike' THEN 1 ELSE 0 END), 0) AS dislikes,
	    p.created_at
	FROM posts p
	LEFT JOIN post_categories pc ON p.id = pc.post_id
	LEFT JOIN categories c ON pc.category_id = c.id
	LEFT JOIN reactPost pr ON p.id = pr.post_id
	GROUP BY p.id
	ORDER BY p.created_at DESC
    `

	rows, err := db.Query(query)
	if err != nil {
		fmt.Println("error querying posts: ", err)
		ErrorController(w, r, http.StatusInternalServerError)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var post models.Post
		if err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.Category,
			&post.Likes, &post.Dislikes, &post.CreatedAt); err != nil {
			fmt.Println("error scanning posts: ", err)
			ErrorController(w, r, http.StatusInternalServerError)
			return nil, err
		}
        comments := []models.Comment{}
        commentQuery := `
        SELECT 
            id,
            content,
            created_at,
            user_id,
            (SELECT username FROM users WHERE id=user_id) AS username
        FROM comments 
        WHERE post_id = ?
        ORDER BY created_at ASC
        `
        
        commentRow, err := db.Query(commentQuery, post.ID)
        if err != nil {
            fmt.Println("error querying comments: ", err)
            ErrorController(w, r, http.StatusInternalServerError)
            return nil, err
        }
        defer commentRow.Close()

        for commentRow.Next() {
            var comment models.Comment
            if err := commentRow.Scan(&comment.ID, &comment.Content,
                &comment.CreatedAt, &comment.UserID); err != nil {
                fmt.Println("error scanning comments: ", err)
                ErrorController(w, r, http.StatusInternalServerError)
                return nil, err
            }
            comments = append(comments, comment)
        }

        reactionsMap, err := fetchCommentReactions(db, post.ID)
        if err != nil {
            return nil, err 
        }

        for i := range comments {
            if reactionList, exists := reactionsMap[comments[i].ID]; exists {
            
                comments[i].Reactions = reactionList 
            }
        }
        post.Comments = comments
        posts = append(posts, post)
    }

	return posts, nil
}
