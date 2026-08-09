package handler

import (
	"encoding/json"
	"net/http"
	"movies-api/internal/service"
)

type GenreHandler struct {
	service *service.GenreService
}

func NewGenreHandler(service *service.GenreService) *GenreHandler {
	return &GenreHandler{
		service: service,
	}
}

func (h *GenreHandler) GetAll(w http.ResponseWriter, r *http.Request ) {

	genres, err := h.service.GetAll()
	if err != nil {
		http.Error(w, "Failed to fetch genres", http.StatusInternalServerError )
		return
	}

	w.Header().Set("Content-Type", "application/json")
	
	w.WriteHeader(http.StatusOK )

	err = json.NewEncoder(w).Encode(genres)

	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError )
	}
}