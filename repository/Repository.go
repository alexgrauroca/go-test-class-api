package repository

import (
	"context"

	"github.com/alexgrauroca/go-test-class-api/models"
)

type Repository interface {
	Close() error

	// Class functions
	GetClasses(ctx context.Context, filters *models.ClassFilter) ([]*models.Class, error)
	InsertClass(ctx context.Context, c *models.Class) error
	TruncateClass(ctx context.Context) error

	// Booking functions
	GetBookings(ctx context.Context) ([]*models.Booking, error)
	InsertBooking(ctx context.Context, c *models.Booking) error
	TruncateBooking(ctx context.Context) error
}

var implementation Repository

func SetRepository(repository Repository) {
	implementation = repository
}

func Close() error {
	return implementation.Close()
}
