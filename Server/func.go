package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"sync"
)

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
	//
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
	//
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

	if config.subscription_port > 0 && config.subscription_network != "" {
		// Subscription-based connection
		// Create a new listener for the subscription-based connections
		subscriptionListener, err := net.Listen(config.subscription_network, fmt.Sprintf(":%d", config.subscription_port))
		if err != nil {
			log.Printf("Failed to listen for subscription-based connections: %v", err)
			return
		}
		defer subscriptionListener.Close()

		// Accept the subscription-based connection
		subscriptionConn, err := subscriptionListener.Accept()
		if err != nil {
			log.Printf("Failed to accept subscription-based connection: %v", err)
			return
		}
		defer subscriptionConn.Close()

		// Relay data between the client and the subscription-based connection
		var wg sync.WaitGroup
		wg.Add(2)
		go func() {
			io.Copy(subscriptionConn, conn)
			subscriptionConn.Close()
			wg.Done()
		}()
		go func() {
			io.Copy(conn, subscriptionConn)
			conn.Close()
			wg.Done()
		}()
		wg.Wait()
	} else if config.promo_code != "" {
		// Promo code-based connection
		// Read the promo code from the client
		buf = make([]byte, 256)
		_, err = io.ReadFull(conn, buf)
		if err != nil {
			return
		}

		promoCodeLen := int(buf[0])
		promoCodeBuf := make([]byte, promoCodeLen)
		_, err = io.ReadFull(conn, promoCodeBuf)
		if err != nil {
			return
		}

		// Check if the provided promo code is valid
		if string(promoCodeBuf) == config.promo_code {
			// Promo code is valid, allow access to the paid network
			// Process the SOCKS5 request and relay data as before
			// ...
		} else {
			// Promo code is invalid, return an error
			_, err = conn.Write([]byte{5, 1, 0, 8, 0, 0, 0, 0, 0, 0}) // Send a SOCKS5 response with "connection not allowed by ruleset" error
			if err != nil {
				return
			}
			return
		}
	} else {
		// Free connection
		// Process the SOCKS5 request and relay data as before
		// ...
	}
}
func checkSubscription(userID int) bool {

	return true // For TON FunC Contact
}
func sendConfigRequest(subscription string) {
	var url string
	switch subscription {
	case "free":
		url = "https://example.com/free-config"
	case "paid":
		url = "https://example.com/paid-config"
	default:
		fmt.Println("Invalid subscription type")
		return
	}

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error sending config request:", err)
		return
	}
	defer resp.Body.Close()

	// Process the response if needed
	// ...

	fmt.Println("Config request sent successfully")
}

func anon() {
	sendConfigRequest("free") // Send request for free config
	sendConfigRequest("paid") // Send request for paid config
}
