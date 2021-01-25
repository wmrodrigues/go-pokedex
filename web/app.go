package web

import (
	"github.com/gorilla/mux"
	"github.com/wmrodrigues/go-pokedex/web/controllers"
	"net/http"
)

type App struct {
	Router     *mux.Router
}

func (a *App) RegisterRoutes() {
	a.Router = mux.NewRouter()

	// static file handlers
	a.Router.PathPrefix("/semantic/").Handler(http.StripPrefix("/semantic/", http.FileServer(http.Dir("./semantic"))))
	a.Router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))

	// route handlers
	a.Router.Handle("/heartbeat", http.HandlerFunc(controllers.HeartbeatIndexHandler)).Methods(http.MethodGet)
	a.Router.Handle("/home", http.HandlerFunc(controllers.HomeIndexHandler)).Methods(http.MethodGet)
}