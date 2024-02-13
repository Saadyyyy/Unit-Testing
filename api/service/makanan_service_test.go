package service

import (
	"errors"
	"testing"

	"github.com/Saadyyyy/Unit-Testing/api/service/mocks"
	"github.com/Saadyyyy/Unit-Testing/models"
	"github.com/stretchr/testify/assert"
)

func Test_Create(t *testing.T) {
	mockRepo := new(mocks.MakanRepositoryMock)
	service := NewMakananService(mockRepo)

	t.Run("success", func(t *testing.T) {
		//membuat mocking data
		mockData := models.Makanan{
			Id:   uint(1),
			Nama: "aisah",
		}

		mockRepo.Mock.On("Create", mockData).Return(mockData, nil)

		result, err := service.Create(mockData)
		assert.NotNil(t, result)
		assert.Nil(t, err)
		assert.Equal(t, &mockData, result)
		mockRepo.Mock.AssertExpectations(t)
	})

	t.Run("Failed", func(t *testing.T) {
		//membuat mocking data
		mockData := models.Makanan{
			Id:   uint(2),
			Nama: "aisah",
		}
		mockRepo.Mock.On("Create", mockData).Return(nil, errors.New("Gagal Membuat makanan"))

		_, err := service.Create(mockData)
		// assert.Nil(t, result)
		assert.NotNil(t, err)
		mockRepo.Mock.AssertExpectations(t)
	})

}
