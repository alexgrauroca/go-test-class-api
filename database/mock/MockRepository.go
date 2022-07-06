package mockdb

type MockRepository struct{}

func NewMockRepository() (*MockRepository, error) {
	return &MockRepository{}, nil
}

func (repo *MockRepository) Close() error {
	return nil
}
