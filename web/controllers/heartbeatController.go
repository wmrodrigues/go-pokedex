package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

// IndexHandler
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	wd, _ := os.Getwd()
	tplPath := fmt.Sprintf("%s/web/templates/heartbeat/index.html", wd)
	templates := template.Must(template.ParseFiles(tplPath))
	err := templates.ExecuteTemplate(w, "index.html", nil)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
