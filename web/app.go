package web

import (
	"github.com/gorilla/mux"
	"github.com/wmrodrigues/web/controllers"
	"net/http"
)

type App struct {
	Router     *mux.Router
}

func (a *App) RegisterRoutes() {
	a.Router = mux.NewRouter()

	a.Router.Handle("/heartbeat", http.HandlerFunc(controllers.IndexHandler)).Methods(http.MethodGet)
}