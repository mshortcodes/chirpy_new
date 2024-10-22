package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

// handlerValidateChirp checks for chirp length and censors bad words.
func handlerValidateChirp(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Body string `json:"body"`
	}

	type response struct {
		CleanedBody string `json:"cleaned_body"`
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

	cleaned := cleanChirp(params.Body)

	respondWithJson(w, http.StatusOK, response{
		CleanedBody: cleaned,
	})
}

// cleanChirp censors bad words.
func cleanChirp(body string) string {
	badWords := map[string]bool{
		"kerfuffle": true,
		"sharbert":  true,
		"fornax":    true,
	}

	words := strings.Fields(body)
	const stars = "****"
	for i, word := range words {
		lowered := strings.ToLower(word)
		if badWords[lowered] {
			words[i] = stars
		}
	}

	cleaned := strings.Join(words, " ")
	return cleaned
}
