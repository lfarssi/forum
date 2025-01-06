package controllers

import (
	"net/http"
)

func RegisterController(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/register" {
		ErrorController(w, r, http.StatusNotFound)
		return
	} else if r.Method != "GET" {
		ErrorController(w, r, http.StatusMethodNotAllowed)
		return
	}
	ParseFileController(w, r, "auth/register", "")

}
