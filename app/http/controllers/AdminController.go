package controllers

import (
	"forum/app/models"
	"html"
	"net/http"
	"strconv"
)

func HandleModRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		ErrorController(w, r, http.StatusMethodNotAllowed, "Method Not Allowed")
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

func CategoryReportController(w http.ResponseWriter, r *http.Request)  {
	if r.Method != http.MethodPost {
		ErrorController(w, r, http.StatusMethodNotAllowed, "Method Not Allowed")
		return
	}
	categorie:= r.FormValue("category_name")
	categorie= html.EscapeString(categorie)
	if categorie == "" {
		ErrorController(w, r, http.StatusBadRequest, "Category Name is required")
		return
	}
	err := models.AddReportCategory(categorie)
	if err != nil {
		ErrorController(w, r, http.StatusInternalServerError, "Failed to Add Category")
		return
	}
	w.WriteHeader(http.StatusOK)

}
func CategoryDeleteReportController(w http.ResponseWriter, r *http.Request)  {
	if r.Method != http.MethodPost {
		ErrorController(w, r, http.StatusMethodNotAllowed, "Method Not Allowed")
		return
	}
	reportId:= r.FormValue("category_id")
	id, err:= strconv.Atoi(reportId)
	if err != nil || reportId == "" {
		ErrorController(w, r, http.StatusBadRequest, "Invalid Report ID")
		return
	}
	err = models.DeleteCategory(id)
	if err != nil {
		ErrorController(w, r, http.StatusInternalServerError, "Failed to Delete Category")
		return
	}
	w.WriteHeader(http.StatusOK)

}

func HandleRepostPost(w http.ResponseWriter, r *http.Request){
if r.Method != http.MethodPost {
		ErrorController(w, r, http.StatusMethodNotAllowed, "Method Not Allowed")
		return
	}

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		ErrorController(w, r, http.StatusBadRequest, "Invalid Form Data")
		return
	}

	reportId := r.FormValue("report-id")
	desicion := r.FormValue("desicion")
	reportID, err := strconv.Atoi(reportId)

	if err != nil || (desicion != "rejected" && desicion != "accepted") {
		ErrorController(w, r, http.StatusBadRequest, "Invalid Data")
		return
	}
	err = models.UpdateReportPostStatus( reportID,desicion)
	if err != nil {
		http.Error(w, "Could not update report status", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}