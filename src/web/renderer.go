package web

import (
	"embed"
	"html/template"
	"log"
)

var caches = map[string]*template.Template{}

//go:embed templates
var fs embed.FS

func loadTemplate(tmpl string) (*template.Template, error) {
	if t, exists := caches[tmpl]; exists {
		log.Printf("The cache found : %s", tmpl)
		return t, nil
	}
	t, err := template.ParseFS(
		fs, "templates/"+tmpl,
	)
	if err != nil {
		return t, err
	}
	t, err = t.ParseFS(
		fs, "templates/layouts/*.tmpl",
	)
	if err != nil {
		return t, err
	}

	caches[tmpl] = t
	return t, nil
}
