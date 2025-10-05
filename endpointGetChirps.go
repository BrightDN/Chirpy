package main

import (
	"net/http"

	"github.com/google/uuid"
)

func (cfg *apiConfig) endpointGetChirps(w http.ResponseWriter, r *http.Request) {

	chirps, err := cfg.Db.GetAllChirps(r.Context())
	if err != nil {
		writeError(w, http.StatusInternalServerError, "An error occured handling your request")
		return
	}

	out := make([]chirpsResp, len(chirps))
	for i, r := range chirps {
		out[i] = chirpsResp{
			Id:        r.ID,
			CreatedAt: r.CreatedAt,
			UpdatedAt: r.UpdatedAt,
			Body:      r.Body,
			UserID:    r.UserID,
		}
	}

	writeJSON(w, http.StatusOK, out)
}

func (cfg *apiConfig) endpointGetChirp(w http.ResponseWriter, r *http.Request) {
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

	writeJSON(w, http.StatusOK, chirpsResp{
		Id:        chirp.ID,
		CreatedAt: chirp.CreatedAt,
		UpdatedAt: chirp.UpdatedAt,
		Body:      chirp.Body,
		UserID:    chirp.UserID,
	})
}
