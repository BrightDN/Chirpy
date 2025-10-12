package main

import (
	"encoding/json"
	"net/http"
	"slices"
	"strings"

	"github.com/BrightDN/Chirpy/internal/auth"
	"github.com/BrightDN/Chirpy/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) endpointCreateChirp(w http.ResponseWriter, r *http.Request) {
	var params params

	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON")
		return
	}

	bt, err := auth.GetBearerToken(r.Header)

	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}
	user, err := auth.ValidateJWT(bt, cfg.Secret)

	if err != nil {
		writeError(w, http.StatusUnauthorized, err.Error())
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
		UserID: user,
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
