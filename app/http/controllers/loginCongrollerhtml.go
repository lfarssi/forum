package controllers

import "net/http"

func LoginController(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/login" {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	} else if r.Method != "GET" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	ParseFileController(w, r, "auth/login", "")
}
