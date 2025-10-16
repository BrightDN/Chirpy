package endpoints

import (
	"net/http"

	"github.com/BrightDN/Chirpy/internal/auth"
	"github.com/BrightDN/Chirpy/internal/jsonConfig"
)

func (cfg *ApiConfig) EndpointRevokeToken(w http.ResponseWriter, r *http.Request) {

	tok, err := auth.GetBearerToken(r.Header)
	if err != nil {
		jsonConfig.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := cfg.Db.RevokeToken(r.Context(), tok); err != nil {
		jsonConfig.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}
	jsonConfig.WriteJSON(w, http.StatusNoContent, "")
}
