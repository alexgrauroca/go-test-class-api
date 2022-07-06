package tests_test

import (
	"context"
	mockdb "go-test-class-api/database/mock"
	"go-test-class-api/models"
	"go-test-class-api/repository"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var c models.Class
var repo repository.Repository

func TestMain(m *testing.M) {
	// Before tests
	c = models.Class{
		Name:      "Testing",
		StartDate: time.Now(),
		EndDate:   time.Now().AddDate(0, 0, 1),
		Capacity:  10,
	}

	repo, _ = mockdb.NewMockRepository()

	// Run tests
	exitVal := m.Run()

	// Exit tests
	os.Exit(exitVal)
}

func TestCreateClass(t *testing.T) {
	assert.NotNil(t, c)
	assert.NotNil(t, c.NewId())
	assert.NotEmpty(t, c.Id)
	assert.Nil(t, repository.InsertClass(context.Background(), &c))
}

func TestReadClass(t *testing.T) {

}

func TestUpdateClass(t *testing.T) {

}

func TestDeleteClass(t *testing.T) {

}
