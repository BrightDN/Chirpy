package endpoints

import (
	"encoding/json"
	"net/http"

	"github.com/BrightDN/Chirpy/internal/auth"
	"github.com/BrightDN/Chirpy/internal/database"
	"github.com/BrightDN/Chirpy/internal/jsonConfig"
	"github.com/google/uuid"
)

func (cfg *ApiConfig) EndpointCreateUser(w http.ResponseWriter, r *http.Request) {
	var p jsonConfig.Params
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		jsonConfig.WriteError(w, http.StatusBadRequest, "invalid JSON")
		return
	}

	if p.Email == "" {
		jsonConfig.WriteError(w, http.StatusBadRequest, "email is required")
		return
	}

	hash, err := auth.HashPassword(p.Password)
	if err != nil {
		jsonConfig.WriteError(w, http.StatusInternalServerError, "Something went wrong, please try again")
	}
	user, err := cfg.Db.CreateUser(r.Context(), database.CreateUserParams{
		ID:             uuid.New(),
		Email:          p.Email,
		HashedPassword: hash,
	})
	if err != nil {
		jsonConfig.WriteError(w, http.StatusInternalServerError, "could not create user")
		return
	}

	jsonConfig.WriteJSON(w, http.StatusCreated, jsonConfig.UserResp{
		Id:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Email:     user.Email,
	})
}
