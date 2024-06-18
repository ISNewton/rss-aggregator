package main

import "net/http"

func handleError(w http.ResponseWriter, r *http.Request) {

	respondWithError(w, 500, "The server encountered an error")

}
