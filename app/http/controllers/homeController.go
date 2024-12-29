package controllers

import (
	"net/http"
)

func HomeController(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		ErrorController(w,r,http.StatusNotFound)
		return
	}
	
	ParseFileController(w,r,"index",nil)
}
