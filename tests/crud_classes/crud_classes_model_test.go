package crud_classes_test

import (
	"context"
	"testing"

	"github.com/alexgrauroca/go-test-class-api/repository"

	"github.com/stretchr/testify/assert"
)

func TestCreateClass(t *testing.T) {
	assert.NotNil(t, c)
	assert.Nil(t, c.NewId())
	assert.NotEmpty(t, c.Id)
	assert.Nil(t, repository.InsertClass(context.Background(), &c))
}
