package postgresdb

import (
	"context"
	"errors"

	"github.com/alexgrauroca/go-test-class-api/models"
)

func (repo *PostgresRepository) InsertBooking(ctx context.Context, booking *models.Booking) error {
	var filters = models.ClassFilter{
		Date: booking.Date,
	}

	classes, err := repo.GetClasses(ctx, &filters)

	if err != nil {
		return err
	}

	if len(classes) == 0 {
		return errors.New("No classes are available for the specified date")
	}

	_, err = repo.db.ExecContext(ctx, "INSERT INTO bookings (id, name, date) VALUES ($1, $2, $3)", booking.Id, booking.Name, booking.Date)

	return err
}

func (repo *PostgresRepository) TruncateBooking(ctx context.Context) error {
	_, err := repo.db.ExecContext(ctx, "TRUNCATE TABLE bookings")
	return err
}
