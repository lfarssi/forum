package config

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	_"github.com/mattn/go-sqlite3"
)

func DatabaseExecution() {
	db, err := sql.Open("sqlite3", "./database/database.db")
	if err != nil {
		fmt.Println(" failed to open database: ", err)
		return
	}
	defer db.Close()

	// Read the schema SQL file
	schema, err := ioutil.ReadFile("./database/schema/schema.sql")
	if err != nil {
		fmt.Println(" failed to read schema file: ", err)
		return
	}

	// Execute the SQL commands in the schema file
	_, err = db.Exec(string(schema))
	if err != nil {
		fmt.Println(" failed to execute schema:", err)
		return
	}
	print("exce succes ")
}
