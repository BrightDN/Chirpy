package main

import (
	"net/http"

	"github.com/BrightDN/Chirpy/internal/auth"
)

func (cfg *apiConfig) endpointRevokeToken(w http.ResponseWriter, r *http.Request) {

	tok, err := auth.GetBearerToken(r.Header)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := cfg.Db.RevokeToken(r.Context(), tok); err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusNoContent, "")
}
