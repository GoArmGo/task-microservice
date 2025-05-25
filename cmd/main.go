package main

import (
	"log"
	"task-microservice/internal/di"
)

func main() {
	// сбор зависимостей через di контейнер
	container, err := di.NewContainer()
	if err != nil {
		log.Fatalf("failed to initialize container: %v", err)
	}

	// пуск сервера
	log.Println("🚀 Server is running on http://localhost:8080")
	if err := container.Router.Run(":8080"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
