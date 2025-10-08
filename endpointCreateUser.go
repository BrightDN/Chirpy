package main

import (
	"encoding/json"
	"net/http"

	"github.com/BrightDN/Chirpy/internal/auth"
	"github.com/BrightDN/Chirpy/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) endpointCreateUser(w http.ResponseWriter, r *http.Request) {
	var p params
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON")
		return
	}

	if p.Email == "" {
		writeError(w, http.StatusBadRequest, "email is required")
		return
	}

	hash, err := auth.HashPassword(p.Password)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Something went wrong, please try again")
	}
	user, err := cfg.Db.CreateUser(r.Context(), database.CreateUserParams{
		ID:             uuid.New(),
		Email:          p.Email,
		HashedPassword: hash,
	})
	if err != nil {
		writeError(w, http.StatusInternalServerError, "could not create user")
		return
	}

	writeJSON(w, http.StatusCreated, userResp{
		Id:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Email:     user.Email,
	})
}
