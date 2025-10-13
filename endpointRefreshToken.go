package main

import (
	"net/http"
	"time"

	"github.com/BrightDN/Chirpy/internal/auth"
)

func (cfg *apiConfig) endpointRefreshToken(w http.ResponseWriter, r *http.Request) {
	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	tok, err := cfg.Db.GetUserFromToken(r.Context(), token)
	if err != nil {
		writeError(w, http.StatusUnauthorized, err.Error())
	}

	if tok.RevokedAt.Valid {
		writeError(w, http.StatusUnauthorized, "token has been revoked")
		return
	}

	if time.Now().After(tok.ExpiresAt) {
		writeError(w, http.StatusUnauthorized, "token has expired")
		return
	}

	newTok, err := auth.MakeJWT(tok.UserID, cfg.Secret)
	if err != nil {
		writeError(w, http.StatusBadRequest, "something went wrong")
		return
	}

	writeJSON(w, http.StatusOK, tokenResp{Token: newTok})
}
