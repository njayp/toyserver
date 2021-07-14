package socketserver

import (
	"net"
	"bufio"
)

type EchoServer struct {
	BaseServer
}

func (server EchoServer) reply(conn net.Conn) {
	message, _ := bufio.NewReader(conn).ReadBytes('\n')
	conn.Write([]byte("hi from echoServer: "))
	conn.Write(message)
	conn.Close()
}
