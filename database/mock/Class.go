package mockdb

import (
	"context"
	"go-test-class-api/models"
)

func (repo *MockRepository) InsertClass(ctx context.Context, c *models.Class) error {
	return nil
}
