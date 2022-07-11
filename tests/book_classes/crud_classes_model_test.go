package crud_bookings_test

import (
	"context"
	"testing"

	"github.com/alexgrauroca/go-test-class-api/repository"

	"github.com/stretchr/testify/assert"
)

func TestCreateBooking(t *testing.T) {
	assert.NotNil(t, booking)
	assert.Nil(t, booking.NewId())
	assert.NotEmpty(t, booking.Id)
	assert.Nil(t, repository.InsertBooking(context.Background(), &booking))
}
