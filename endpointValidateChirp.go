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
		w.Header().Set("Content-Type", "application/json")
		resp := apiError{Error: "Something went wrong"}
		data, _ := json.Marshal(resp)
		w.WriteHeader(400)
		w.Write(data)
		return
	}

	if len(params.Body) > 140 {
		w.Header().Set("Content-Type", "application/json")
		resp := apiError{Error: "Chirp is too long"}
		data, _ := json.Marshal(resp)
		w.WriteHeader(400)
		w.Write(data)
		return
	}

	cleaned := ReplaceProfanity(params.Body)

	w.Header().Set("Content-Type", "application/json")
	resp := chirpValidateResp{Cleaned_Body: cleaned}
	data, _ := json.Marshal(resp)
	w.WriteHeader(200)
	w.Write(data)
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
