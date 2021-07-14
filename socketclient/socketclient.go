package socketclient

import (
	"bufio"
	"net"
)

func SocketPost(message string, addr string) string {
	conn, err := net.Dial("tcp", addr)
	if err == nil {
		conn.Write([]byte(message))
		message, _ := bufio.NewReader(conn).ReadString('\n')
		return message
	} else {
		return err.Error()
	}
}
