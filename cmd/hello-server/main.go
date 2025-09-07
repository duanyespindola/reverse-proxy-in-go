package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	hello_server "github.com/duanyespindola/reverse-proxy-in-go/internal/hello-server"
)

func main() {
	// 1. Define and parse command-line flags
	serverNumber := flag.Int("number", 1, "The server number, used in the response message.")
	port := flag.Int("port", 8081, "The port the server will listen on.")
	flag.Parse()

	// 2. Create the handler using our internal package
	handler := hello_server.HelloServerHandler(*serverNumber)

	// 3. Configure the server address
	addr := fmt.Sprintf(":%d", *port)
	log.Printf("Starting hello server #%d on address %s", *serverNumber, addr)

	// 4. Start the HTTP server
	// http.ListenAndServe starts the server and blocks forever until it fails.
	// If it fails (e.g., port is already in use), it returns an error.
	// We wrap it in log.Fatal so that if an error occurs, the program will
	// print the error message and exit immediately.
	if err := http.ListenAndServe(addr, handler); err != nil {
		log.Fatalf("could not start server: %v", err)
	}
}
