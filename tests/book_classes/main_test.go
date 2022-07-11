package crud_bookings_test

import (
	"encoding/json"
	"log"
	"os"
	"testing"

	mockdb "go-test-class-api/database/mock"
	"go-test-class-api/models"
	"go-test-class-api/repository"
)

var booking models.Booking
var repo repository.Repository
var bookingJson []byte

func TestMain(m *testing.M) {
	// Before tests
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

	// bookingJson is used to send correctly the api body, with the os.File it doesn't work
	b, err := os.ReadFile("./data/booking.json")

	if err != nil {
		log.Fatal("Error reading 2 booking.json: ", err)
	}

	bookingJson = []byte(string(b))

	// Init mock db repository
	repo, _ = mockdb.NewMockRepository()
	repository.SetRepository(repo)

	// Run tests
	exitVal := m.Run()

	// Exit tests
	os.Exit(exitVal)
}
