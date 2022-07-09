package crud_classes_test

import (
	"bytes"
	"encoding/json"
	"go-test-class-api/handlers"
	"go-test-class-api/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type ClassesResponse struct {
	Classes []*models.Class `json:"classes"`
}

func TestApiV1InsertClass(t *testing.T) {
	// Creating new test request
	req, err := http.NewRequest("POST", "/api/v1/classes", bytes.NewBuffer(classJson))
	assert.Nil(t, err)

	if err != nil {
		return
	}

	// Execute test request
	rec := httptest.NewRecorder()
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
func TestApiV1GetClasses(t *testing.T) {
	// Creating new test request
	req, err := http.NewRequest("GET", "/api/v1/classes", nil)
	assert.Nil(t, err)

	if err != nil {
		return
	}

	// Execute test request
	// The first request will return a not found status.
	rec := httptest.NewRecorder()
	handlers.GetClasses(rec, req)

	res := rec.Result()

	// Assertions
	assert.Equal(t, http.StatusNotFound, res.StatusCode)

	// The second one will return an ok status.
	rec = httptest.NewRecorder()
	handlers.GetClasses(rec, req)

	res = rec.Result()
	defer res.Body.Close()

	// Assertions
	assert.Equal(t, http.StatusOK, res.StatusCode)

	// Read body
	var classesResponse = ClassesResponse{}
	err = json.NewDecoder(res.Body).Decode(&classesResponse)
	assert.Nil(t, err)

	if err != nil {
		return
	}

	assert.Equal(t, 1, len(classesResponse.Classes))
}
