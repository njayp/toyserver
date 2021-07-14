package socketserver

import (
	"net"
)

const SIMPLEGREETING = "hi from baseServer\n"

type SimpleHandler struct {
}

func (handler SimpleHandler) handle(conn net.Conn) {
	message := []byte(SIMPLEGREETING)
	conn.Write(message)
	conn.Close()
}