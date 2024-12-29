package main

import (
	"database/sql"
	"forum/config"
	"forum/routes"
	"log"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("sqlite3", "./database/database.db")
	if err != nil {
		log.Fatal("Error opening database: ", err)
	}
}

func main() {
	defer db.Close()
	config.DatabaseExecution(db)
	config.Seeders(db)
	routes.Router(db)

}
