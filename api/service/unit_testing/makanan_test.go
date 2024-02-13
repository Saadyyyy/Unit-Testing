package unittesting

import (
	"testing"

	"github.com/Saadyyyy/Unit-Testing/api/service"
	"github.com/Saadyyyy/Unit-Testing/api/service/mocks"
	"github.com/Saadyyyy/Unit-Testing/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_CreateSucces(t *testing.T) {
	repo := new(mocks.MakanRepositoryMock)
	service := service.NewMakananService(repo)
	mockConext, _ := gin.CreateTestContext(nil)
	var data *models.Makanan

	createData := models.Makanan{
		Id:   uint(1),
		Nama: "laode",
	}
	repo.Mock.On("Create", createData).Return(data, nil)
	result, err := service.Create(mockConext)
	assert.NoError(t, err)
	// repo.Mock.AssertCalled(t, "Create", data)
	assert.Equal(t, data, result)
}
