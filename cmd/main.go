package main

import (
	"fmt"
	"github.com/TimmyTurner98/project/server"
)

func main() {
	// Создание экземпляра сервера
	server := &server.Server{}

	// Запуск сервера на порту 8080
	err := server.Run("8080")
	if err != nil {
		fmt.Printf("Ошибка при запуске сервера: %v\n", err)
	}
}
