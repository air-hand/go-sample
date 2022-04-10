package main

import (
	"local.packages/web"
)

func main() {
	println("Start server.")
	server := web.NewServer(80)
	server.Serve()
}
