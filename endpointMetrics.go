package main

import (
	"fmt"
	"net/http"
)

func (cfg *apiConfig) endpointMetrics(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	count := cfg.fileserverHits.Load()
	welcomeText := "Welcome, Chirpy Admin"
	fullText := fmt.Sprintf(`
	<html>
		<body>
			<h1>%s</h1>
			<p>Chirpy has been visited %d times!</p>
		</body>
	</html>`, welcomeText, count)
	w.Write([]byte(fullText))
}
