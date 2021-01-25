package http

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

func Render(templatePath string, w http.ResponseWriter, data interface{}) error {
	wd, _ := os.Getwd()
	tplPath := fmt.Sprintf("%s%s", wd, templatePath)
	templates := template.Must(template.ParseFiles(tplPath))
	err := templates.ExecuteTemplate(w, "index.html", data)
	return err
}
