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
}

func createMainClient() {

	currentLang, langerr := getSystemLanguage()

	if langerr != nil {
		raise(Antares_vpn.ErrorsGlobal.value.wrongLanguage)
	}

	mainClient := UserClient{
		name:             "Bob",
		lang:             currentLang,
		coins:            nil,
		vpnLocation:      "std",
		theme:            "std",
		fingerprint:      hashUnique(time.Now().UnixMilli(), salt(Antares_vpn.SALT_SIZE)),
		passworder:       "no",
		connected:        "",
		subscriptionDays: 0,
		system:           runtime.GOOS,
	}
}
