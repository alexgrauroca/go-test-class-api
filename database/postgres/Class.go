package postgresdb

import (
	"context"
	"go-test-class-api/models"
)

func (repo *PostgresRepository) InsertClass(ctx context.Context, c *models.Class) error {
	_, err := repo.db.ExecContext(ctx, "INSERT INTO classes (id, name, start_date, end_date, capacity) VALUES ($1, $2, $3, $4, $5)", c.Id, c.Name, c.StartDate, c.EndDate, c.Capacity)
	return err
}

func (repo *PostgresRepository) TruncateClass(ctx context.Context) error {
	_, err := repo.db.ExecContext(ctx, "TRUNCATE TABLE classes")
	return err
}
