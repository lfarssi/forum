package routes

import (
	"fmt"
	"forum/app/http/controllers"
	"forum/app/http/controllers/auth"
	"forum/app/http/controllers/static"
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
	http.HandleFunc("/resources/", static.CssJsController)
	fmt.Println("Server is running on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err!= nil {
        fmt.Println("err starting the server : ", err)
		return
    }

}
