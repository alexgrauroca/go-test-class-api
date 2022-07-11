package main

import (
	"context"
	"flag"
	"log"
	"os"

	"github.com/alexgrauroca/go-test-class-api/server"

	"github.com/alexgrauroca/go-test-class-api/router"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file: ", err)
	}

	appType := flag.String("run-mode", "production", "specify app running type: production or test")
	flag.Parse()

	databaseUrl := os.Getenv("DATABASE_PRODUCTION_URL")
	port := os.Getenv("PRODUCTION_PORT")

	if *appType == "test" {
		log.Println("Running on test mode")
		databaseUrl = os.Getenv("DATABASE_TEST_URL")
		port = os.Getenv("TEST_PORT")
	}

	s, err := server.NewServer(context.Background(), &server.Config{
		Port:        port,
		DatabaseUrl: databaseUrl,
	})

	if err != nil {
		log.Fatal("Error starting new server: ", err)
	}

	s.Start(router.BindRoutes)
}
