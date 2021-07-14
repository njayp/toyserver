package socketserver

import (
	"net"
)

type Handler interface {
	handle(net.Conn)
}

func StartServer(handler Handler) {
	listener, err := net.Listen("tcp", ":8080")
	if err == nil {
		go serve(handler, listener)
	} else {
		println(err.Error())
	}
}

func serve(handler Handler, listener net.Listener) {
	for {
		conn, err := listener.Accept()
		if err == nil {
			go handler.handle(conn)
		} else {
			println(err.Error())
		}
	}
}