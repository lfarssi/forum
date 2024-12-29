package config

import (
	"database/sql"
	"fmt"
	"io/ioutil"
)

func Seeders(db *sql.DB) {
	schema, err := ioutil.ReadFile("./database/schema/seeders.sql")
	if err != nil {
		fmt.Println(" failed to read schema file seeders: ", err)
		return
	}
	_, err = db.Exec(string(schema))
	if err != nil {
		fmt.Println(" failed to execute schema seeders :", err)
		return
	}
}
