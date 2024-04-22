package main

import (
	"fmt"
	"net/http"
	"os/exec"
)

const (
	freeConfigURL = "https://example.com/free-config"
	paidConfigURL = "https://example.com/paid-config"
)

func checkSubscription(userID int) bool {

	return true // For TON FunC Contact
}

func sendConfigRequest(subscription string) {
	var url string
	switch subscription {
	case "free":
		url = freeConfigURL
	case "paid":
		url = paidConfigURL
	default:
		panic("Invalid subscription type")
		return
	}

	resp, err := http.Get(url)
	if err != nil {
		panic(fmt.Sprintf("Error sending config request:", err))
		return
	}
	defer resp.Body.Close()

	// Process the response if needed
	// ...

	fmt.Println("Config request sent successfully")
}

// Функция для добавления пользователя
func addUser(username, password string) error {
	cmd := exec.Command("ssserver", "-u", "-c", "/путь/к/конфигурации/shadowsocks.json", "-p", "порт", "-k", "пароль", "-m", "метод", "--adduser", username, password)
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

// Функция для удаления пользователя
func removeUser(username string) error {
	cmd := exec.Command("ssserver", "-u", "-c", "/путь/к/конфигурации/shadowsocks.json", "-p", "порт", "-k", "пароль", "-m", "метод", "--removeuser", username)
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
