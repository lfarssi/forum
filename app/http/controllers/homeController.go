package controllers

import (
	errorcont "forum/app/http/controllers/error"
	utils "forum/app/http/controllers/utils"

	"net/http"
)

func HomeController(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		errorcont.ErrorController(w, r, http.StatusMethodNotAllowed)
		return
	}
	if r.URL.Path != "/" {
		errorcont.ErrorController(w, r, http.StatusNotFound)
		return
	}
	utils.TemplateController(w, r, "/index", nil)
}
