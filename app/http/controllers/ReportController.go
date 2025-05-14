package controllers

import (
	"forum/app/models"
	"net/http"
	"strconv"
)

func DeleteReportHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		ErrorController(w, r, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	reportID := r.FormValue("report_id")
	if reportID == "" {
		ErrorController(w, r, http.StatusBadRequest, "Report ID is required")
		return
	}

	// Convert reportID to int
	reportIDInt, err := strconv.Atoi(reportID)
	if err != nil {
		ErrorController(w, r, http.StatusBadRequest, "Invalid Report ID")
		return
	}

	err = models.DeleteReport(reportIDInt)
	if err != nil {
		ErrorController(w, r, http.StatusInternalServerError, "Failed to delete report")
		return
	}

	// Redirect to moderator dashboard after deletion
	http.Redirect(w, r, "/moderator/dashboard", http.StatusSeeOther)
}
