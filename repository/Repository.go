package repository

import (
	"context"
	"go-test-class-api/models"
)

type Repository interface {
	Close() error

	InsertClass(ctx context.Context, c *models.Class) error
}

var implementation Repository

func SetRepository(repository Repository) {
	implementation = repository
}

func Close() error {
	return implementation.Close()
}
