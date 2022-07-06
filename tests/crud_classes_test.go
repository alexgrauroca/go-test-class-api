package tests_test

import (
	"go-test-class-api/models"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var c models.Class

func TestMain(m *testing.M) {
	// Before tests
	c = models.Class{
		Name:      "Testing",
		StartDate: time.Now(),
		EndDate:   time.Now().AddDate(0, 0, 1),
		Capacity:  10,
	}

	// Run tests
	exitVal := m.Run()

	// Exit tests
	os.Exit(exitVal)
}

func TestCreateClass(t *testing.T) {
	assert.NotNil(t, c)
	assert.Nil(t, c.InsertClass())
	assert.NotEmpty(t, c.Id)
}

func TestReadClass(t *testing.T) {

}

func TestUpdateClass(t *testing.T) {

}

func TestDeleteClass(t *testing.T) {

}
