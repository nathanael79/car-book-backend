package repository

import (
	"book-car/model"
	"book-car/pkg/pagination"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CarRepository struct {
	db *gorm.DB
}

func CarRepositoryImpl(db *gorm.DB) *CarRepository {
	return &CarRepository{db: db}
}

func (cr *CarRepository) FindOne(id uuid.UUID) (*model.Car, error) {
	var car model.Car
	result := cr.db.Preload("CarType").Where("id = ?", id).First(&car)

	if result.Error != nil {
		return nil, result.Error
	}

	return &car, nil
}

func (cr *CarRepository) FindAll(pagination pagination.Pagination) ([]model.Car, error) {
	var cars []model.Car

	result := cr.db.Scopes(pagination.Scope()).Preload("CarType").Find(&cars)

	if result.Error != nil {
		return nil, result.Error
	}

	return cars, nil
}

func (cr *CarRepository) Create(car *model.Car) (*model.Car, error) {
	result := cr.db.Create(car)

	if result.Error != nil {
		return nil, result.Error
	}

	activeCar := cr.db.Preload("CarType").Preload("CarType.CarBrand").First(car, "id = ?", car.ID)

	if activeCar.Error != nil {
		return nil, activeCar.Error
	}

	return car, nil
}

func (cr *CarRepository) Update(activeCar *model.Car, patchCar *model.Car) (*model.Car, error) {
	result := cr.db.Model(activeCar).Updates(patchCar)

	if result.Error != nil {
		return nil, result.Error
	}

	if err := cr.db.First(activeCar, "id = ?", activeCar.ID).Error; err != nil {
		return nil, err
	}

	return activeCar, nil
}

func (cr *CarRepository) Delete(car *model.Car) error {
	result := cr.db.Delete(car)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
