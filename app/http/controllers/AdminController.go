package controllers

import (
	"forum/app/models"
	"net/http"
	"strconv"
)



func HandleModRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		ErrorController(w, r, http.StatusMethodNotAllowed, "Method Not Allowed")
		return
	}

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		ErrorController(w, r, http.StatusBadRequest, "Invalid Form Data")
		return
	}

	userIDStr := r.FormValue("user_id")
	role := r.FormValue("role")
	userID, err := strconv.Atoi(userIDStr)

	if err != nil || (role != "user" && role != "moderator") {
		ErrorController(w, r, http.StatusBadRequest, "Invalid Data")
		return
	}

	err = models.UpdateUserRole(userID, role)
	if err != nil {
		ErrorController(w, r, http.StatusInternalServerError, "Failed to Change User Role")
		return
	}
	err = models.UpdateModRequestStatus(userID, "accepted")
	if err != nil {
		http.Error(w, "Could not update request status", http.StatusInternalServerError)
		return
	}

	// Only delete the mod request if the role is "user" (i.e., request refused)
	if role == "user" {
	err = models.DeleteModRequest(userID)
	if err != nil {
		ErrorController(w, r, http.StatusInternalServerError, "Failed to Delete Mod Request")
		return
	}
}

	w.WriteHeader(http.StatusOK)
}

