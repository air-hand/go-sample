package main

import (
	"local.packages/fundamentals"
	"local.packages/web"
)

func main() {
	println("Start server.")
	fundamentals.StringSample("any input")
	web.Serve()
}
