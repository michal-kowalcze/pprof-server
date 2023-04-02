package server

import (
	"github.com/michal-kowalcze/pprof-server/internal/driver"
	"net/http"
)

type profProvider interface {
	GetProfile(id string) ([]byte, error)
}

func ProfHandler(provider profProvider) http.HandlerFunc {
	return driver.ProfHandler(provider)
}
