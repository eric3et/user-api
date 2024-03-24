package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/eric3et/go_tutorial_1/api"
	"github.com/eric3et/go_tutorial_1/internal/tools"
)

func ListUser(w http.ResponseWriter, r *http.Request) {
	users, err := tools.DBListUser()
	if err != nil {
		api.InternalErrorHandler(w, err)
		return
	}

	//Write Response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}
