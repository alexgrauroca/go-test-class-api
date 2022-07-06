package crud_classes_test

import (
	"bytes"
	"encoding/json"
	"go-test-class-api/handlers"
	"go-test-class-api/models"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApiV1InsertClass(t *testing.T) {
	// Creating new test request
	req, err := http.NewRequest("POST", "/api/v1/classes", bytes.NewBuffer(classJson))
	assert.Nil(t, err)

	if err != nil {
		return
	}

	// Execute test request
	rec := httptest.NewRecorder()
	log.Println(req.Body)
	log.Println(req.ContentLength)
	handlers.InsertClass(rec, req)

	res := rec.Result()
	defer res.Body.Close()

	// Assertions
	assert.Equal(t, http.StatusCreated, res.StatusCode)

	// Read body
	var newC = models.Class{}
	err = json.NewDecoder(res.Body).Decode(&newC)
	assert.Nil(t, err)

	if err != nil {
		return
	}

	assert.NotEmpty(t, newC.Id)
	assert.Equal(t, c.Name, newC.Name)
	assert.Equal(t, c.StartDate, newC.StartDate)
	assert.Equal(t, c.EndDate, newC.EndDate)
	assert.Equal(t, c.Capacity, newC.Capacity)
}
