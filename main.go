package main

import (
	"github.com/rchargel/goblog/web"
)

func main() {
	server := web.NewServer()
	server.Start()
}
