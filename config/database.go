package config

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	_"github.com/mattn/go-sqlite3"
)

func DatabaseExecution(db *sql.DB) {
	
	schema, err := ioutil.ReadFile("./database/schema/schema.sql")
	if err != nil {
		fmt.Println(" failed to read schema file database: ", err)
		return
	}
	_, err = db.Exec(string(schema))
	if err != nil {
		fmt.Println(" failed to execute schema database:", err)
		return
	}
	// print("exce succes ")
}
