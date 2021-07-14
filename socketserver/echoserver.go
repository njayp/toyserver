package socketserver

import (
	"net"
)

type EchoServer struct {
	BaseServer
}

func (server EchoServer) reply(listener net.Conn) {

}
