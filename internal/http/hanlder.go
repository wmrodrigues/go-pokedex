package http

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

func Render(templatePath string, w http.ResponseWriter) error {
	wd, _ := os.Getwd()
	tplPath := fmt.Sprintf("%s%s", wd, templatePath)
	templates := template.Must(template.ParseFiles(tplPath))
	err := templates.ExecuteTemplate(w, "index.html", nil)
	return err
}
