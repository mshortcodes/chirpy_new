package main

import (
	"net/http"

	"github.com/mshortcodes/chirpy_new/internal/auth"
)

// handlerRevoke revokes a refresh token.
func (cfg *apiConfig) handlerRevoke(w http.ResponseWriter, r *http.Request) {
	refreshToken, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "token is required", err)
		return
	}

	_, err = cfg.dbQueries.RevokeRefreshToken(r.Context(), refreshToken)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't revoke token", err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
