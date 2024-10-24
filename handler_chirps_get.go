package main

import (
	"net/http"

	"github.com/google/uuid"
)

// handlerChirpsGet retrieves all chirps in the database.
func (cfg *apiConfig) handlerChirpsGet(w http.ResponseWriter, r *http.Request) {
	dbChirps, err := cfg.dbQueries.GetChirps(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't get chirps", err)
		return
	}

	var chirps []Chirp

	for _, dbChirp := range dbChirps {
		chirps = append(chirps, Chirp{
			ID:        dbChirp.ID,
			CreatedAt: dbChirp.CreatedAt,
			UpdatedAt: dbChirp.UpdatedAt,
			Body:      dbChirp.Body,
			UserID:    dbChirp.UserID,
		})
	}

	respondWithJson(w, http.StatusOK, chirps)
}

// handlerChirpsGetByID retrieves a single chirp by ID.
func (cfg *apiConfig) handlerChirpsGetByID(w http.ResponseWriter, r *http.Request) {
	chirpID := r.PathValue("chirpID")

	uuid, err := uuid.Parse(chirpID)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid chirp ID", err)
		return
	}

	chirp, err := cfg.dbQueries.GetChirp(r.Context(), uuid)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "chirp not found", err)
		return
	}

	respondWithJson(w, http.StatusOK, Chirp{
		ID:        chirp.ID,
		CreatedAt: chirp.CreatedAt,
		UpdatedAt: chirp.UpdatedAt,
		Body:      chirp.Body,
		UserID:    chirp.UserID,
	})
}
