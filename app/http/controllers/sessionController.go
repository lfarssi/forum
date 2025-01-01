package controllers

import (
	"database/sql"
	"fmt"
	"time"
)

func SessionController(db *sql.DB) {
	expireAt := time.Now()
	res, err := db.Exec("delete from sessions where expires_at < ?", expireAt)
	fmt.Println("res: ", res)
	if err != nil {
		fmt.Println("error deleting expired sessions: ", err)
		return
	}else{
		rowsAffected, _ := res.RowsAffected()
        fmt.Printf("Deleted %d expired sessions\n", rowsAffected)
	}
}
