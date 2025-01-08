package models

import (
	"database/sql"
	"fmt"
	"io/ioutil"

	_ "github.com/mattn/go-sqlite3"

)

var Database *sql.DB

func DatabaseExecution() {
	err := error(nil)
	Database, err = sql.Open("sqlite3", "./database/database.db")
	if err != nil {
		fmt.Println(" failed to open database: ", err)
		return
	}
	//defer Database.Close()

	// Read the schema SQL file
	schema, err := ioutil.ReadFile("./database/schema/schema.sql")
	if err != nil {
		fmt.Println(" failed to read schema file: ", err)
		return
	}

	// Execute the SQL commands in the schema file
	_, err = Database.Exec(string(schema))
	if err != nil {
		fmt.Println(" failed to execute schema:", err)
		return
	}
}
