package main

import (
	main2 "awesomeProject"
	"time"
)

type NetworkDaemon struct {
	endpoint   string
	idleTime   time.Duration
	updateTime time.Duration
	threads    uint8
	maxRetries uint8
	method     string
	state      string
}

type daemonStates interface {
	run()
	drop()
	update()
}

func (d NetworkDaemon) run() {
	for {

		res, er := req(d, MAIN_URL_PATH, nil, nil)

		if er != nil {
			d.drop()
			break
		}

		time.Sleep(d.updateTime * time.Second)
		return
	}
}
func (d NetworkDaemon) drop()   {}
func (d NetworkDaemon) update() {}

func main() {}
