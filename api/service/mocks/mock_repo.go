package mocks

import (
	"github.com/Saadyyyy/Unit-Testing/models"
	"github.com/stretchr/testify/mock"
)

type MakanRepositoryMock struct {
	Mock mock.Mock
}

func (repo *MakanRepositoryMock) Create(makan models.Makanan) (*models.Makanan, error) {
	arg := repo.Mock.Called(makan)
	if arg.Get(0) == nil {
		return nil, arg.Error(1)
	} else {
		makann := arg.Get(0).(models.Makanan)
		return &makann, nil
	}

}
