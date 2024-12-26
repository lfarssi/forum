package controllers

import (
	"forum/config"
	"net/http"
)
func HomeController(w http.ResponseWriter, r *http.Request) {
	config.DatabaseExecution()
	

}
