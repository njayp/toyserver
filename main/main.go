package main

import (
	"fmt"

	"github.com/njayp/toyserver/socketclient"
	"github.com/njayp/toyserver/socketserver"
	"github.com/njayp/toyserver/httpserver"
)

func runSocket() {
	addr := "localhost:8080"
	socketserver.StartServer(socketserver.EchoHandler{}, addr)
	println(socketclient.SocketPost("here there be dragons\n", addr))
}

func runHttp() {
	httpserver.StartServer("localhost:8080")
	fmt.Scanln()
}

func main() {
	runHttp()
	//runSocket()
}
