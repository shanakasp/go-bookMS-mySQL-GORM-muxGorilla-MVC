package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/princesp/go-bookms/models"
)



func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	var books []models.Book

	result := models.DB.Find(&books)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}