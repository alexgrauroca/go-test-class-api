package main

import (
	"context"
	"go-test-class-api/server"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file: ", err)
	}

	s, err := server.NewServer(context.Background(), &server.Config{
		Port:        os.Getenv("PORT"),
		DatabaseUrl: os.Getenv("DATABASE_URL"),
	})

	if err != nil {
		log.Fatal("Error starting new server: ", err)
	}

	s.Start(router.BindRoutes)
}
