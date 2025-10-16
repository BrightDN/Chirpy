package endpoints

import (
	"net/http"
	"sort"

	"github.com/BrightDN/Chirpy/internal/database"
	"github.com/BrightDN/Chirpy/internal/jsonConfig"
	"github.com/google/uuid"
)

func (cfg *ApiConfig) EndpointGetChirps(w http.ResponseWriter, r *http.Request) {
	sid := r.URL.Query().Get("author_id")
	ssort := r.URL.Query().Get("sort")

	var chirps []database.Chirp
	var err error
	if sid == "" {
		chirps, err = cfg.Db.GetAllChirps(r.Context())
		if err != nil {
			jsonConfig.WriteError(w, http.StatusInternalServerError, "An error occured handling your request")
			return
		}
	} else {
		s, err := uuid.Parse(sid)
		if err != nil {
			jsonConfig.WriteError(w, http.StatusBadRequest, "Invalid ID format")
			return
		}
		chirps, err = cfg.Db.GetAllChirpsFromAuthor(r.Context(), s)
		if err != nil {
			jsonConfig.WriteError(w, http.StatusInternalServerError, "An error occured handling your request")
			return
		}
	}
	if ssort == "desc" {
		sort.Slice(chirps, func(i, j int) bool {
			return i > j
		})
	}

	out := make([]jsonConfig.ChirpsResp, len(chirps))
	for i, r := range chirps {
		out[i] = jsonConfig.ChirpsResp{
			Id:        r.ID,
			CreatedAt: r.CreatedAt,
			UpdatedAt: r.UpdatedAt,
			Body:      r.Body,
			UserID:    r.UserID,
		}
	}

	jsonConfig.WriteJSON(w, http.StatusOK, out)
}

func (cfg *ApiConfig) EndpointGetChirp(w http.ResponseWriter, r *http.Request) {
	chirpID := r.PathValue("chirpID")
	if chirpID == "" {
		jsonConfig.WriteError(w, http.StatusBadRequest, "Include a chirpID")
		return
	}

	chirpUUID, err := uuid.Parse(chirpID)
	if err != nil {
		jsonConfig.WriteError(w, http.StatusBadRequest, "Invalid ID format")
		return
	}

	chirp, err := cfg.Db.GetChirp(r.Context(), chirpUUID)
	if err != nil {
		jsonConfig.WriteError(w, http.StatusNotFound, "No chirp with this ID exists")
		return
	}

	jsonConfig.WriteJSON(w, http.StatusOK, jsonConfig.ChirpsResp{
		Id:        chirp.ID,
		CreatedAt: chirp.CreatedAt,
		UpdatedAt: chirp.UpdatedAt,
		Body:      chirp.Body,
		UserID:    chirp.UserID,
	})
}
