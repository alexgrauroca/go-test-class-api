package postgres_test

import (
	"context"
	"go-test-class-api/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepositoryInsertBooking(t *testing.T) {
	// Create class for bookings tests
	newClass := c
	newClass.StartDate = booking.Date.AddDate(0, 0, -2)
	newClass.EndDate = booking.Date.AddDate(0, 0, 2)
	newClass.NewId()

	assert.Nil(t, repository.InsertClass(context.Background(), &newClass))

	assert.NotNil(t, booking)
	assert.Nil(t, booking.NewId())
	assert.NotEmpty(t, booking.Id)
	assert.Nil(t, repository.InsertBooking(context.Background(), &booking))

	// Testing diferent dates, to validate every possible combination
	errorBooking := booking
	errorBooking.Id = ""
	errorBooking.Date = newClass.StartDate.AddDate(0, 0, -2)

	assert.NotNil(t, errorBooking)
	assert.Nil(t, errorBooking.NewId())
	assert.NotEmpty(t, errorBooking.Id)
	assert.NotNil(t, repository.InsertBooking(context.Background(), &errorBooking))
}
