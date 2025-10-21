package repository

import (
	"book-car/model"
	"book-car/pkg/pagination"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CarTypeRepository struct {
	db *gorm.DB
}

func CarTypeRepositoryImpl(db *gorm.DB) *CarTypeRepository {
	return &CarTypeRepository{db: db}
}

func (ct *CarTypeRepository) FindOne(ID uuid.UUID) (*model.CarType, error) {
	var carType model.CarType

	result := ct.db.Debug().Preload("CarBrand").Where("id = ?", ID).First(&carType)

	if result.Error != nil {
		return nil, result.Error
	}

	return &carType, nil
}

func (ct *CarTypeRepository) FindAll(pagination *pagination.Pagination) (*[]model.CarType, error) {
	var carTypes []model.CarType

	result := ct.db.Preload("CarBrand").Scopes(pagination.Scope()).Find(&carTypes)

	if result.Error != nil {
		return nil, result.Error
	}

	return &carTypes, nil
}

func (ct *CarTypeRepository) FindAllByCarBrandID(carBrandID uuid.UUID) (*[]model.CarType, error) {
	var carTypes []model.CarType

	result := ct.db.Where("car_brand_id = ?", carBrandID).Find(&carTypes)

	if result.Error != nil {
		return nil, result.Error
	}

	return &carTypes, nil
}

func (ct *CarTypeRepository) Create(carType *model.CarType) (*model.CarType, error) {
	result := ct.db.Create(carType)

	if result.Error != nil {
		return nil, result.Error
	}

	var active model.CarType
	if err := ct.db.Preload("CarBrand").First(&active, "id = ?", carType.ID).Error; err != nil {
		return nil, err
	}

	return &active, nil
}

func (ct *CarTypeRepository) Update(active *model.CarType, patch *model.CarType) (*model.CarType, error) {
	result := ct.db.Model(active).Updates(patch)

	if result.Error != nil {
		return nil, result.Error
	}

	if err := ct.db.First(active, "id = ?", active.ID).Error; err != nil {
		return nil, err
	}

	return active, nil
}

func (ct *CarTypeRepository) Delete(carType *model.CarType) error {
	result := ct.db.Delete(carType)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
