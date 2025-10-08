package main

import (
	"encoding/json"
	"net/http"

	"github.com/BrightDN/Chirpy/internal/auth"
)

func (cfg *apiConfig) endpointLogin(w http.ResponseWriter, r *http.Request) {
	var p params
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON")
		return
	}

	user, err := cfg.Db.GetUser(r.Context(), p.Email)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "something went wrong processing the request, please try again")
		return
	}

	isSame, err := auth.ComparePasswordHash(p.Password, user.HashedPassword)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if !isSame {
		writeError(w, http.StatusUnauthorized, "The given email or password does not match")
	}

	writeJSON(w, http.StatusOK, userResp{
		Id:        user.ID,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	})
}
