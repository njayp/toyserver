package socketserver

import (
	"net"
)

type BaseHandler struct {
}

func (handler BaseHandler) handle(conn net.Conn) {
	message := []byte("hi from baseServer\n")
	conn.Write(message)
	conn.Close()
}
