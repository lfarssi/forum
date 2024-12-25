package controllers

import (
	"forum/app/models"
	"html/template"
	"net/http"
)

func ErrorController(w http.ResponseWriter, r *http.Request, statusCode int) {
	templatePath := "./resources/views/error.html"
	w.WriteHeader(statusCode)

	tmp, err := template.ParseFiles(templatePath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}
	errorType := models.Error{
		StatusCode:   statusCode,
		ErrorMessage: http.StatusText(statusCode),
	}
	if err := tmp.Execute(w, errorType); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
