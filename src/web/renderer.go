package web

import (
	"html/template"
	"log"
)

var caches = map[string]*template.Template{}

func loadTemplate(tmpl string) (*template.Template, error) {
	if t, exists := caches[tmpl]; exists {
		log.Printf("The cache found : %s", tmpl)
		return t, nil
	}
	// TODO: I want to use path as relative... how?
	t, err := template.ParseFiles(
		"/opt/app/src/web/" + "templates/" + tmpl,
	)
	if err != nil {
		return t, err
	}
	t, err = t.ParseGlob(
		"/opt/app/src/web/templates/layouts/*.tmpl",
	)
	if err != nil {
		return t, err
	}

	caches[tmpl] = t
	return t, nil
}
