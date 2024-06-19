package main

import (
	"encoding/json"
	"github.com/ISNewton/rss-aggregator/internal/database/schema"
	"github.com/google/uuid"
	"log"
	"net/http"
	"time"
)

func (apiCfg apiConfig) handleCreateUser(w http.ResponseWriter, r *http.Request) {

	type parameters struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}

	err := decoder.Decode(&params)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
	}

	log.Println(params)

	user, userError := apiCfg.DB.CreateUser(r.Context(), schema.CreateUserParams{
		ID:        uuid.New(),
		Name:      params.Name,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})

	if userError != nil {
		respondWithError(w, http.StatusInternalServerError, userError.Error())
	}

	respondWithJson(w, 200, user)

}
