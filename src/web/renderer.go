package web

import (
	"bytes"
	"embed"
	"fmt"
	"html/template"
	"io/fs"
	"log"
)

//go:embed templates
var templates embed.FS

type TemplateCache struct {
	caches map[string]*template.Template
}

func newTemplateCache() TemplateCache {
	cache := TemplateCache{
		caches: make(map[string]*template.Template),
	}
	return cache
}

func (cache *TemplateCache) Get(tmpl string) *template.Template {
	if t, exists := cache.caches[tmpl]; exists {
		return t
	}
	return nil
}

func (cache *TemplateCache) Warmup(templates embed.FS) {
	cache.caches = make(map[string]*template.Template)
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
		cache.caches[tmpl] = t
	}
}

type TemplateRenderer struct {
	UseCache      bool
	templateCache *TemplateCache
}

func (renderer *TemplateRenderer) RenderToBuffer(tmpl string, data interface{}) *bytes.Buffer {
	if !renderer.UseCache {
		panic("todo")
	}
	tmpl = fmt.Sprintf("templates/%s", tmpl)
	t := renderer.templateCache.Get(tmpl)
	if t == nil {
		log.Fatalf("The cache has not been found : %s", tmpl)
		return nil
	}
	buffer := new(bytes.Buffer)
	err := t.Execute(buffer, data)
	if err != nil {
		log.Fatalln("Failed to execute template:", err)
		return nil
	}
	return buffer
}

func NewTemplateRenderer(useCache bool) *TemplateRenderer {
	templateCache := newTemplateCache()
	if useCache {
		templateCache.Warmup(templates)
	}
	renderer := TemplateRenderer{
		UseCache:      useCache,
		templateCache: &templateCache,
	}
	return &renderer
}
