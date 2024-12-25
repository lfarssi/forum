package main

import (
	"forum/app/http/config"
	"forum/routes"
	"os"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		routes.Router()
	} else if len(args) > 1 {
		config.Config(args)
	}
}
