package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/BrightDN/Chirpy/internal/database"
	"github.com/BrightDN/Chirpy/internal/endpoints"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	godotenv.Load()
	dbURL := os.Getenv("DB_URL")
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Something went wrong loading the postgres db: %v", err)
	}

	dbQueries := database.New(db)

	const filepathRoot = "."
	const port = "8080"

	apiCfg := endpoints.ApiConfig{
		Db:       dbQueries,
		Platform: os.Getenv("PLATFORM"),
		Secret:   os.Getenv("JWTSECRET"),
		PolkaKey: os.Getenv("POLKA_KEY"),
	}

	mu := http.NewServeMux()
	fileServer := http.FileServer(http.Dir(filepathRoot))
	handler := http.StripPrefix("/app", fileServer)

	mu.Handle("/app/", apiCfg.MiddlewareMetricsInc(handler))
	mu.HandleFunc("GET /admin/healthz", endpoints.EndpointReadiness)

	mu.HandleFunc("GET /admin/metrics", apiCfg.EndpointMetrics)
	mu.HandleFunc("POST /admin/reset", apiCfg.EndpointReset)

	mu.HandleFunc("GET /api/chirps", apiCfg.EndpointGetChirps)
	mu.HandleFunc("GET /api/chirps/{chirpID}", apiCfg.EndpointGetChirp)
	mu.HandleFunc("DELETE /api/chirps/{chirpID}", apiCfg.EndpointDeleteChirp)

	mu.HandleFunc("POST /api/chirps", apiCfg.EndpointCreateChirp)
	mu.HandleFunc("POST /api/users", apiCfg.EndpointCreateUser)
	mu.HandleFunc("PUT /api/users", apiCfg.EndpointUpdateUserData)
	mu.HandleFunc("POST /api/login", apiCfg.EndpointLogin)
	mu.HandleFunc("POST /api/refresh", apiCfg.EndpointRefreshToken)
	mu.HandleFunc("POST /api/revoke", apiCfg.EndpointRevokeToken)
	mu.HandleFunc("POST /api/polka/webhooks", apiCfg.EndpointUpgradeWebhook)

	server := http.Server{
		Addr:    ":" + port,
		Handler: mu,
	}

	log.Printf("Serving files from %s on port: %s\n", filepathRoot, port)
	log.Fatal(server.ListenAndServe())
}
