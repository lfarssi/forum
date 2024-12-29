package controllers

import (
	"fmt"
	"forum/app/models"
	"net/http"
	"text/template"
)

func ErrorController(w http.ResponseWriter, r *http.Request, StatusCode int) {
	errPage := "resources/views/error.html"
	tmp, err := template.ParseFiles(errPage)
	if err != nil {
		fmt.Println("error Parse :", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}
	w.WriteHeader(StatusCode)
	errType := models.ErrorType{
		Code:    StatusCode,
		Message: http.StatusText(StatusCode),
	}
	if err := tmp.Execute(w, errType); err != nil {
		fmt.Println("error Execute :", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
