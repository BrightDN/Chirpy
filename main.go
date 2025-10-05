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
		log.Fatalf("Something went wrong: %v", err)
	}

	dbQueries := database.New(db)

	const filepathRoot = "."
	const port = "8080"

	apiCfg := apiConfig{
		Db:       dbQueries,
		Platform: os.Getenv("PLATFORM"),
	}

	mu := http.NewServeMux()
	fileServer := http.FileServer(http.Dir(filepathRoot))
	handler := http.StripPrefix("/app", fileServer)

	mu.Handle("/app/", apiCfg.middlewareMetricsInc(handler))
	mu.HandleFunc("GET /admin/healthz", endpointReadiness)

	mu.HandleFunc("GET /admin/metrics", apiCfg.endpointMetrics)
	mu.HandleFunc("POST /admin/reset", apiCfg.endpointReset)
	mu.HandleFunc("POST /api/validate_chirp", apiCfg.endpointValidateChirp)
	mu.HandleFunc("POST /api/users", apiCfg.endpointCreateUser)

	server := http.Server{
		Addr:    ":" + port,
		Handler: mu,
	}

	log.Printf("Serving files from %s on port: %s\n", filepathRoot, port)
	log.Fatal(server.ListenAndServe())
}
