package main

import (
	//"bufio"
	"fmt"
	"net"

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

func TestHangup() {
	socketserver.StartServer(socketserver.SessionHandler{}, ADDR)
	conn, _ := net.Dial("tcp", ADDR)
	message := "egg"
	for i := 0; i < 100; i++ {
		message += message
	}
	conn.Write([]byte(message))
	conn.Close()
	fmt.Scanln()
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
	socketserver.StartServer(socketserver.SessionHandler{}, ADDR)

	/*
		conn, _ := net.Dial("tcp", ADDR)
		const size = 50000 // 5E4 is fine, 7E4 too large
		bytearray := make([]byte, size)
		//println(bytearray[size - 1])
		n, err := conn.Write(bytearray)
		println(n, err)
		conn.Write([]byte("\n"))


	*/

	var second string
	fmt.Scanln(&second)
}
