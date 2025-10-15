package main

import (
	"encoding/json"
	"net/http"
)

func (cfg *apiConfig) endpointUpgradeWebhook(w http.ResponseWriter, r *http.Request) {
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
