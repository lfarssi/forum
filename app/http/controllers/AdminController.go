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

// func GetPendingReportsHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != "GET" {
// 		ErrorController(w, r, http.StatusMethodNotAllowed, "Method not allowed")
// 		return
// 	}

// 	// Check if user is admin (you should implement proper session checking)
// 	// session, _ := store.Get(r, "session")
// 	// userRole := session.Values["role"]
// 	// if userRole != "admin" {
// 	//     ErrorController(w, r, http.StatusForbidden, "Access denied")
// 	//     return
// 	// }

// 	// Get pending reports from the database
// 	pendingReports, err := models.GetPendingReports()
// 	if err != nil {
// 		ErrorController(w, r, http.StatusInternalServerError, "Failed to fetch pending reports")
// 		return
// 	}

// 	// Return the reports as JSON
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(pendingReports)
// }

// // AdminReportDecisionHandler - For admins to approve/reject report actions
// func AdminReportDecisionHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != "POST" {
// 		ErrorController(w, r, http.StatusMethodNotAllowed, "Method not allowed")
// 		return
// 	}

// 	// Check if user is admin
// 	// session, _ := store.Get(r, "session")
// 	// userRole := session.Values["role"]
// 	// if userRole != "admin" {
// 	//     ErrorController(w, r, http.StatusForbidden, "Access denied")
// 	//     return
// 	// }

// 	reportID := r.FormValue("report_id")
// 	decision := r.FormValue("decision")

// 	if reportID == "" || decision == "" {
// 		ErrorController(w, r, http.StatusBadRequest, "Report ID and decision are required")
// 		return
// 	}

// 	// Validate decision
// 	if decision != "approved" && decision != "rejected" {
// 		ErrorController(w, r, http.StatusBadRequest, "Invalid decision. Must be 'approved' or 'rejected'")
// 		return
// 	}

// 	// Convert reportID to int
// 	reportIDInt, err := strconv.Atoi(reportID)
// 	if err != nil {
// 		ErrorController(w, r, http.StatusBadRequest, "Invalid Report ID")
// 		return
// 	}

// 	// Update report status in database
// 	err = models.UpdateReportStatus(reportIDInt, decision)
// 	if err != nil {
// 		ErrorController(w, r, http.StatusInternalServerError, "Failed to update report status")
// 		return
// 	}

// 	// Return success response
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(map[string]string{"status": "success", "message": "Decision saved successfully"})
// }
