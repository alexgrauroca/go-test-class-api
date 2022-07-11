package repository

import (
	"context"

	"github.com/alexgrauroca/go-test-class-api/models"
)

func InsertBooking(ctx context.Context, c *models.Booking) error {
	return implementation.InsertBooking(ctx, c)
}

func TruncateBooking(ctx context.Context) error {
	return implementation.TruncateBooking(ctx)
}
