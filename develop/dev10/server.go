package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
	defer listener.Close()

	log.Println("Server started, listening on port 8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %v", err)
			continue
		}
		go func(c net.Conn) {
			defer c.Close()
			log.Printf("Client connected: %s", c.RemoteAddr())
			fmt.Fprintln(c, "What's your name?")
			var name string
			fmt.Fscanln(c, &name)
			fmt.Fprintf(c, "Your name is %s!\n", name)
			log.Printf("Client %s disconnected", c.RemoteAddr())
		}(conn)
	}
}
