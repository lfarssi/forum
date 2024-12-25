package controllers

import (
	"fmt"
	"forum/app/http/controllers/error"
	"html/template"
	"net/http"
)

func TemplateController(w http.ResponseWriter, r *http.Request, temp string, data any) {
	res, err := template.ParseFiles("resources/views/" + temp + ".html")
	if err != nil {
		fmt.Println("error parsing")
		controllers.ErrorController(w, r, http.StatusInternalServerError)
		return
	}
	if err = res.Execute(w, data); err != nil {
		fmt.Println("error executing template")
		controllers.ErrorController(w, r, http.StatusInternalServerError)
		return
	}
}
