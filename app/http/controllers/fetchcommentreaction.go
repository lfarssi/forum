package controllers

import (
	"database/sql"
	"fmt"
)

// Fetches reactions for comments associated with a specific post
func fetchCommentReactions(db *sql.DB, postID int) (map[int][]string, error) {
	reactions := make(map[int][]string)

	commentQuery := `
	SELECT 
		rc.comment_id,
		rc.react_type
	FROM reactComment rc
	JOIN comments c ON rc.comment_id = c.id
	WHERE c.post_id = ?
	`

	rows, err := db.Query(commentQuery, postID)
	if err != nil {
		fmt.Println("error querying comment reactions: ", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var commentID int
		var reactType string
		if err := rows.Scan(&commentID, &reactType); err != nil {
			fmt.Println("error scanning comment reactions: ", err)
			return nil, err
		}
		reactions[commentID] = append(reactions[commentID], reactType)
	}

	return reactions, nil
}
