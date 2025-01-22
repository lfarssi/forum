package controllers

import (
	"forum/app/models"
	"forum/utils"
	"net/http"
)

func HomeController(w http.ResponseWriter, r *http.Request) {
	var logedIn bool
	if !utils.IsLoggedIn(r) {
		logedIn = false
	} else {
		logedIn = true
	}
	data := models.Data{
		IsLoggedIn: logedIn,
	}
	if r.Method == "GET" {
		if r.URL.Path != "/" {
			ErrorController(w, r, http.StatusNotFound, http.StatusText(http.StatusNotFound))
			return
		}
		ParseFileController(w, r, "users/index", data)
	} else {
		ErrorController(w, r, http.StatusMethodNotAllowed, "")
		return
	}
}
