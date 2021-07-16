package socketserver

import (
	"bufio"
	"fmt"
	"net"
)

const SESSIONGREETING = "hi from sessionServer: you have sent me %v messages, your last message was: "

type SessionHandler struct {
}

func (handler SessionHandler) handle(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	msgcount := 0

	for scanner.Scan() {
		msgcount += 1
		conn.Write([]byte(fmt.Sprintf(SESSIONGREETING, msgcount)))
		// TODO custom scanner function that does not delete "\n"
		conn.Write([]byte(scanner.Text() + "\n"))
	}

	//fmt.Printf("Session connection closing after receiving %v messages\n", msgcount)
	conn.Close()
}
