package controllers

import "net/http"

func PostController(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		ErrorController(w, r, http.StatusMethodNotAllowed, "")
		return
	}
	if r.URL.Path != "/posts" {
		ErrorController(w, r, http.StatusNotFound, "")
	}
	ParseFileController(w, r, "users/create_poste", nil)
}
