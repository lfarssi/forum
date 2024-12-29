package controllers

import (
	"database/sql"
	"fmt"
	"forum/app/models"
	"net/http"
)

func fetchCategories(w http.ResponseWriter,r *http.Request, db *sql.DB) ([]models.Category, error) {
	categories := []models.Category{}
	rows, err := db.Query("SELECT id, name FROM categories")
	if err != nil {
		fmt.Println("error querying categories: ", err)
		ErrorController(w, r, http.StatusInternalServerError)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var category models.Category
		if err := rows.Scan(&category.ID, &category.Name); err != nil {
			fmt.Println("error scanning categories: ", err)
			ErrorController(w, r, http.StatusInternalServerError)
			return nil, err
		}
		categories = append(categories, category)
	}

	return categories, nil
}
