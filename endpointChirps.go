package main

import (
	"encoding/json"
	"net/http"
	"slices"
	"strings"

	"github.com/BrightDN/Chirpy/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) endpointCreateChirp(w http.ResponseWriter, r *http.Request) {
	var params params

	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON")
		return
	}

	if len(params.Body) > 140 {
		writeError(w, http.StatusBadRequest, "Chirp is too long")
		return
	}

	cleaned := ReplaceProfanity(params.Body)

	chirp, err := cfg.Db.CreateChirp(r.Context(), database.CreateChirpParams{
		ID:     uuid.New(),
		Body:   cleaned,
		UserID: params.User,
	})

	if err != nil {
		writeError(w, http.StatusInternalServerError, "An error occured processing your chirp")
		return
	}

	writeJSON(w, http.StatusCreated, chirpsResp{
		Id:        chirp.ID,
		CreatedAt: chirp.CreatedAt,
		UpdatedAt: chirp.UpdatedAt,
		Body:      chirp.Body,
		UserID:    chirp.UserID,
	})
}

func ReplaceProfanity(text string) string {
	var wordsToFilter = []string{
		"kerfuffle",
		"sharbert",
		"fornax",
	}

	fields := strings.Fields(text)

	for i := range fields {
		if slices.Contains(wordsToFilter, strings.ToLower(fields[i])) {
			fields[i] = "****"
		}
	}

	return strings.Join(fields, " ")
}
