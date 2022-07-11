package crud_bookings_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alexgrauroca/go-test-class-api/handlers"
	"github.com/alexgrauroca/go-test-class-api/models"

	"github.com/stretchr/testify/assert"
)

type BookingsResponse struct {
	Bookings []*models.Booking `json:"bookings"`
}

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

func TestApiV1GetBookings(t *testing.T) {
	// Creating new test request
	req, err := http.NewRequest("GET", "/api/v1/bookings", nil)
	assert.Nil(t, err)

	if err != nil {
		return
	}

	// Execute test request
	// The first request will return a not found status.
	rec := httptest.NewRecorder()
	handlers.GetBookings(rec, req)

	res := rec.Result()

	// Assertions
	assert.Equal(t, http.StatusNotFound, res.StatusCode)

	// The second one will return an ok status.
	rec = httptest.NewRecorder()
	handlers.GetBookings(rec, req)

	res = rec.Result()
	defer res.Body.Close()

	// Assertions
	assert.Equal(t, http.StatusOK, res.StatusCode)

	// Read body
	var bookingsResponse = BookingsResponse{}
	err = json.NewDecoder(res.Body).Decode(&bookingsResponse)
	assert.Nil(t, err)

	if err != nil {
		return
	}

	assert.Equal(t, 1, len(bookingsResponse.Bookings))
}
