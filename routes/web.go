package routes

import (
	"fmt"
	"forum/app/http/controllers"
	"forum/app/http/controllers/auth"
	"net/http"
)

func Router() {
	http.HandleFunc("/", controllers.HomeController)
	http.HandleFunc("/posts", controllers.PostController)
	http.HandleFunc("/categories", controllers.CategoryController)
	http.HandleFunc("/comments", controllers.CommentController)
	http.HandleFunc("/reacts", controllers.ReactController)
	http.HandleFunc("/login", auth.LoginController)
	http.HandleFunc("/register", auth.RegisterController)
	http.HandleFunc("/logout", auth.LogoutController)










	
	fmt.Println("Server running on http://localhost:3000")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println("Error starting server: ", err)
	}
}
