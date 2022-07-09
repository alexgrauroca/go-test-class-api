package postgresdb

import (
	"context"
	"errors"
	"go-test-class-api/models"
)

func (repo *PostgresRepository) InsertBooking(ctx context.Context, booking *models.Booking) error {
	/*_, err := repo.db.ExecContext(ctx, "INSERT INTO bookings (id, name, date) VALUES ($1, $2, $3)", booking.Id, booking.Name, booking.Date)
	return err*/

	return errors.New("Not implemented")
}

func (repo *PostgresRepository) TruncateBooking(ctx context.Context) error {
	/*_, err := repo.db.ExecContext(ctx, "TRUNCATE TABLE bookings")
	return err*/

	return errors.New("Not implemented")
}
