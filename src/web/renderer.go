package web

import (
	"embed"
	"fmt"
	"html/template"
	"io/fs"
	"log"
)

var caches = map[string]*template.Template{}

//go:embed templates
var templates embed.FS

func warmupTemplatesCache() {
	tmpls, err := fs.Glob(templates, "templates/*.tmpl")
	if err != nil {
		log.Fatalln("glob err:", err)
		return
	}

	for _, tmpl := range tmpls {
		t, err := template.ParseFS(
			templates, tmpl,
		)
		if err != nil {
			log.Fatalln("Parse error:", err)
			return
		}
		t, err = t.ParseFS(
			templates, "templates/layouts/*.tmpl",
		)
		if err != nil {
			log.Fatalln("Parse layouts error:", err)
			return
		}
		caches[tmpl] = t
	}
}

func findTemplate(tmpl string) *template.Template {
	tmpl = fmt.Sprintf("templates/%s", tmpl)
	if t, exists := caches[tmpl]; exists {
		return t
	}
	log.Fatalf("The cache has not been found : %s", tmpl)
	return nil
}
