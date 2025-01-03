package auth

import (
	"fmt"
	"forum/app/http/controllers"
	"net/http"
)

func RegisterController(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/register" {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	} else if r.Method != "GET" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	username := r.URL.Query().Get("username")
	email := r.URL.Query().Get("email")
	password := r.URL.Query().Get("password")
	fmt.Println("username: ", username)
	fmt.Println("email: ", email)
	fmt.Println("password: ", password)
	controllers.ParseFileController(w, r, "auth/register", "")

}
