package client

import (
	"net"
	//"time"
	"bufio"
)

func Ping() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		println(err.Error())
	}

	//conn.SetDeadline(time.Now().Add(time.Second * 5))

	conn.Write([]byte("repeat after me\n"))
	message, _ := bufio.NewReader(conn).ReadString('\n')
	println(string(message))
}
