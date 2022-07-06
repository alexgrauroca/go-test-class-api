package crud_classes_test

import (
	"context"
	"go-test-class-api/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateClass(t *testing.T) {
	assert.NotNil(t, c)
	assert.Nil(t, c.NewId())
	assert.NotEmpty(t, c.Id)
	assert.Nil(t, repository.InsertClass(context.Background(), &c))
}

func TestReadClass(t *testing.T) {

}

func TestUpdateClass(t *testing.T) {

}

func TestDeleteClass(t *testing.T) {

}
