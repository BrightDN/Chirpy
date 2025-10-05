package main

import (
	"net/http"
)

func (cfg *apiConfig) endpointReset(w http.ResponseWriter, r *http.Request) {

	if cfg.Platform != "dev" {
		writeError(w, http.StatusForbidden, "Change to DEV mode for this functionality")
		return
	}

	if err := cfg.Db.DeleteUsers(r.Context()); err != nil {
		writeError(w, http.StatusInternalServerError, "An issue occured while processing your request")
		return
	}

	cfg.fileserverHits.Store(0)

	w.WriteHeader(http.StatusOK)

}
