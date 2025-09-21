package repository

import (
	"book-car/model"

	"gorm.io/gorm"
)

type CarBrandRepository struct {
	db *gorm.DB
}

func CarBrandRepositoryImpl(db *gorm.DB) *CarBrandRepository {
	return &CarBrandRepository{db: db}
}

func (cb *CarBrandRepository) FindByID(id string) (*model.CarBrand, error) {
	var carBrand model.CarBrand

	result := cb.db.Where("id = ?", id).First(&carBrand)

	if result.Error != nil {
		return nil, result.Error
	}

	return &carBrand, nil
}

func (cb *CarBrandRepository) FindAll() (*[]model.CarBrand, error) {
	var carBrands []model.CarBrand

	result := cb.db.Find(&carBrands)

	if result.Error != nil {
		return nil, result.Error
	}

	return &carBrands, nil
}

func (cb *CarBrandRepository) Create(carBrand model.CarBrand) (*model.CarBrand, error) {
	result := cb.db.Create(&carBrand)

	if result.Error != nil {
		return nil, result.Error
	}

	return &carBrand, nil
}

func (cb *CarBrandRepository) Update(carBrand model.CarBrand) (*model.CarBrand, error) {
	result := cb.db.Save(&carBrand)

	if result.Error != nil {
		return nil, result.Error
	}

	return &carBrand, nil

}

func (cb *CarBrandRepository) DeleteByID(carBrand *model.CarBrand) error {
	result := cb.db.Delete(carBrand)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
