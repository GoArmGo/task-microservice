package main

import (
	"log"
	"task-microservice/internal/db"
	"task-microservice/pkg/config"
)

func main() {

	con := config.LoadConfig()
	db := db.NewPostgres(con)

	if err := db.Ping(); err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}

	log.Println("✅ Successfully connected to the database!")
}
