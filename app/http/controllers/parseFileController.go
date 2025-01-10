package controllers

import (
	"fmt"
	"net/http"
	"text/template"
)

func ParseFileController(w http.ResponseWriter, r *http.Request, filename string, data any) {
	filepath := "./resources/views/" + filename + ".html"
	components := []string{
		"./resources/views/components/navbar.html",
		"./resources/views/components/footer.html",
		"./resources/views/components/menu.html",
	}
	allFiles := append([]string{filepath}, components...)
	// fmt.Println("filepath: ", filepath)
	temp, err := template.ParseFiles(allFiles...)
	if err != nil {
		fmt.Println("error parsing; ", err)
		ErrorController(w, r, http.StatusInternalServerError, "")
		return
	}
	err1 := temp.Execute(w, data)
	if err1 != nil {
		fmt.Println("error executing; ", err1)
		ErrorController(w, r, http.StatusInternalServerError, "")
		return
	}
}
