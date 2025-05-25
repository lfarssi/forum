package models

import (
	"database/sql"
	"time"
)

type ModReq struct {
	UserID       string `json:"user_id"`
	Request_date string `json:"request_date"`
	Status       string `json:"status"`
	Review_date  string `json:"review_date"`
	Reason       string
}

func AddModRequest(reason string, userId int) error {
	query := "INSERT INTO moderator_requests (user_id, reason, request_date, status) VALUES (?, ?, ?, ?)"
	stmt, err := Database.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(userId, reason, time.Now().Format("2006-01-02 15:04:05"), "pending")
	if err != nil {
		return err
	}
	return nil
}

func GetRequestInfo(userID int) (string, error) {
	query := `
	SELECT status  
	FROM moderator_requests 
	WHERE user_id = ?
	ORDER BY request_date DESC
	LIMIT 1
	`
	var status string
	err := Database.QueryRow(query, userID).Scan(&status)
	if err != nil {
		if err == sql.ErrNoRows {
			return "No Request", nil
		}
		return "", err
	}
	return status, nil
}

func GetAllModRequests() ([]ModReq, error) {
	query := `
	SELECT user_id, reason, request_date, status
	FROM moderator_requests
	ORDER BY request_date DESC
	`
	rows, err := Database.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var requests []ModReq
	for rows.Next() {
		var req ModReq
		err := rows.Scan(&req.UserID, &req.Reason, &req.Request_date, &req.Status)
		if err != nil {
			return nil, err
		}
		requests = append(requests, req)
	}
	return requests, nil
}
