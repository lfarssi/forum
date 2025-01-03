package static

import (
	"forum/app/http/controllers"
	"net/http"
	"os"
	"strings"
)

func CssJsController(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		controllers.ErrorController(w, r, http.StatusMethodNotAllowed)
		return
	}
	filePath := strings.TrimPrefix(r.URL.Path, "/resources/")
	fullPath := "resources/" + filePath

	info, err := os.Stat(fullPath) 
	if err != nil {
		if os.IsNotExist(err) {
			controllers.ErrorController(w, r, http.StatusNotFound)
		} else {
			controllers.ErrorController(w, r, http.StatusInternalServerError)
		}
		return
	}
	if info.IsDir() {
		controllers.ErrorController(w, r, http.StatusForbidden)
		return
	}	

	fs := http.Dir("resources")
	http.StripPrefix("/resources/", http.FileServer(fs)).ServeHTTP(w, r)
	
}


