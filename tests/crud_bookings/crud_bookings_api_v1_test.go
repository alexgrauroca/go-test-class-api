package crud_bookings_test

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

func TestApiV1InsertBooking(t *testing.T) {
	// Creating new test request
	req, err := http.NewRequest("POST", "/api/v1/bookings", bytes.NewBuffer(bookingJson))
	assert.Nil(t, err)

	if err != nil {
		return
	}

	// Execute test request
	rec := httptest.NewRecorder()
	handlers.InsertBooking(rec, req)

	res := rec.Result()
	defer res.Body.Close()

	// Assertions
	assert.Equal(t, http.StatusCreated, res.StatusCode)

	// Read body
	var newBooking = models.Booking{}
	err = json.NewDecoder(res.Body).Decode(&newBooking)
	assert.Nil(t, err)

	if err != nil {
		return
	}

	assert.NotEmpty(t, newBooking.Id)
	assert.Equal(t, booking.Name, newBooking.Name)
	assert.Equal(t, booking.Date, newBooking.Date)
}
