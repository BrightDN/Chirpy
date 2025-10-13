package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/BrightDN/Chirpy/internal/database"
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

	apiCfg := apiConfig{
		Db:       dbQueries,
		Platform: os.Getenv("PLATFORM"),
		Secret:   os.Getenv("JWTSECRET"),
	}

	mu := http.NewServeMux()
	fileServer := http.FileServer(http.Dir(filepathRoot))
	handler := http.StripPrefix("/app", fileServer)

	mu.Handle("/app/", apiCfg.middlewareMetricsInc(handler))
	mu.HandleFunc("GET /admin/healthz", endpointReadiness)

	mu.HandleFunc("GET /admin/metrics", apiCfg.endpointMetrics)
	mu.HandleFunc("POST /admin/reset", apiCfg.endpointReset)

	mu.HandleFunc("GET /api/chirps", apiCfg.endpointGetChirps)
	mu.HandleFunc("GET /api/chirps/{chirpID}", apiCfg.endpointGetChirp)

	mu.HandleFunc("POST /api/chirps", apiCfg.endpointCreateChirp)
	mu.HandleFunc("POST /api/users", apiCfg.endpointCreateUser)
	mu.HandleFunc("POST /api/login", apiCfg.endpointLogin)
	mu.HandleFunc("POST /api/refresh", apiCfg.endpointRefreshToken)
	mu.HandleFunc("POST /api/revoke", apiCfg.endpointRevokeToken)

	server := http.Server{
		Addr:    ":" + port,
		Handler: mu,
	}

	log.Printf("Serving files from %s on port: %s\n", filepathRoot, port)
	log.Fatal(server.ListenAndServe())
}
