package main

import (
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"github.com/ISNewton/rss-aggregator/internal/database/schema"
	"github.com/ISNewton/rss-aggregator/models"
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

	user, userError := apiCfg.DB.CreateUser(r.Context(), schema.CreateUserParams{
		ID:        uuid.New(),
		Name:      params.Name,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})

	if userError != nil {
		respondWithError(w, http.StatusInternalServerError, userError.Error())
	}

	respondWithJson(w, 200, models.ConvertToUserModel(user))

}

func (apiCfg apiConfig) handleIssueApiKey(w http.ResponseWriter, r *http.Request) {

	type parameters struct {
		UserId uuid.UUID `json:"user_id"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}

	err := decoder.Decode(&params)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
	}

	generatedApiKey := generateApiKey()

	log.Println(params.UserId)

	apiKey, err := apiCfg.DB.IssueApiKey(r.Context(), schema.IssueApiKeyParams{
		ID:     params.UserId,
		ApiKey: sql.NullString{String: generatedApiKey, Valid: true},
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}

	respondWithJson(w, 200, apiKey.String)
}

func (apicfg apiConfig) handleRetrieveUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		ApiKey string `json:"api_key"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}

	err := decoder.Decode(&params)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")

	}

	user, err := apicfg.DB.GetUserByApiKey(r.Context(), sql.NullString{String: params.ApiKey, Valid: true})

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}

	respondWithJson(w, 200, models.ConvertToUserModel(user))

}

func generateApiKey() string {
	b := make([]byte, 30)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return base64.StdEncoding.EncodeToString(b)
}
