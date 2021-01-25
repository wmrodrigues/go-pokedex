package controllers

import (
	internalHttp "github.com/wmrodrigues/go-pokedex/internal/http"
	"net/http"
)

// HomeIndexHandler
func HomeIndexHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	err = internalHttp.Render("/web/templates/heartbeat/index.html", w)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
