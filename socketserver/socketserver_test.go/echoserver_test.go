package socketserver_test

import (
	"github.com/njayp/toyserver/socketserver"
	"github.com/njayp/toyserver/socketclient"
	"testing"
)

func TestEcho(t *testing.T) {
	addr := "localhost:8011"
	socketserver.StartServer(socketserver.EchoHandler{}, addr)
	message := "test message\n" // must end with \n
	resp := socketclient.SocketPing(message, addr)
	if resp != socketserver.ECHOGREETING + message {
		t.Errorf("echo server fail, recieved: " + resp)
	}
}
