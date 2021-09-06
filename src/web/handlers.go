package web

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/lestrrat-go/strftime"
)

func home(w http.ResponseWriter, r *http.Request) {
	t, _ := loadTemplate("home.page.tmpl")
	err := t.Execute(w, struct {
		Title string
	}{
		Title: "Home",
	})
	if err != nil {
		log.Fatalln("Failed to renderer template", err)
	}
}

func about(w http.ResponseWriter, r *http.Request) {
	t, _ := loadTemplate("about.page.tmpl")
	err := t.Execute(w, struct {
		Title string
	}{
		Title: "About",
	})
	if err != nil {
		log.Fatalln("Failed to renderer template", err)
	}
}

func now_time(w http.ResponseWriter, r *http.Request) {
	p, err := strftime.New("%Y-%m-%d %H:%M:%S")
	if err != nil {
		return
	}
	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return
	}
	fmt.Fprintf(w, "Now %s", p.FormatString(time.Now().In(loc)))
}
