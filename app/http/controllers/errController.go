package controllers

import (
	"forum/app/models"
	"net/http"
	"text/template"
)

func ErrorController(w http.ResponseWriter, r *http.Request , StatusCode int) {
	errPage := "resources/views/error.html"
	tmp, err := template.ParseFiles(errPage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}
	w.WriteHeader(StatusCode)
	errType := models.ErrorType{
	   Code : StatusCode,
	   Message : http.StatusText(StatusCode),
	}
	if err := tmp.Execute(w, errType); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
