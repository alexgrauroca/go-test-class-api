package mockdb

import "github.com/alexgrauroca/go-test-class-api/models"

type MockRepository struct {
	// This is used for API testing
	Classes  []*models.Class
	Bookings []*models.Booking
}

func NewMockRepository() (*MockRepository, error) {
	return &MockRepository{}, nil
}

func (repo *MockRepository) Close() error {
	return nil
}
