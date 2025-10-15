package main

import (
	"sync/atomic"

	"github.com/BrightDN/Chirpy/internal/database"
)

type apiConfig struct {
	fileserverHits atomic.Int32
	Db             *database.Queries
	Platform       string
	Secret         string
	PolkaKey       string
}
