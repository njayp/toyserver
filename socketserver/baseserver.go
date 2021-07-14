package socketserver

import (
	"net"
)

type BaseServer struct {
}

func (server BaseServer) reply(conn net.Conn) {
	message := []byte("hi from baseServer\n")
	conn.Write(message)
	conn.Close()
}
