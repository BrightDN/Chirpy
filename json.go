package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

func writeError(w http.ResponseWriter, status int, msg string) {
	writeJSON(w, status, apiError{Error: msg})
}

type params struct {
	Body  string `json:"body,omitempty"`
	Email string `json:"email,omitempty"`
}

type chirpValidateResp struct {
	CleanedBody string `json:"cleaned_body,omitempty"`
	Valid       bool   `json:"valid,omitempty"`
}

type createUserResp struct {
	Id        uuid.UUID `json:"id,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	Email     string    `json:"email,omitempty"`
}

type apiError struct {
	Error string `json:"error"`
}
