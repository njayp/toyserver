package client

import (
	"net"
	"bufio"
)

func Ping() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err == nil {
		conn.Write([]byte("here there be dragons\n"))
		message, _ := bufio.NewReader(conn).ReadString('\n')
		println(string(message))
	} else {
		println(err.Error())
	}
}
