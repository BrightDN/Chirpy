package main

import (
	"log"
	"net/http"
)

func main() {
	const filepathRoot = "."
	const port = "8080"

	apiCfg := apiConfig{}

	mu := http.NewServeMux()
	fileServer := http.FileServer(http.Dir(filepathRoot))
	handler := http.StripPrefix("/app", fileServer)

	mu.Handle("/app/", apiCfg.middlewareMetricsInc(handler))
	mu.HandleFunc("GET /admin/healthz", endpointReadiness)

	mu.HandleFunc("GET /admin/metrics", apiCfg.endpointMetrics)
	mu.HandleFunc("POST /admin/reset", apiCfg.endpointReset)
	mu.HandleFunc("POST /api/validate_chirp", apiCfg.endpointValidateChirp)

	server := http.Server{
		Addr:    ":" + port,
		Handler: mu,
	}

	log.Printf("Serving files from %s on port: %s\n", filepathRoot, port)
	log.Fatal(server.ListenAndServe())
}
