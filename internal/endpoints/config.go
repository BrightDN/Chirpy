package endpoints

import (
	"sync/atomic"

	"github.com/BrightDN/Chirpy/internal/database"
)

type ApiConfig struct {
	fileserverHits atomic.Int32
	Db             *database.Queries
	Platform       string
	Secret         string
	PolkaKey       string
}
