package repository

import (
	"book-car/model"
	"book-car/pkg/pagination"

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

func (cb *CarBrandRepository) FindAll(pagination pagination.Pagination) (*[]model.CarBrand, error) {
	var carBrands []model.CarBrand

	result := cb.db.Scopes(pagination.Scope()).Find(&carBrands)

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

func (cb *CarBrandRepository) Update(patch *model.CarBrand, active *model.CarBrand) (*model.CarBrand, error) {
	// Update field non-zero dari patch ke row active (berdasar PK di active)
	if err := cb.db.Model(active).Updates(patch).Error; err != nil {
		return nil, err
	}

	// Reload agar UpdatedAt / trigger DB terambil (opsional jika DB tak support RETURNING)
	if err := cb.db.First(active, "id = ?", active.ID).Error; err != nil {
		return nil, err
	}

	return active, nil
}

func (cb *CarBrandRepository) Delete(carBrand *model.CarBrand) error {
	result := cb.db.Delete(carBrand)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
