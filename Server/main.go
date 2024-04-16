package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"sync"
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

func handle_connection(conn net.Conn, config *Server_config) {
	defer conn.Close()

	// Read the SOCKS5 greeting from the client
	buf := make([]byte, 3)
	_, err := io.ReadFull(conn, buf)
	if err != nil {
		return
	}

	// Check the SOCKS5 version and number of authentication methods
	if buf[0] != 5 || buf[1] != 1 {
		return
	}

	// Read the authentication method selected by the client
	_, err = io.ReadFull(conn, buf[:1])
	if err != nil {
		return
	}

	// Check if the client selected the correct authentication method
	if buf[0] != config.auth_method {
		return
	}

	// Send the SOCKS5 greeting back to the client
	_, err = conn.Write([]byte{5, config.auth_method})
	if err != nil {
		return
	}

	// Read the SOCKS5 request from the client
	buf = make([]byte, 10)
	_, err = io.ReadFull(conn, buf)
	if err != nil {
		return
	}

	// Check the SOCKS5 version and command
	if buf[0] != 5 || buf[1] != 1 {
		return
	}

	usernameLen := int(buf[0])
    usernameBuf := make([]byte, usernameLen)
    _, err = io.ReadFull(conn, usernameBuf)
    if err != nil {
        return
    }

    _, err = io.ReadFull(conn, buf[:1])
    if err != nil {
        return
    }

    passwordLen := int(buf[0])
    passwordBuf := make([]byte, passwordLen)
    _, err = io.ReadFull(conn, passwordBuf)
    if err != nil {
        return
    }

    // Check if the provided username and password are valid
    if string(usernameBuf) != config.auth_username || string(passwordBuf) != config.auth_password {
        return
    }

	// Read the destination address and port
	addrType := buf[3]
	var addr string
	var port uint16
	switch addrType {
	case 1: // IPv4 address
		addr = fmt.Sprintf("%d.%d.%d.%d", buf[4], buf[5], buf[6], buf[7])
		port = binary.BigEndian.Uint16(buf[8:10])
	case 3: // Domain name
		domainLen := int(buf[4])
		domainBuf := make([]byte, domainLen)
		_, err = io.ReadFull(conn, domainBuf)
		if err != nil {
			log.Printf("Failed to read domain name: %v", err)
			return
		}
		addr = string(domainBuf)
		port = binary.BigEndian.Uint16(buf[5+domainLen : 7+domainLen])
	case 4: // IPv6 address
		addr = fmt.Sprintf("%x:%x:%x:%x:%x:%x:%x:%x", buf[4], buf[5], buf[6], buf[7], buf[8], buf[9], buf[10], buf[11])
		port = binary.BigEndian.Uint16(buf[12:14])
	default:
		log.Println("Unsupported address type")
		return
	}

	if port == 80 && buf[0] == 0 {
		log.Println("User is trying to connect to port 80 without authentication. Connection rejected.")
		_, err = conn.Write([]byte{5, 1, 0, 8, 0, 0, 0, 0, 0, 0}) // Send a SOCKS5 response with "connection not allowed by ruleset" error
		if err != nil {
			log.Printf("Failed to send SOCKS5 response: %v", err)
			return
		}
		return
	}
	// Connect to the destination
	destConn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", addr, port))
	if err != nil {
		return
	}
	defer destConn.Close()

	// Send the SOCKS5 response to the client
	_, err = conn.Write([]byte{5, 0, 0, 1, 0, 0, 0, 0, 0, 0})
	if err != nil {
		return
	}

	// Relay data between the client and the destination
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		io.Copy(destConn, conn)
		destConn.Close()
		wg.Done()
	}()
	go func() {
		io.Copy(conn, destConn)
		conn.Close()
		wg.Done()
	}()
	wg.Wait()
}
