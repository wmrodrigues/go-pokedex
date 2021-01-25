package controllers

import (
	internalHttp "github.com/wmrodrigues/go-pokedex/internal/http"
	"net/http"
)

// HeartbeatIndexHandler
func HeartbeatIndexHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	err = internalHttp.Render("/web/templates/heartbeat/index.html", w, nil)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
