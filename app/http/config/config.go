package config

import (
	"fmt"
	"log"
	"os"
)

func Config(args []string) {
	if len(args) > 2 {
		fmt.Println("Invalid arguments")
		return
	}
	cmd := args[1]
	fmt.Println(cmd)
	if cmd == "--migrate" || cmd == "-m" {
		if _, err := os.Stat("./http/database/database.db"); err == nil {
			fmt.Println("The database already exists. Re-migrating will delete all data.")
			fmt.Print("Are you sure you want to proceed? (yes/no): ")
			var confirmation string
			fmt.Scanln(&confirmation)
			if confirmation != "yes" {
				fmt.Println("Migration canceled.")
				return
			}
			err = os.Remove("./back-end/database/database.db")
			if err != nil {
				log.Fatalf("Failed to delete the existing database: %v", err)
			}
			fmt.Println("Existing database deleted.")
		}
		err := Migrate("./back-end/database/db.sql", "./back-end/database/database.db")
		if err != nil {
			log.Fatalf("Migration failed: %v", err)
		}
		fmt.Println("Database migration completed successfully.")
		return
	}else if cmd == "--seed" {
		
		err := Seeders("./back-end/database/database.db", "./back-end/database/seeder.sql")
		if err != nil {
			log.Fatalf("Seeder failed: %v", err)
		}
		fmt.Println("Seeder completed successfully.")
	} else {
		fmt.Println("Invalid command try : --migrate or --seed")
		return
	}



}