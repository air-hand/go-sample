package web

import (
	"fmt"
	"net/http"
	"time"

	"github.com/lestrrat-go/strftime"
)

func home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.page.tmpl")
}

func about(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.page.tmpl")
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
