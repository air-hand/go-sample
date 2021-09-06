package web

import (
	"html/template"
	"log"
	"net/http"
)

func renderTemplate(w http.ResponseWriter, tmpl string) {
	// TODO: I want to use path as relative... how?
	parsed, _ := template.ParseFiles("/opt/app/src/web/" + "templates/" + tmpl)
	err := parsed.Execute(w, nil)
	if err != nil {
		log.Println("error parsing :", err)
		return
	}
}
