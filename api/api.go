package api

import (
	"encoding/json"
	"net/http"
)

// Error Response
type Error struct {
	Message string
}

func writeError(w http.ResponseWriter, message string, code int) {
	resp := Error{
		Message: message,
	}

	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

var (
	NotFoundErrorHandler = func(w http.ResponseWriter, err error) {
		writeError(w, err.Error(), http.StatusNotFound)
	}
	RequestErrorHandler = func(w http.ResponseWriter, err error) {
		writeError(w, err.Error(), http.StatusBadRequest)
	}
	InternalErrorHandler = func(w http.ResponseWriter, err error) {
		writeError(w, "An Unexpected Error Occured.", http.StatusInternalServerError)
	}
)
