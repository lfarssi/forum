package controllers

import (
	"html"
	"net/http"
	"strconv"
	"time"

	"forum/app/models"
)


func CreatCommentController(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		ErrorController(w, r, http.StatusMethodNotAllowed, "")
		return
	}
	r.ParseForm()
	var err error
	var comment models.Comment
	comment.Content = html.EscapeString(r.FormValue("content"))  
	comment.CreatedAt= time.Now()
	comment.PostID, err =strconv.Atoi( r.FormValue("post_id"))
	if err != nil{
		ErrorController(w,r , http.StatusBadRequest, "Post number Invalid")
		return 
	}
	if len(comment.Content)==0{
		ErrorController(w,r , http.StatusBadRequest, "Empty String")
		return
	}else if len(comment.Content)>=10000{
		ErrorController(w,r , http.StatusRequestEntityTooLarge, "")
		return
	}
	userId,err:=models.GetUserId(r)
	if err != nil{
		ErrorController(w,r , http.StatusInternalServerError, "")
		return 
	}
	comment.UserID = userId 
	
	if err = models.CreatComment(comment) ; err != nil {
		ErrorController(w, r, http.StatusInternalServerError, "")
		return
	}
	http.Redirect(w,r , "/", http.StatusFound)
}
