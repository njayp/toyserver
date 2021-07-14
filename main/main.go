package main

import (
	"fmt"

	"github.com/njayp/toyserver/client"
	"github.com/njayp/toyserver/socketserver"
)

func main() {
	socketserver.StartServer(socketserver.EchoServer{})
	for {
		fmt.Scanln()
		client.Ping()
	}
}