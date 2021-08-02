package socketserver

import (
	"bufio"
	"bytes"
	"fmt"
	"net"
	"time"
)

const SESSIONGREETING = "hi from sessionServer: you have sent me %v messages, your last message was: "

type SessionHandler struct {
}

func splitf(data []byte, atEOF bool) (advance int, token []byte, err error) {

	// built-in MaxScanTokenSize = 64 * 1024 bytes
	/*
		const maxbytes = 10000
		if len(data) > maxbytes {
			return 0, nil, bufio.ErrTooLong
		}
	*/

	// "\r\n\r\n" to get header
	if i := bytes.Index(data, []byte("\r\n\r\n")); i >= 0 {
		return i + 1, data[:i+1], nil
	}

	// if client conn closes before message is complete
	// server conn closes immediately rather than timeout
	if atEOF {
		return //0, nil, bufio.ErrFinalToken
	}

	// request more data
	return 0, nil, nil
}

func (handler SessionHandler) handle(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	msgcount := 0

	scanner.Split(splitf)
	conn.SetDeadline(time.Now().Add(time.Second * 10))

	for scanner.Scan() {
		msgcount += 1
		conn.Write([]byte(fmt.Sprintf(SESSIONGREETING, msgcount)))
		conn.Write([]byte(scanner.Text()))
		conn.SetDeadline(time.Now().Add(time.Second * 10))
		println(scanner.Text())
	}

	fmt.Printf("Session connection closing after receiving %v messages\n", msgcount)
	conn.Close()
}
