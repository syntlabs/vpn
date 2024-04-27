package main

import (
	"context"
	"crypto/rand"
	"encoding/base32"
	"fmt"
	"github.com/cameo-engineering/tonconnect"
	"golang.org/x/exp/maps"
	"log"
	"os"
	"time"
)

var config_path string = os.Getenv("")

func main() {
	s, err := tonconnect.NewSession()
	if err != nil {
		log.Fatal(err)
	}
	data := make([]byte, 32)
	_, err = rand.Read(data)
	if err != nil {
		log.Fatal(err)
	}

	connreq, err := tonconnect.NewConnectRequest(
		"https://raw.githubusercontent.com/syntlabs/vpn/master/tonconnect-manifest.json",
		tonconnect.WithProofRequest(base32.StdEncoding.EncodeToString(data)),
	)
	if err != nil {
		log.Fatal(err)
	}

	deeplink, err := s.GenerateDeeplink(*connreq, tonconnect.WithBackReturnStrategy())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Deeplink: %s\n\n", deeplink)

	wrapped := tonconnect.WrapDeeplink(deeplink)
	fmt.Printf("Wrapped deeplink: %s\n\n", wrapped)

	for _, wallet := range tonconnect.Wallets {
		link, err := s.GenerateUniversalLink(wallet, *connreq)
		fmt.Printf("%s: %s\n\n", wallet.Name, link)
		if err != nil {
			log.Fatal(err)
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	res, err := s.Connect(ctx, (maps.Values(tonconnect.Wallets))...)
	if err != nil {
		log.Fatal(err)
	}

	var addr string
	network := "mainnet"
	for _, item := range res.Items {
		if item.Name == "ton_addr" {
			addr = item.Address
			if item.Network == -3 {
				network = "testnet"
			}
		}
	}
	fmt.Printf(
		"%s %s for %s is connected to %s with %s address\n\n",
		res.Device.AppName,
		res.Device.AppVersion,
		res.Device.Platform,
		network,
		addr,
	)

	sendConfigRequest(addr) // Send request for config
}
