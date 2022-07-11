package postgres_test

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"testing"

	postgresdb "go-test-class-api/database/postgres"
	"go-test-class-api/models"
	"go-test-class-api/repository"

	"github.com/joho/godotenv"
)

var booking models.Booking
var c models.Class
var repo repository.Repository

func TestMain(m *testing.M) {
	// Before tests
	// Read class.json
	classFile, err := os.Open("./data/class.json")
	defer classFile.Close()

	if err != nil {
		log.Fatal("Error reading class.json: ", err)
	}

	// Decode class.json to models.Class
	if err := json.NewDecoder(classFile).Decode(&c); err != nil {
		log.Fatal("Error decoding class.json: ", err)
	}

	// Read booking.json
	bookingFile, err := os.Open("./data/booking.json")
	defer bookingFile.Close()

	if err != nil {
		log.Fatal("Error reading booking.json: ", err)
	}

	// Decode booking.json to models.Class
	if err := json.NewDecoder(bookingFile).Decode(&booking); err != nil {
		log.Fatal("Error decoding booking.json: ", err)
	}

	if err := godotenv.Load("./../../.env"); err != nil {
		log.Fatal("Error loading .env file: ", err)
	}

	// Init mock db repository
	repo, _ = postgresdb.NewPostgresRepository(os.Getenv("DATABASE_TEST_URL"))
	repository.SetRepository(repo)

	// Run tests
	exitVal := m.Run()

	cleanTables()

	// Exit tests
	os.Exit(exitVal)
}

func cleanTables() {
	repository.TruncateBooking(context.Background())
	repository.TruncateClass(context.Background())
}
