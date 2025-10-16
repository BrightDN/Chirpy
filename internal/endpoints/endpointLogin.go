package endpoints

import (
	"encoding/json"
	"net/http"

	"github.com/BrightDN/Chirpy/internal/auth"
	"github.com/BrightDN/Chirpy/internal/database"
	"github.com/BrightDN/Chirpy/internal/jsonConfig"
)

func (cfg *ApiConfig) EndpointLogin(w http.ResponseWriter, r *http.Request) {
	var p jsonConfig.Params
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		jsonConfig.WriteError(w, http.StatusBadRequest, "invalid JSON")
		return
	}

	user, err := cfg.Db.GetUser(r.Context(), p.Email)
	if err != nil {
		jsonConfig.WriteError(w, http.StatusInternalServerError, "something went wrong processing the request, please try again")
		return
	}

	isSame, err := auth.ComparePasswordHash(p.Password, user.HashedPassword)
	if err != nil {
		jsonConfig.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if !isSame {
		jsonConfig.WriteError(w, http.StatusUnauthorized, "The given email or password does not match")
		return
	}

	tok, err := auth.MakeJWT(user.ID, cfg.Secret)
	if err != nil {
		jsonConfig.WriteError(w, http.StatusBadRequest, "Something went wrong")
		return
	}

	rt, err := auth.MakeRefreshToken()
	if err != nil {
		jsonConfig.WriteError(w, http.StatusBadRequest, "something went wrong")
		return
	}

	refreshTok, err := cfg.Db.CreateRefreshToken(r.Context(), database.CreateRefreshTokenParams{
		Token:  rt,
		UserID: user.ID,
	})
	if err != nil {
		jsonConfig.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	jsonConfig.WriteJSON(w, http.StatusOK, jsonConfig.UserResp{
		Id:           user.ID,
		Email:        user.Email,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
		AuthToken:    tok,
		RefreshToken: refreshTok.Token,
		IsChirpyRed:  user.IsChirpyRed,
	})
}
