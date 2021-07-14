package httpserver

import (
	"io"
	"net/http"
)

func StartServer() {
	h1 := func(w http.ResponseWriter, _ *http.Request) {
		println("server recieved request on \\")
		io.WriteString(w, "Hello from \\!\n")
	}
	h2 := func(w http.ResponseWriter, request *http.Request) {
		println("server recieved request on \\echo")
		body, _ := io.ReadAll(request.Body)
		io.WriteString(w, "Hello from \\echo!\n")
		w.Write(body)
	}

	http.HandleFunc("/", h1)
	http.HandleFunc("/echo", h2)

	go http.ListenAndServe(":8080", nil)
}