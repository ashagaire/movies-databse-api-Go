package handler

import (
	"encoding/json"
	"movies-api/internal/models"
	"movies-api/internal/service"
	"net/http"
)

type GenreHandler struct {
	service *service.GenreService
}

func NewGenreHandler(service *service.GenreService) *GenreHandler {
	return &GenreHandler{
		service: service,
	}
}

func (h *GenreHandler) Create(w http.ResponseWriter, r *http.Request) {

	var genre models.Genre

	err := json.NewDecoder(r.Body).Decode(&genre)
	if err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	err = h.service.Create(&genre)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	err = json.NewEncoder(w).Encode(genre)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func (h *GenreHandler) GetAll(w http.ResponseWriter, r *http.Request) {

	genres, err := h.service.GetAll()
	if err != nil {
		http.Error(w, "Failed to fetch genres", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(genres)

	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
