package main

import (
	"flag"
)

func main() {

	var port string

	flag.StringVar(&port, "port", "8099", "port to listen on")
	flag.Parse()

	server := NewServer(port)
	StartServer(server)
}
