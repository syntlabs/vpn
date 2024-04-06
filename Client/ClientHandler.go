package main

import (
	"runtime"
	"time"
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
	subscriptionDays uint16
	coins            map[string]float32
	vpnconfig VpnConfig
	dialer Dialer
	transport Transport
}

type Usermethods interface {
	update(newSpecs map[string]interface) UserClient
	remove()
}

func (user UserClient) update(newSpecs  map[string]interface) UserClient {

	for field, update := range newSpecs {

	}
}

func createMainClient() {

	currentLang, langerr := getSystemLanguage()

	if langerr != nil {
		raise(Err.val.WrongLanguage)
	}

	startVpnConfig = VpnConfig{{""}, {""}}

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
		fingerprint:      hashUnique(time.Now().UnixMilli(), salt(SALT_SIZE)),
		passworder:       "no",
		connected:        "",
		subscriptionDays: 0,
		system:           runtime.GOOS,
		vpnconfig: startVpnConfig,
		dialer: dialer,
		transport: transport
	}
}
