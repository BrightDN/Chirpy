package main

import (
	"net/http"

	"github.com/BrightDN/Chirpy/internal/auth"
	"github.com/google/uuid"
)

func (cfg *apiConfig) endpointDeleteChirp(w http.ResponseWriter, r *http.Request) {
	authTok, err := auth.GetBearerToken(r.Header)
	if err != nil {
		writeError(w, http.StatusUnauthorized, err.Error())
		return
	}

	userID, err := auth.ValidateJWT(authTok, cfg.Secret)
	if err != nil {
		writeError(w, http.StatusUnauthorized, err.Error())
		return
	}

	chirpID := r.PathValue("chirpID")
	if chirpID == "" {
		writeError(w, http.StatusBadRequest, "Include a chirpID")
		return
	}

	chirpUUID, err := uuid.Parse(chirpID)
	if err != nil {
		writeError(w, http.StatusBadRequest, "Invalid ID format")
		return
	}

	chirp, err := cfg.Db.GetChirp(r.Context(), chirpUUID)
	if err != nil {
		writeError(w, http.StatusNotFound, "No chirp with this ID exists")
		return
	}

	if userID != chirp.UserID {
		writeError(w, http.StatusForbidden, "you are not the owner of this chirp")
		return
	}

	if err := cfg.Db.DeleteChirp(r.Context(), chirp.ID); err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeJSON(w, http.StatusNoContent, "")
}
