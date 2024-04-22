package main

import (
	"fmt"
	"os"
)

var config_path string = os.Getenv("")

func main() {
	err := addUser("новый_пользователь", "пароль_нового_пользователя")
	if err != nil {
		panic(fmt.Sprintf("Ошибка при добавлении пользователя:", err))
	} else {
		fmt.Println("Пользователь успешно добавлен")
	}
	sendConfigRequest("free") // Send request for free config
	sendConfigRequest("paid") // Send request for paid config

	err = removeUser("пользователь_для_удаления")
	if err != nil {
		panic(fmt.Sprintf("Ошибка при удалении пользователя:", err))
	} else {
		fmt.Println("Пользователь успешно удален")
	}
}
