package mockdb

import (
	"context"

	"github.com/alexgrauroca/go-test-class-api/models"
)

func (repo *MockRepository) GetBookings(ctx context.Context) ([]*models.Booking, error) {
	var emptyBookings []*models.Booking
	bookings := repo.Bookings

	// At the first request, mock will return an empty response (404 - not found for the API).
	// At the second one, it will return a response with 1 result.
	// This is to make the mock testing easier
	if len(repo.Bookings) > 0 {
		repo.Bookings = emptyBookings
	} else {
		var booking = models.Booking{}
		repo.Bookings = append(emptyBookings, &booking)
	}

	return bookings, nil
}

func (repo *MockRepository) InsertBooking(ctx context.Context, c *models.Booking) error {
	return nil
}

func (repo *MockRepository) TruncateBooking(ctx context.Context) error {
	return nil
}
