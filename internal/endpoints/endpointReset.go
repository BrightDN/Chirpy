package endpoints

import (
	"net/http"

	"github.com/BrightDN/Chirpy/internal/jsonConfig"
)

func (cfg *ApiConfig) EndpointReset(w http.ResponseWriter, r *http.Request) {

	if cfg.Platform != "dev" {
		jsonConfig.WriteError(w, http.StatusForbidden, "Change to DEV mode for this functionality")
		return
	}

	if err := cfg.Db.DeleteUsers(r.Context()); err != nil {
		jsonConfig.WriteError(w, http.StatusInternalServerError, "An issue occured while processing your request")
		return
	}

	cfg.fileserverHits.Store(0)

	w.WriteHeader(http.StatusOK)

}
