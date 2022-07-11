package postgresdb

import (
	"context"
	"errors"
	"log"

	"github.com/alexgrauroca/go-test-class-api/models"
)

func (repo *PostgresRepository) GetBookings(ctx context.Context) ([]*models.Booking, error) {
	rows, err := repo.db.QueryContext(ctx, "SELECT id, name, date FROM bookings ORDER BY date ASC, name ASC")

	if err != nil {
		return nil, err
	}

	// Closing returned dataset
	defer func() {
		err = rows.Close()

		if err != nil {
			log.Fatal(err)
		}
	}()

	var bookings []*models.Booking

	for rows.Next() {
		var booking = models.Booking{}

		if err := rows.Scan(&booking.Id, &booking.Name, &booking.Date); err != nil {
			return nil, err
		}

		bookings = append(bookings, &booking)
	}

	return bookings, nil
}

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
