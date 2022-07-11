package crud_classes_test

import (
	"encoding/json"
	"log"
	"os"
	"testing"

	mockdb "github.com/alexgrauroca/go-test-class-api/database/mock"
	"github.com/alexgrauroca/go-test-class-api/models"
	"github.com/alexgrauroca/go-test-class-api/repository"
)

var c models.Class
var repo repository.Repository
var classJson []byte

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

	// classJson is used to send correctly the api body, with the os.File it doesn't work
	b, err := os.ReadFile("./data/class.json")

	if err != nil {
		log.Fatal("Error reading 2 class.json: ", err)
	}

	classJson = []byte(string(b))

	// Init mock db repository
	repo, _ = mockdb.NewMockRepository()
	repository.SetRepository(repo)

	// Run tests
	exitVal := m.Run()

	// Exit tests
	os.Exit(exitVal)
}
