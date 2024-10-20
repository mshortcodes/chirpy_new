package main

import (
	"fmt"
	"net/http"
)

// increments fileServerHits when /app is visited
func (cfg *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg.fileServerHits.Add(1)
		next.ServeHTTP(w, r)
	})
}

// writes fileServerHits to the response
func (cfg *apiConfig) handlerMetrics(w http.ResponseWriter, r *http.Request) {
	msg := fmt.Sprintf("Hits: %d", cfg.fileServerHits.Load())
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(msg))
}

// resets fileServerHits
func (cfg *apiConfig) handlerReset(w http.ResponseWriter, r *http.Request) {
	cfg.fileServerHits.Store(0)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hits reset to 0"))
}
