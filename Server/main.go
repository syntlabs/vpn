package main

import (
	"os"
)

var config_path string = os.Getenv("")

const maxThreads = 10

var goPoolin = make(chan struct{}, maxThreads)

func main() {

	sendConfigRequest("free") // Send request for free config
	sendConfigRequest("paid") // Send request for paid config
}
