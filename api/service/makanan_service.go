package service

import (
	"github.com/Saadyyyy/Unit-Testing/api/repository"
	"github.com/Saadyyyy/Unit-Testing/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type MakananService interface {
	Create(ctx *gin.Context) (*models.Makanan, error)
}

type MakananServiceImpl struct {
	repo repository.MakananRepository
}

func NewMakananService(repo repository.MakananRepository) MakananService {
	return &MakananServiceImpl{repo: repo}
}

func (us *MakananServiceImpl) Create(ctx *gin.Context) (*models.Makanan, error) {
	input := models.Makanan{}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		return nil, err
	}

	validate := validator.New()
	err := validate.Struct(input)
	if err != nil {
		return nil, err
	}

	result, err := us.repo.Create(input)
	if err != nil {
		return nil, err
	}

	return result, nil
}
