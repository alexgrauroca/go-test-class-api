package postgres_test

import (
	"context"
	"go-test-class-api/models"
	"go-test-class-api/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepositoryInsertClass(t *testing.T) {
	errorClass := c

	assert.NotNil(t, c)
	assert.Nil(t, c.NewId())
	assert.NotEmpty(t, c.Id)
	assert.Nil(t, repository.InsertClass(context.Background(), &c))

	// This class can't be inserted, because there's an existing class with similar date range
	assert.NotNil(t, errorClass)
	assert.Nil(t, errorClass.NewId())
	assert.NotEmpty(t, errorClass.Id)
	assert.NotNil(t, repository.InsertClass(context.Background(), &errorClass))
}

func TestRepositoryGetClasses(t *testing.T) {
	var expectedClasses []*models.Class
	var filters = models.ClassFilter{}
	classes, err := repository.GetClasses(context.Background(), &filters)

	assert.Nil(t, err)
	assert.IsType(t, expectedClasses, classes)
}
