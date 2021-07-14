package socketserver

import (
	"net"
)

type Server interface {
	Start() Server
	serve(net.Listener)
	reply(net.Conn)
}

type BaseServer struct {
}

func (server BaseServer) Start() Server {
	listener, _ := net.Listen("tcp", ":8080")
	go server.serve(listener)
	return server
}

func (server BaseServer) serve(listener net.Listener) {
	for {
		conn, err := listener.Accept()
		if err == nil {
			go server.reply(conn)
		} else {
			println(err.Error())
		}
	}
}

func (server BaseServer) reply(listener net.Conn) {}