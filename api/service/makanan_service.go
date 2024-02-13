package service

import (
	"github.com/Saadyyyy/Unit-Testing/api/repository"
	"github.com/Saadyyyy/Unit-Testing/models"
	"github.com/go-playground/validator/v10"
)

type MakananService interface {
	Create(input models.Makanan) (*models.Makanan, error)
}

type MakananServiceImpl struct {
	repo repository.MakananRepository
}

func NewMakananService(repo repository.MakananRepository) MakananService {
	return &MakananServiceImpl{repo: repo}
}

func (us *MakananServiceImpl) Create(input models.Makanan) (*models.Makanan, error) {

	validate := validator.New()
	validate.Struct(input)

	result, err := us.repo.Create(input)
	if err != nil {
		return nil, err
	}

	return result, nil
}
