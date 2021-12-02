package repos_test

import (
	"testing"

	"github.com/PizzaStruct/TenseiAPI/internal/pkg/models"
	"github.com/stretchr/testify/mock"
)

type MockGenreRepo struct {
	mock.Mock
}

func (m *MockGenreRepo) GetGenres() []models.Genre {
	args := m.Called()
	return args.Get(0).([]models.Genre)
}

func (m *MockGenreRepo) InsertGenre(genre string) error {
	args := m.Called(genre)
	return args.Error(0)
}

func (m *MockGenreRepo) RemoveGenre(id_hex string) error {
	args := m.Called(id_hex)
	return args.Error(0)
}

func TestGetGenres(t *testing.T) {
	mockobj := new(MockGenreRepo)
	mockobj.On("GetGenres").Return([]models.Genre{})
	mockobj.GetGenres()
	mockobj.AssertExpectations(t)
}
