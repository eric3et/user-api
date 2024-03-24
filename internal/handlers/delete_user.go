package handlers

import (
	"net/http"
	"strconv"

	"github.com/eric3et/go_tutorial_1/api"
	"github.com/eric3et/go_tutorial_1/internal/tools"
	"github.com/go-chi/chi/v5"
)

func DeleteUser(w http.ResponseWriter, r *http.Request) {

	// Parse parameters from the URL query string
	id_s := chi.URLParam(r, "id")

	id, err := strconv.Atoi(id_s)
	if err != nil {
		api.InternalErrorHandler(w, err)
		return
	}

	tools.DBDeleteUser(id)

	//Write Response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}
