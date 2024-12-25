package config

import (
	"database/sql"
	"fmt"
	"io/ioutil"
)

func Migrate(schemaFile string, databaseFile string) error {
	db, err := sql.Open("sqlite3", databaseFile)
	if err != nil {
		return fmt.Errorf("failed to open database: %v", err)
	}
	defer db.Close()
	schema, err := ioutil.ReadFile(schemaFile)
	if err != nil {
		return fmt.Errorf("failed to read schema file: %v", err)
	}
	_, err = db.Exec(string(schema))
	if err != nil {
		return fmt.Errorf("failed to execute schema: %v", err)
	}

	return nil

}
