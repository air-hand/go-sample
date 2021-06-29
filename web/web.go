package web

import (
	"fmt"
	"log"
	"net/http"
)

func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hi there, I love %s!", request.URL.Path[1:])
}

func Serve() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":80", nil))
}
