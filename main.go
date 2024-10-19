package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	const port = "8080"
	const filepathRoot = "."

	mux := http.NewServeMux()
	mux.Handle("/app/", http.StripPrefix("/app", http.FileServer(http.Dir(filepathRoot))))
	mux.HandleFunc("/healthz", handlerReadiness)

	srv := &http.Server{
		Handler: mux,
		Addr:    ":" + port,
	}

	fmt.Printf("Serving files from %s on port: %s\n", filepathRoot, port)
	log.Fatal(srv.ListenAndServe())
}
