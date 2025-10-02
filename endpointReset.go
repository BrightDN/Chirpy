package main

import (
	"net/http"
)

func (cfg *apiConfig) endpointReset(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(http.StatusText(http.StatusOK)))
	cfg.fileserverHits.Store(0)
}
