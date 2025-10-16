package endpoints

import (
	"net/http"
	"time"

	"github.com/BrightDN/Chirpy/internal/auth"
	"github.com/BrightDN/Chirpy/internal/jsonConfig"
)

func (cfg *ApiConfig) EndpointRefreshToken(w http.ResponseWriter, r *http.Request) {
	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		jsonConfig.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	tok, err := cfg.Db.GetUserFromToken(r.Context(), token)
	if err != nil {
		jsonConfig.WriteError(w, http.StatusUnauthorized, err.Error())
	}

	if tok.RevokedAt.Valid {
		jsonConfig.WriteError(w, http.StatusUnauthorized, "token has been revoked")
		return
	}

	if time.Now().After(tok.ExpiresAt) {
		jsonConfig.WriteError(w, http.StatusUnauthorized, "token has expired")
		return
	}

	newTok, err := auth.MakeJWT(tok.UserID, cfg.Secret)
	if err != nil {
		jsonConfig.WriteError(w, http.StatusBadRequest, "something went wrong")
		return
	}

	jsonConfig.WriteJSON(w, http.StatusOK, jsonConfig.TokenResp{Token: newTok})
}
