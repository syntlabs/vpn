package Utils

import (
	"crypto/tls"
	"net"
	"net/http"
	"runtime"
	"time"
	"utils.go"
)

type UserClient struct {
	name,
	lang,
	vpnLocation,
	theme,
	fingerprint,
	passworder,
	connected,
	system string
	subscription uint16 // in minutes
	balance            float32
}

type Usermethods interface {
	update(newSpecs map[string]interface{}) UserClient
	remove()
}

func (user UserClient) update(newSpecs  map[string]interface{}) UserClient {}

func createMainClient() {

	currentLang, langerr := getSystemLanguage()

	if langerr != nil {
		raise(Err.val.WrongLanguage)
	}

	var mneumonic := generateMneumonic(12)

	cert, certPool = certnpool()

	startVpnConfig = VpnConfig{{""}, {""}}

	config := &tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      certPool,
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

	mainClient := UserClient{
		name:             "Bob",
		lang:             currentLang,
		coins:            nil,
		vpnLocation:      "std",
		theme:            "std",
		fingerprint:      hashMnem(mneumonic),
		passworder:       "no",
		connected:        "",
		subscriptionDays: 0,
		system:           runtime.GOOS,
		vpnconfig: startVpnConfig,
		dialer: dialer,
		transport: transport
	}
}
