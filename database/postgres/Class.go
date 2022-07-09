package postgresdb

import (
	"context"
	"go-test-class-api/models"
	"log"
)

func (repo *PostgresRepository) GetClasses(ctx context.Context) ([]*models.Class, error) {
	rows, err := repo.db.QueryContext(ctx, "SELECT id, name, start_date, end_date, capacity FROM classes ORDER BY name ASC")

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

	var classes []*models.Class

	for rows.Next() {
		var c = models.Class{}

		if err := rows.Scan(&c.Id, &c.Name, &c.StartDate, &c.EndDate, &c.Capacity); err != nil {
			return nil, err
		}

		classes = append(classes, &c)
	}

	return classes, nil
}

func (repo *PostgresRepository) InsertClass(ctx context.Context, c *models.Class) error {
	_, err := repo.db.ExecContext(ctx, "INSERT INTO classes (id, name, start_date, end_date, capacity) VALUES ($1, $2, $3, $4, $5)", c.Id, c.Name, c.StartDate, c.EndDate, c.Capacity)
	return err
}

func (repo *PostgresRepository) TruncateClass(ctx context.Context) error {
	_, err := repo.db.ExecContext(ctx, "TRUNCATE TABLE classes")
	return err
}
