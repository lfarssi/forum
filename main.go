package main

import (
	"forum/config"
	"forum/routes"
)

func main() {
	config.DatabaseExecution()
	routes.Router()

}
