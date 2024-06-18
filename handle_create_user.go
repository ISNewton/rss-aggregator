package main

import (
	"encoding/json"
	"net/http"
)

func (apiCfg apiConfig) handleCreateUser(w http.ResponseWriter, r *http.Request) {

	type parameters struct {
		name string `name`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}

	err := decoder.Decode(&params)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
	}

	apiCfg.DB

}
