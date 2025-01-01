package controllers


import (
	"net/http"
)

func JsController(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		ErrorController(w, r, http.StatusMethodNotAllowed)
		return
	}
	// filePath := "resources/js" + r.URL.Path[len("/js/"):]
	// if _, err := os.Stat(filePath); os.IsNotExist(err) {
	// 	fmt.Println("File does not exist:", filePath)
	// 	ErrorController(w, r, http.StatusNotFound)
	// 	return
	// }
	fs := http.Dir("resources/js")
	http.StripPrefix("/js/", http.FileServer(fs)).ServeHTTP(w, r)
}
