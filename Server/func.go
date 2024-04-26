package main

import (
	"fmt"
	"net/http"
)


func checkSubscription(addr string) string {

	return "Free" // TBA
}
var subscription string
func sendConfigRequest(addr string) {
	subscription = checkSubscription(addr)
	var url string
	switch subscription {
	case "Free":
		url = "https://example.com/free-config"
	case "Paid":
		url = "https://example.com/paid-config"
	default:
		fmt.Println("Invalid subscription type")
		return
	}

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error sending config request:", err)
		return
	}
	defer resp.Body.Close()

	// Process the response if needed
	// ...

	fmt.Println(subscription + " config request sent successfully")
}


