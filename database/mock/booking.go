package mockdb

import (
	"context"

	"github.com/alexgrauroca/go-test-class-api/models"
)

func (repo *MockRepository) InsertBooking(ctx context.Context, c *models.Booking) error {
	return nil
}

func (repo *MockRepository) TruncateBooking(ctx context.Context) error {
	return nil
}
