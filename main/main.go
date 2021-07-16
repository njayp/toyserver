package main

import (
	//"bufio"
	"fmt"
	//"strings"

	"github.com/njayp/toyserver/httpserver"
	"github.com/njayp/toyserver/socketclient"
	"github.com/njayp/toyserver/socketserver"
)

const ADDR = "localhost:8080"

func runSocket() {
	socketserver.StartServer(socketserver.EchoHandler{}, ADDR)
	println(socketclient.SocketPost("here there be dragons\n", ADDR))
}

func runHttp() {
	httpserver.StartServer(ADDR)
	fmt.Scanln()
}

func runSession() {
	socketserver.StartServer(socketserver.SessionHandler{}, ADDR)
	println(socketclient.SessionPost(ADDR, 3))
}

func main() {
	//runHttp()
	//runSocket()

	/*
	const input = "i must not fear\nfear is the mind killer\n"
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		fmt.Printf("%s\n", scanner.Text())
	}
	*/

	runSession()
}
