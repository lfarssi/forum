package controllers

import (
	"net/http"
)

func HomeController(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		ParseFileController(w, r, "guests/index","")
		return
	} else if r.URL.Path == "/home" {
		ParseFileController(w, r, "users/index" , "")
		
	} else {
		ErrorController(w, r, http.StatusNotFound, "")
	}
}