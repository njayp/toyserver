package socketserver

import (
	"net"
	"bufio"
)

const ECHOGREETING = "hi from echoServer: "

type EchoHandler struct {
}

func (handler EchoHandler) handle(conn net.Conn) {
	message, _ := bufio.NewReader(conn).ReadBytes('\n')
	conn.Write([]byte(ECHOGREETING))
	conn.Write(message)
	conn.Close()
}
