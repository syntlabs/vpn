package main

import (
	"time"
)

type NetworkDaemon struct {
	endpoint   string
	idleTime   time.Duration
	updateTime time.Duration
	threads    uint8
	maxRetries uint8
	state      string
}

type daemonStates interface {
	run()
	infinityPolling()
}

func keepTunnel(_type string) {
	// Watch for free or paid period and keep related settings
}

func keepSecureConnection() {
	// Watch for secure connection
}

func (d NetworkDaemon) run() {

	d.infinityPolling()
}

func (d NetworkDaemon) infinityPolling() {

	go func() {
		
		for {

			checkInternet()
			keepSecureConnection()

			if mainClient.connected {

				if mainClient.subscription > 0 {
					keepTunnel("Paid")
				} else {
					keepTunel("Free")
				}
			}

			time.sleep(d.updateTime * time.Second)
		}
	}()
}

func main() {}
