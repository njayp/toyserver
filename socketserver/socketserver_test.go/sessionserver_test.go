package socketserver_test

import (
	"github.com/njayp/toyserver/socketserver"
	"github.com/njayp/toyserver/socketclient"
	"testing"
	"fmt"
)

func TestSession(t *testing.T) {
	const addr, nomsg = "localhost:8012", 6
	
	var expectedmessages string
	for i := 1; i <= nomsg; i++ {
		expectedmessages += fmt.Sprintf(socketserver.SESSIONGREETING + socketclient.SESSIONMESSAGE, i, i)
	}

	socketserver.StartServer(socketserver.SessionHandler{}, addr)
	messages := socketclient.SessionPost(addr, nomsg)
	if messages != expectedmessages {
		t.Errorf("session server fail, client recieved: " + messages)
	}
}
