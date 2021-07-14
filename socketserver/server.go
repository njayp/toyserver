package socketserver

import (
	"net"
)

type Server interface {
	reply(net.Conn)
}

func StartServer(server Server) Server {
	listener, err := net.Listen("tcp", ":8080")
	if err == nil {
		go serve(server, listener)
	} else {
		println(err.Error())
	}
	return server
}

func serve(server Server, listener net.Listener) {
	for {
		conn, err := listener.Accept()
		if err == nil {
			go server.reply(conn)
		} else {
			println(err.Error())
		}
	}
}