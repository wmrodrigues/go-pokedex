package main

import (
	"github.com/wmrodrigues/go-pokedex/web"
	"log"
	"net/http"
)

func main() {
	log.Println("starting project...")

	var app web.App
	app.RegisterRoutes()

	log.Fatal(http.ListenAndServe(":8087", app.Router))
}
