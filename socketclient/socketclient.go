package socketclient

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

// delimiter strategy
// messages sent must end with /n

func SocketPost(message string, addr string) string {
	conn, err := net.Dial("tcp", addr)
	if err == nil {
		conn.Write([]byte(message))
		message, _ := bufio.NewReader(conn).ReadString('\n')
		conn.Close()
		return message
	} else {
		return err.Error()
	}
}

const SESSIONMESSAGE = "This is message number %v\n"

func SessionPost(addr string, nomsg int) string {
	conn, _ := net.Dial("tcp", addr)
	var messages string

	for i := 1; i <= nomsg; i++ {
		conn.Write([]byte(fmt.Sprintf(SESSIONMESSAGE, i)))
		message, _ := bufio.NewReader(conn).ReadString('\n')
		messages += message
		time.Sleep(time.Second)
	}

	conn.Close()
	return messages
}