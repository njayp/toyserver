package socketserver_test

import (
	"fmt"
	//"net"
	"testing"
	//"time"

	"github.com/njayp/toyserver/socketclient"
	"github.com/njayp/toyserver/socketserver"
)

const ADDR = "localhost:8012"

func TestSession(t *testing.T) {
	const nomsg = 20
	
	var expectedmessages string
	for i := 1; i <= nomsg; i++ {
		expectedmessages += fmt.Sprintf(socketserver.SESSIONGREETING + socketclient.SESSIONMESSAGE, i, i)
	}

	socketserver.StartServer(socketserver.SessionHandler{}, ADDR)
	messages := socketclient.SessionPost(ADDR, nomsg)
	if messages != expectedmessages {
		t.Errorf("session server fail, client recieved: " + messages)
	}
}

/*
func TestValidHangup(t *testing.T) {
	socketserver.StartServer(socketserver.SessionHandler{}, ADDR)
	conn, _ := net.Dial("tcp", ADDR)
	conn.Write([]byte("egg\n"))
	conn.Close()
}

func TestInvalidHangup(t *testing.T) {
	socketserver.StartServer(socketserver.SessionHandler{}, ADDR)
	conn, _ := net.Dial("tcp", ADDR)
	conn.Write([]byte("egg"))
	conn.Close()
}


func TestTimeout(t *testing.T) {
	socketserver.StartServer(socketserver.SessionHandler{}, ADDR)
	conn, _ := net.Dial("tcp", ADDR)
	time.Sleep(time.Second*20)
	conn.Close()
}
*/