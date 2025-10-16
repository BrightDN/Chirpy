package endpoints

import (
	"encoding/json"
	"net/http"

	"github.com/BrightDN/Chirpy/internal/auth"
	"github.com/BrightDN/Chirpy/internal/jsonConfig"
)

func (cfg *ApiConfig) EndpointUpgradeWebhook(w http.ResponseWriter, r *http.Request) {

	ak, err := auth.GetApiKey(r.Header)
	if err != nil {
		jsonConfig.WriteError(w, http.StatusUnauthorized, err.Error())
		return
	}
	if ak != cfg.PolkaKey {
		jsonConfig.WriteError(w, http.StatusUnauthorized, "Unauthorized request")
		return
	}

	var whr jsonConfig.UpgradeHookResp
	if err := json.NewDecoder(r.Body).Decode(&whr); err != nil {
		jsonConfig.WriteError(w, http.StatusBadRequest, "invalid JSON")
		return
	}

	if whr.Event != "user.upgraded" {
		jsonConfig.WriteJSON(w, http.StatusNoContent, "")
		return
	}

	if err := cfg.Db.UpgradeUserRed(r.Context(), whr.Data.UserId); err != nil {
		jsonConfig.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	jsonConfig.WriteJSON(w, http.StatusNoContent, "")
}
