package main

import (
	"fmt"
	"net/http"
)

// handlerMetrics writes fileServerHits to the response.
func (cfg *apiConfig) handlerMetrics(w http.ResponseWriter, r *http.Request) {
	msg := fmt.Sprintf(`
<html>
	<body>
	<h1>Welcome, Chirpy Admin</h1>
	<p>Chirpy has been visited %d times!</p>
	</body>
</html>`, cfg.fileServerHits.Load())
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(msg))
}

// middlewareMetricsInc increments fileServerHits when /app is visited.
func (cfg *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg.fileServerHits.Add(1)
		next.ServeHTTP(w, r)
	})
}
