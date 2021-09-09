package web

import (
	"net/http"
	"time"

	"github.com/lestrrat-go/strftime"
)

type Handler struct {
	renderer *TemplateRenderer
}

func (handler *Handler) Home(w http.ResponseWriter, r *http.Request) {
	buffer := handler.renderer.RenderToBuffer("home.page.tmpl", struct {
		Title string
	}{
		Title: "Home",
	})
	if buffer == nil {
		return
	}
	buffer.WriteTo(w)
}

func (handler *Handler) About(w http.ResponseWriter, r *http.Request) {
	buffer := handler.renderer.RenderToBuffer("about.page.tmpl", struct {
		Title string
	}{
		Title: "About",
	})
	if buffer == nil {
		return
	}
	buffer.WriteTo(w)
}

func (handler *Handler) NowTime(w http.ResponseWriter, r *http.Request) {
	p, err := strftime.New("%Y-%m-%d %H:%M:%S")
	if err != nil {
		return
	}
	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return
	}
	now_string := p.FormatString(time.Now().In(loc))

	buffer := handler.renderer.RenderToBuffer("now.page.tmpl", struct {
		Title string
		Time  string
	}{
		Title: "Now",
		Time:  now_string,
	})
	if buffer == nil {
		return
	}
	buffer.WriteTo(w)
}
