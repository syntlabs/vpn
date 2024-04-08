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
}}

type VpnConfig struct {
	ip,
	country []string
}

func checkInternet() {

	response, err = req("https://www.google.com", nil, nil)

	if response.StatusCode != 200 {
		raise(Err.net.NoConnection)
		someViewNoInternet()
	}
}

func req(url string, headers, payload map[string]interface{}) (*http.Response, error) {

	request, req_err := http.NewRequest("GET", url, headers, payload)

	if req_err != nil {
		raise(Err.net.BrokenRequest)
	}

	if headers != nil {
		for key, val := range headers {
			request.Header.Add(key, val)
		}
	}
	if payload != nil {
		for key, val := range payload {
			request.PostForm.Add(key, val)
		}
	}

	response, res_err := mainClient.Do(request)

	if res_err != nil {
		raise(Err.net.BrokenResponse)
	}

	return response, nil
}

