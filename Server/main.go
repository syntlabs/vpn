package main

import (
	"fmt"
	"os"
)

var config_path string = os.Getenv("")

func main() {

	sendConfigRequest("free") // Send request for free config
	sendConfigRequest("paid") // Send request for paid config
	
	// Пример использования функций
	err := addUser("новый_пользователь", "пароль_нового_пользователя")
	if err != nil {
		fmt.Println("Ошибка при добавлении пользователя:", err)
	} else {
		fmt.Println("Пользователь успешно добавлен")
	}

	err = removeUser("пользователь_для_удаления")
	if err != nil {
		fmt.Println("Ошибка при удалении пользователя:", err)
	} else {
		fmt.Println("Пользователь успешно удален")
	}
}
