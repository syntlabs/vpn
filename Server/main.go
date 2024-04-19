package main

import (
	"log"
	"net"
	"os"
)

var config_path string = os.Getenv("")

func main() {

	config := config_from(config_path)

	listener, err := net.Listen("tcp", config.listen_address)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Failed to accept connection: %v", err)
			continue
		}

		go handle_connection(conn, config)
	}
}
