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

type 

func handler() {
	proxyURL, err := url.Parse(
		fmt.Sprintf("%s%s%d", TRANSFER_PROTOCOL, MAIN_ROUTE, MAIN_PORT),
	)

	if err != nil {
		raise(ErrorsGlobal.network.BrokenResponse)
	}

	client := &http.Client{Transport: transport}

	req, err := http.NewRequest("GET", "https://example.com", nil)
	if err != nil {
		raise(ErrorsGlobal.network.BrokenRequest)
		return
	}
	resp, err := client.Do(req)
	if err != nil {
		raise(ErrorsGlobal.network.BrokenRequest)
		return
	}
	defer resp.Body.Close()

	body, err := json.Marshal(resp.Body)
	if err != nil {
		raise(ErrorsGlobal.network.JsonConversionError)
		return
	}
}

func loadNetInit() {
	config := &tls.Config{
		RootCAs: certPool,
	}

	transport := &http.Transport{
		TLSClientConfig: config,
	}

	client := &http.Client{
		Transport: transport,
	}
}

func serviceAvailable(d NetworkDaemon) bool {

	go func() (avail bool, code int) {
		for {
			resp, _ := req(d, MAIN_URL_PATH, nil, nil)

			time.Sleep(d.updateTime * time.Second)
			fmt.Print("Network daemon is running...")

			if resp.StatusCode != 200 {
				avail = false
			}

			return avail, resp.StatusCode
		}
	}()
}


func req(
	daemon NetworkDaemon, url string, headers, payload map[string]string,
) (*http.Response, error) {

	daemonClient := &http.Client{}
	request, req_err := http.NewRequest(daemon.method, url, nil)

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

	response, res_err := daemonClient.Do(request)

	if res_err != nil {
		raise(Err.net.BrokenResponse)
	}

	return response, nil
}

