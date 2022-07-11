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

	// Testing diferent ranges, to validate every possible combination
	assert.NotNil(t, errorClass)
	assert.Nil(t, errorClass.NewId())
	assert.NotEmpty(t, errorClass.Id)
	assert.NotNil(t, repository.InsertClass(context.Background(), &errorClass))

	errorClass.Id = ""
	errorClass.EndDate = c.StartDate
	errorClass.StartDate = c.StartDate.AddDate(0, 0, -2)

	assert.NotNil(t, errorClass)
	assert.Nil(t, errorClass.NewId())
	assert.NotEmpty(t, errorClass.Id)
	assert.NotNil(t, repository.InsertClass(context.Background(), &errorClass))

	errorClass.Id = ""
	errorClass.EndDate = c.EndDate.AddDate(0, 0, 2)
	errorClass.StartDate = c.EndDate

	assert.NotNil(t, errorClass)
	assert.Nil(t, errorClass.NewId())
	assert.NotEmpty(t, errorClass.Id)
	assert.NotNil(t, repository.InsertClass(context.Background(), &errorClass))

	errorClass.Id = ""
	errorClass.EndDate = c.EndDate.AddDate(0, 0, 2)
	errorClass.StartDate = c.EndDate

	assert.NotNil(t, errorClass)
	assert.Nil(t, errorClass.NewId())
	assert.NotEmpty(t, errorClass.Id)
	assert.NotNil(t, repository.InsertClass(context.Background(), &errorClass))

	errorClass.Id = ""
	errorClass.StartDate = c.EndDate.AddDate(0, 0, 1)

	assert.NotNil(t, errorClass)
	assert.Nil(t, errorClass.NewId())
	assert.NotEmpty(t, errorClass.Id)
	assert.Nil(t, repository.InsertClass(context.Background(), &errorClass))
}

func TestRepositoryGetClasses(t *testing.T) {
	var expectedClasses []*models.Class
	var filters = models.ClassFilter{}
	classes, err := repository.GetClasses(context.Background(), &filters)

	assert.Nil(t, err)
	assert.IsType(t, expectedClasses, classes)
}
