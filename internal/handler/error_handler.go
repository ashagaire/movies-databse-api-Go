package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"movies-api/internal/models/errors"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func HandleError(w http.ResponseWriter, message string, err error) {

	var statusCode int

	if message == "" && err != nil {
		message = err.Error()
	} else if message == "" {
		message = "Server Error"
	}

	switch {
	case errors.Is(err, models.ErrNotFound):
		statusCode = http.StatusNotFound
	case errors.Is(err, models.ErrInvalidInput):
		statusCode = http.StatusBadRequest
	case errors.Is(err, models.ErrConflict):
		statusCode = http.StatusConflict // 409 Conflict
	// For any other error (like database connection issues), return a 500
	default:

		statusCode = http.StatusInternalServerError
		message = err.Error()

	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(ErrorResponse{Error: message})
}
