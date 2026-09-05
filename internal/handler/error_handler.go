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

func HandleError(w http.ResponseWriter, err error ) {
	var statusCode int
	var message string

	switch {

	case errors.Is(err, models.ErrNotFound):
		statusCode = http.StatusNotFound
		message = err.Error( )
	case errors.Is(err, models.ErrInvalidInput):
		statusCode = http.StatusBadRequest
		message = err.Error( )
	case errors.Is(err, models.ErrConflict):
		statusCode = http.StatusConflict // 409 Conflict
		message = err.Error( )
	
	// For any other error (like database connection issues), return a 500
	default:
		
		statusCode = http.StatusInternalServerError
		message = err.Error( ) 

	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(ErrorResponse{Error: message})
}
