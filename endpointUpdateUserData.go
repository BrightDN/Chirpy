package main

import (
	"encoding/json"
	"net/http"

	"github.com/BrightDN/Chirpy/internal/auth"
	"github.com/BrightDN/Chirpy/internal/database"
)

func (cfg *apiConfig) endpointUpdateUserData(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	t, err := auth.GetBearerToken(r.Header)
	if err != nil {
		writeError(w, http.StatusUnauthorized, err.Error())
		return
	}

	user, err := auth.ValidateJWT(t, cfg.Secret)
	if err != nil {
		writeError(w, http.StatusUnauthorized, err.Error())
		return
	}

	var p params
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON")
		return
	}

	hp, err := auth.HashPassword(p.Password)
	if err != nil {
		writeError(w, http.StatusUnauthorized, err.Error())
		return
	}

	dbResp, err := cfg.Db.AlterUserData(r.Context(), database.AlterUserDataParams{
		Email:          p.Email,
		HashedPassword: hp,
		ID:             user,
	})

	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeJSON(w, http.StatusOK, userResp{
		Id:        dbResp.ID,
		Email:     dbResp.Email,
		CreatedAt: dbResp.CreatedAt,
		UpdatedAt: dbResp.UpdatedAt,
	})
}
