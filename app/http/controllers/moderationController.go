package controllers

import (
	"encoding/json"
	"fmt"
	"forum/app/models"
	"html"
	"net/http"
	"strings"
)

func RequestModeration(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		ErrorController(w, r, http.StatusMethodNotAllowed, "")
		return
	}

	reason := html.EscapeString(r.PostFormValue("reason"))
	if strings.TrimSpace(reason) == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Error:reason  field  empty ")
		return
	} else if len(reason) > 255 || len(reason) < 4 {

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusRequestEntityTooLarge)
		json.NewEncoder(w).Encode("Error: reason or Content field's too large")
		return
	}

	userId, err := models.GetUserId(r)
	if err != nil {
		LogoutController(w, r)
		return
	}
	role, err := models.GetRoleUser(userId)
	if err != nil {
		http.Error(w, "Error: could not verify user role", http.StatusInternalServerError)
		return
	}
	if role != "user" {
		http.Error(w, "Error: only users with role 'user' can request moderation", http.StatusForbidden)
		return
	}
	err1 := models.AddModRequest(reason, userId)
	if err1 != nil {
		fmt.Println(err1)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Error: Cannot create post")
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Moderation request submitted successfully.")

}
