package postgres_test

import (
	"context"
	"go-test-class-api/models"
	"go-test-class-api/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepositoryInsertClass(t *testing.T) {
	assert.NotNil(t, c)
	assert.Nil(t, c.NewId())
	assert.NotEmpty(t, c.Id)
	assert.Nil(t, repository.InsertClass(context.Background(), &c))
}

func TestRepositoryGetClasses(t *testing.T) {
	var expectedClasses []*models.Class
	classes, err := repository.GetClasses(context.Background())

	assert.Nil(t, err)
	assert.IsType(t, expectedClasses, classes)
}
