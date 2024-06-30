package main

import (
	"fmt"
	"go-server/api"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from go server"))
	}

func main() {

	server := http.NewServeMux()

	// adding a handler
	server.HandleFunc("/", sayHello)
	server.HandleFunc("/api/ex", api.Get)
	server.HandleFunc("/api/ex/new", api.Post)

	// starting the server
	err := http.ListenAndServe(":8000", server)

	if err != nil {
		fmt.Println("error while opening the server")
	} else {
		fmt.Println("Server started on port 8000")
	}
}