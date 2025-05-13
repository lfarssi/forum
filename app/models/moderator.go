package models

import "time"

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
