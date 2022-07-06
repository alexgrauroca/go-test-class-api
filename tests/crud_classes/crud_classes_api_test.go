package crud_classes_test

import (
	"net/http"
	"os"
	"testing"

	testhelpers "go-test-class-api/tests/helpers"

	"github.com/stretchr/testify/assert"
)

func TestApiCreateClass(t *testing.T) {
	classJson, err := os.Open("./data/class.json")
	defer classJson.Close()

	assert.Nil(t, err)

	if err == nil {
		req, _ := http.NewRequest("POST", "/api/v1/classes", classJson)
		response := testhelpers.TestRequest(req)

		assert.Equal(t, http.StatusCreated, response.Code)
		assert.NotEmpty(t, response.Body.String())
	}
}
