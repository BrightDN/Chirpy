package main

import (
	"encoding/json"
	"net/http"

	"github.com/BrightDN/Chirpy/internal/auth"
)

func (cfg *apiConfig) endpointUpgradeWebhook(w http.ResponseWriter, r *http.Request) {

	ak, err := auth.GetApiKey(r.Header)
	if err != nil {
		writeError(w, http.StatusUnauthorized, err.Error())
		return
	}
	if ak != cfg.PolkaKey {
		writeError(w, http.StatusUnauthorized, "Unauthorized request")
		return
	}

	var whr upgradeHookResp
	if err := json.NewDecoder(r.Body).Decode(&whr); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON")
		return
	}

	if whr.Event != "user.upgraded" {
		writeJSON(w, http.StatusNoContent, "")
		return
	}

	if err := cfg.Db.UpgradeUserRed(r.Context(), whr.Data.UserId); err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeJSON(w, http.StatusNoContent, "")
}
