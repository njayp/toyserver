package socketserver

import (
	"net"
	"bufio"
)

type EchoHandler struct {
}

func (handler EchoHandler) handle(conn net.Conn) {
	message, _ := bufio.NewReader(conn).ReadBytes('\n')
	conn.Write([]byte("hi from echoServer: "))
	conn.Write(message)
	conn.Close()
}
