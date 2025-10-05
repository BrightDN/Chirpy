package main

import (
	"encoding/json"
	"net/http"
	"slices"
	"strings"
)

func (cfg *apiConfig) endpointValidateChirp(w http.ResponseWriter, r *http.Request) {

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

	writeJSON(w, http.StatusOK, chirpValidateResp{CleanedBody: cleaned})
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
