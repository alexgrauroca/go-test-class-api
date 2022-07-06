package repository

import (
	"context"
	"go-test-class-api/models"
)

func InsertClass(ctx context.Context, c *models.Class) error {
	return implementation.InsertClass(ctx, c)
}

func TruncateClass(ctx context.Context) error {
	return implementation.TruncateClass(ctx)
}
