package models

import (
	"time"
)

type Session struct {
	ID int
	UserId    int
	Token     string
	ExpiredAt time.Time
}

func CreateSession(id int, token string, expired time.Time) error {
	query := `
	INSERT INTO sessionss (user_id, token, expired_at) 
	VALUES (?, ?, ?) 
	ON CONFLICT DO UPDATE SET token = EXCLUDED.token , expired_at = CURRENT_TIMESTAMP
	`
	stm, err := Database.Prepare(query)
	if err != nil {
		return err
	}
	defer stm.Close()
	_, err = stm.Exec(id, token, expired)
	if err != nil {
		return err
	}
	return nil
}

func GetSession(id string) (*Session, error) {
	query := `SELECT id, user_id, token, expired_at FROM sessionss WHERE token = ?`
	stm, err := Database.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stm.Close()
	var session Session
	err = stm.QueryRow(id).Scan(&session.ID, &session.UserId, &session.Token, &session.ExpiredAt)
	if err != nil {
		return nil , err
	}
	if time.Now().After(session.ExpiredAt) {
		return nil, err
	}
	return &session, nil
}