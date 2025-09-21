package service

import (
	"book-car/dto"
	"book-car/model"
	"book-car/pkg/pagination"
	"book-car/repository"
	"log"
)

type CarBrandService struct {
	carBrandRepository *repository.CarBrandRepository
}

func CarBrandServiceImpl(carBrandRepository *repository.CarBrandRepository) *CarBrandService {
	return &CarBrandService{carBrandRepository: carBrandRepository}
}

func (cb *CarBrandService) FindAll(pagination pagination.Pagination) (*[]model.CarBrand, error) {
	result, err := cb.carBrandRepository.FindAll(pagination)

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

func (cb *CarBrandService) Update(ID string, carBrandRequest dto.CarBrandRequest) (*model.CarBrand, error) {
	activeCarBrand, err := cb.carBrandRepository.FindByID(ID)
	if err != nil {
		log.Println("car brand is not found")
		return nil, err
	}

	newCarBrandData := model.CarBrand{
		Name: carBrandRequest.Name,
	}

	updatedCarBrand, err := cb.carBrandRepository.Update(&newCarBrandData, activeCarBrand)
	if err != nil {
		return nil, err
	}

	return updatedCarBrand, nil
}

func (cb *CarBrandService) Delete(ID string) error {
	activeCarBrand, err := cb.carBrandRepository.FindByID(ID)
	if err != nil {
		log.Println("car brand is not found")
		return err
	}

	result := cb.carBrandRepository.Delete(activeCarBrand)
	if result != nil {
		return result
	}

	return nil
}
