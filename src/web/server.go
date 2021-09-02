package web

import (
	"net/http"
)

func Serve() {
	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)
	http.HandleFunc("/now", now_time)
	_ = http.ListenAndServe(":80", nil)
}
