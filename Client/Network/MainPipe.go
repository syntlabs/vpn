package main

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"os"
	"time"
)

const MAIN_PORT uint16 = 9090
const MAIN_ROUTE string = ""
const SECURE_PROTOCOL string = "socks5"
const TRANSFER_PROTOCOL string = "SSH"

var vpn_ips = map[string][]string{
	"USA":     {"1.1.1.1", "0.0.0.00", "2.2.2.2", "3.3.3.3"},
	"Italy":   {"1.1.1.1", "0.0.0.00", "2.2.2.2", "3.3.3.3"},
	"Germany": {"1.1.1.1", "0.0.0.00", "2.2.2.2", "3.3.3.3"},
}

type Pipe struct {
	port uint16
	route,
	sec_prot,
	tr_prot string
	bandwidth  float32
	speed      float32
	multiLevel bool
}

type NetConfig struct {
	pipe     Pipe
	login    string
	password string
}

type pumping interface {
	pump(any) any
}

func (p Pipe) pump() {
	proxyURL, err := url.Parse(
		fmt.Sprintf("%s%s%d", TRANSFER_PROTOCOL, MAIN_ROUTE, MAIN_PORT),
	)
	if err != nil {

		raise(ErrorsGlobal.network.brokenResponse)
		return
	}

	dialer := &net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
	}
	transport := &http.Transport{
		Proxy:                 http.ProxyURL(proxyURL),
		DialTLSContext:        dialer.Dial,
		TLSClientConfig:       tlsConfig,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

	client := &http.Client{Transport: transport}

	req, err := http.NewRequest("GET", "https://example.com", nil)
	if err != nil {
		raise(ErrorsGlobal.network.brokenRequest)
		return
	}
	resp, err := client.Do(req)
	if err != nil {
		raise(ErrorsGlobal.network.brokenRequest)
		return
	}
	defer resp.Body.Close()

	body, err := json.Marshal(resp.Body)
	if err != nil {
		raise(ErrorsGlobal.network.jsonConversionError)
		return
	}

	fmt.Println(string(body))
}

func main() {
	mainPipe := Pipe{
		port:     MAIN_PORT,
		route:    MAIN_ROUTE,
		sec_prot: SECURE_PROTOCOL,
		tr_prot:  TRANSFER_PROTOCOL,
	}

	netconf := NetConfig{
		pipe:     mainPipe,
		login:    os.Getenv("login"),
		password: os.Getenv("password"),
	}
}
