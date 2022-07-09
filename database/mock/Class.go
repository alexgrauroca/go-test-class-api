package mockdb

import (
	"context"
	"go-test-class-api/models"
)

func (repo *MockRepository) GetClasses(ctx context.Context) ([]*models.Class, error) {
	var emptyClasses []*models.Class
	classes := repo.Classes

	// At the first request, mock will return an empty response (404 - not found for the API).
	// At the second one, it will return a response with 1 result.
	// This is to make the mock testing easier
	if len(repo.Classes) > 0 {
		repo.Classes = emptyClasses
	} else {
		var c = models.Class{}
		repo.Classes = append(emptyClasses, &c)
	}

	return classes, nil
}

func (repo *MockRepository) InsertClass(ctx context.Context, c *models.Class) error {
	return nil
}

func (repo *MockRepository) TruncateClass(ctx context.Context) error {
	return nil
}
