package main

import (
	"net/http"
	"time"

	"github.com/mshortcodes/chirpy_new/internal/auth"
)

// handlerRefresh generates new access token (JWT) if provided a refresh token.
func (cfg *apiConfig) handlerRefresh(w http.ResponseWriter, r *http.Request) {
	type response struct {
		Token string `json:"token"`
	}

	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "auth token required", err)
		return
	}

	user, err := cfg.dbQueries.GetUserFromRefreshToken(r.Context(), token)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "couldn't get user for refresh token", err)
		return
	}

	accessToken, err := auth.MakeJWT(user.ID, cfg.jwtSecret, time.Hour)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't make access token", err)
		return
	}

	respondWithJson(w, http.StatusOK, response{
		Token: accessToken,
	})
}
