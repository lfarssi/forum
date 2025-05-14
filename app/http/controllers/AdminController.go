package controllers

import (
	"forum/app/models"
	"net/http"
	"strconv"
)

func HandleModRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		ErrorController(w, r, http.StatusMethodNotAllowed, "Method Not Allowed sss")
		return
	}

	err := r.ParseMultipartForm(10 << 20) 
	if err != nil {
		ErrorController(w, r, http.StatusBadRequest, "Invalid Form Data")
		return
	}

	userIDStr := r.FormValue("user_id")
	action := r.FormValue("action")
	userID, err := strconv.Atoi(userIDStr)

	if err != nil {
		ErrorController(w, r, http.StatusBadRequest, "Invalid User ID")
		return
	}

	if action == "accept" {
		err = models.UpdateUserRole(userID, "moderator")
		if err != nil {
			ErrorController(w, r, http.StatusInternalServerError, "Failed to Promote User")
			return
		}
	}

	// In both accept and refuse: delete the request
	err = models.DeleteModRequest(userID)
	if err != nil {
		ErrorController(w, r, http.StatusInternalServerError, "Failed to Delete Request")
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
