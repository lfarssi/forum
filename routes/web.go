package routes

import (
	"database/sql"
	"fmt"
	"forum/app/http/controllers"
	"forum/app/http/controllers/auth"
	"net/http"
)

func Router(db *sql.DB) {
	controllers.SessionController(db)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		controllers.HomeController(w, r, db) 
	})
	http.HandleFunc("/add-post", func(w http.ResponseWriter, r *http.Request) {
		controllers.PostController(w, r, db) 
	})
	http.HandleFunc("/submit-add-post", func(w http.ResponseWriter, r *http.Request) {
		controllers.PostController(w, r, db) 
	})
	// http.HandleFunc("/submit-add-post", controllers.PostController)
	http.HandleFunc("/categories", controllers.CategoryController)
	http.HandleFunc("/comments", controllers.CommentController)
	http.HandleFunc("/reacts", controllers.ReactController)
	http.HandleFunc("/login", auth.LoginController)
	http.HandleFunc("/register", auth.RegisterController)
	http.HandleFunc("/logout", auth.LogoutController)
	

	http.HandleFunc("/filter-posts", func(w http.ResponseWriter, r *http.Request) {
        category := r.URL.Query().Get("category") 
        controllers.FilterPostsController(w, r, db, category)
    })


//serve js and css 
	http.HandleFunc("/css/", controllers.CssController)
	http.HandleFunc("/js/", controllers.JsController)


	fmt.Println("Server is running on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("err starting the server : ", err)
		return
	}

}
