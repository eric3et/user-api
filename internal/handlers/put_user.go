package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/eric3et/go_tutorial_1/api"
	"github.com/eric3et/go_tutorial_1/internal/tools"
	log "github.com/sirupsen/logrus"
)

func PutUser(w http.ResponseWriter, r *http.Request) {
	// Read the request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Error(err)
		api.RequestErrorHandler(w, err)
		return
	}
	defer r.Body.Close()

	// Parse the JSON body into a User struct
	var user tools.User
	err = json.Unmarshal(body, &user)
	if err != nil {
		api.InternalErrorHandler(w, err)
		return
	}

	// Add user to DB
	tools.DBPutUser(user)

	//Write Response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}
