package main

import (
	"fmt"

	"forum/app/models"
	"forum/routes"
	"net/http"
)


func main() {
	models.DatabaseExecution()
	defer models.CloseDatabase()
	routes.WebRouter()
	routes.ApiRouter()
	fmt.Println("Server is running on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("err starting the server : ", err)
		return
	}
}
