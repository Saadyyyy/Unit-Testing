package repository

import (
	"github.com/Saadyyyy/Unit-Testing/models"
	"gorm.io/gorm"
)

type MakananRepository interface {
	Create(makan models.Makanan) (*models.Makanan, error)
}

type MakananRepositoryImpl struct {
	db *gorm.DB
}

func NewMakananRepository(db *gorm.DB) MakananRepository {
	return &MakananRepositoryImpl{db: db}
}

func (ur *MakananRepositoryImpl) Create(makan models.Makanan) (*models.Makanan, error) {
	if err := ur.db.Create(&makan).Error; err != nil {
		return nil, err
	}
	return &makan, nil
}
