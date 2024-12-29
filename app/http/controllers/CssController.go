package controllers

import (
	"net/http"
)

func CssController(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		ErrorController(w, r, http.StatusMethodNotAllowed)
		return
	}
	// filePath := "resources/css" + r.URL.Path[len("/css/"):]
	// if _, err := os.Stat(filePath); os.IsNotExist(err) {
	// 	fmt.Println("File does not exist:", filePath)
	// 	ErrorController(w, r, http.StatusNotFound)
	// 	return
	// }
	fs := http.Dir("resources/css")
	http.StripPrefix("/css/", http.FileServer(fs)).ServeHTTP(w, r)
}
