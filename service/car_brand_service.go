package service

import (
	"book-car/dto"
	"book-car/model"
	"book-car/repository"
)

type CarBrandService struct {
	carBrandRepository *repository.CarBrandRepository
}

func CarBrandServiceImpl(carBrandRepository *repository.CarBrandRepository) *CarBrandService {
	return &CarBrandService{carBrandRepository: carBrandRepository}
}

func (cb *CarBrandService) FindAll() (*[]model.CarBrand, error) {
	result, err := cb.carBrandRepository.FindAll()

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (cb *CarBrandService) FindOneByID(ID string) (*model.CarBrand, error) {
	result, err := cb.carBrandRepository.FindByID(ID)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (cb *CarBrandService) Create(carBrandRequest *dto.CarBrandRequest) (*model.CarBrand, error) {
	carBrandModel := model.CarBrand{Name: carBrandRequest.Name}

	result, err := cb.carBrandRepository.Create(carBrandModel)

	if err != nil {
		return nil, err
	}

	return result, nil
}
