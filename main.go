package main

import (
	"fmt"
	"os"
)

func main() {
	host := os.Getenv("TARGET_HOST")
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
		fmt.Println("⚠️ Переменная APP_PORT не найдена. Используем значение по умолчанию.")
	}
	if host == "" {
		host = "localhost"
		fmt.Println("⚠️ Переменная TARGET_HOST не найдена. Используем значение по умолчанию.")
	}
	fmt.Printf("Host: %s, Port: %s\n", host, port)
}
