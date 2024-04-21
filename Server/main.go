package main

import (
	"os"
)

var config_path string = os.Getenv("")

func main() {

	sendConfigRequest("free") // Send request for free config
	sendConfigRequest("paid") // Send request for paid config
}
