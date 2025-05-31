package routes

import (
	"forum/app/http/controllers"
	"forum/app/http/middleware"
	"net/http"
)

func ApiRouter() {

	http.HandleFunc("/signIn", controllers.LoginController)
	http.HandleFunc("/signUp",controllers.RegisterController)
	http.HandleFunc("/logout", controllers.LogoutController)
	// http.HandleFunc("/delete_post", controllers.DeleteController)
	http.HandleFunc("/create_post", middleware.AuthMiddleware(middleware.RateLimitMiddleware( controllers.CreatePosts)))
	http.HandleFunc("/react",middleware.AuthMiddleware(middleware.RateLimitMiddleware( controllers.ReactPostController)))
	http.HandleFunc("/create_comment",middleware.AuthMiddleware(middleware.RateLimitMiddleware( controllers.CreateCommentController)))
	http.HandleFunc("/reqmod",middleware.AuthMiddleware(middleware.RateLimitMiddleware( controllers.RequestModeration)))
	http.HandleFunc("/report_post",middleware.AuthMiddleware(middleware.RateLimitMiddleware( controllers.ReportPostController)))
	http.HandleFunc("/delete_report",middleware.AuthMiddleware(middleware.RateLimitMiddleware( controllers.DeleteReportHandler)))
	http.HandleFunc("/delete_post",middleware.AuthMiddleware(middleware.RateLimitMiddleware( controllers.DeletePostHandler)))
	http.HandleFunc("/delete_comment",middleware.AuthMiddleware(middleware.RateLimitMiddleware( controllers.DeleteCommentHandler)))

	http.HandleFunc("/get_reported_posts",middleware.AuthMiddleware(middleware.RateLimitMiddleware( controllers.GetReportedPostsHandler)))
	http.HandleFunc("/moderator/handle_report", middleware.AuthMiddleware(controllers.HandleModeratorReport))
	http.HandleFunc("/handleRequest",middleware.AuthMiddleware( controllers.HandleModRequest))
	http.HandleFunc("/repot-post-responce",middleware.AuthMiddleware( controllers.HandleRepostPost))
	http.HandleFunc("/add-categorie-report",middleware.AuthMiddleware( controllers.CategoryReportController))
	http.HandleFunc("/delete-categorie-report",middleware.AuthMiddleware( controllers.CategoryDeleteReportController))





	http.HandleFunc("/resources/", controllers.CssJsController)
	

}
