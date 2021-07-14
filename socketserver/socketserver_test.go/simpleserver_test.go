package socketserver_test

import (
	"github.com/njayp/toyserver/socketserver"
	"github.com/njayp/toyserver/socketclient"
	"testing"
)

func TestSimple(t *testing.T) {
	addr := "localhost:8010"
	socketserver.StartServer(socketserver.SimpleHandler{}, addr)
	resp := socketclient.SocketPost("", addr)
	if resp != socketserver.SIMPLEGREETING {
		t.Errorf("basic server fail, recieved: " + resp)
	}
}
