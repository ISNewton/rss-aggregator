package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithJson(w http.ResponseWriter, code int, data interface{}) {
	jsonPayload, err := json.Marshal(data)

	if err != nil {
		log.Printf("Error marshalling json: %v", err)
		w.WriteHeader(500)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(jsonPayload)

}

func respondWithError(w http.ResponseWriter, code int, message string) {
	if code >= 500 {
		type errorResponse struct {
			Error string `json:"error"`
			Code  int    `json:"code"`
		}

		respondWithJson(w, code, errorResponse{Error: message, Code: code})
	}

}
