package main

import (
	"encoding/json"
	"net/http"
)

// validates a chirp's length
func handlerValidateChirp(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Body string `json:"body"`
	}

	type response struct {
		Valid bool `json:"valid"`
	}

	var params parameters

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't decode parameters", err)
		return
	}

	const maxChirpLen = 140
	if len(params.Body) > maxChirpLen {
		respondWithError(w, http.StatusBadRequest, "chirp is too long", nil)
		return
	}

	respondWithJson(w, http.StatusOK, response{
		Valid: true,
	})
}
