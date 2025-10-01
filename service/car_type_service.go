package service

import (
	"book-car/dto"
	"book-car/model"
	"book-car/pkg/pagination"
	"book-car/repository"

	"github.com/google/uuid"
)

type CarTypeService struct {
	carTypeRepository  *repository.CarTypeRepository
	carBrandRepository *repository.CarBrandRepository
}

func CarTypeServiceImpl(carTypeRepository *repository.CarTypeRepository, carBrandRepository *repository.CarBrandRepository) *CarTypeService {
	return &CarTypeService{
		carTypeRepository:  carTypeRepository,
		carBrandRepository: carBrandRepository,
	}
}

func (cts *CarTypeService) FindOneByID(ID uuid.UUID) (*model.CarType, error) {
	result, err := cts.carTypeRepository.FindOne(ID)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (cts *CarTypeService) FindAll(pagination *pagination.Pagination) (*[]model.CarType, error) {
	result, err := cts.carTypeRepository.FindAll(pagination)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (cts *CarTypeService) Create(request *dto.CarTypeRequest) (*model.CarType, error) {
	_, err := cts.carBrandRepository.FindByID(request.CarBrandID)

	if err != nil {
		return nil, err
	}

	newCarType := &model.CarType{
		Name:       request.Name,
		CarBrandID: request.CarBrandID,
	}

	result, err := cts.carTypeRepository.Create(newCarType)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (cts *CarTypeService) Update(ID uuid.UUID, request *dto.CarTypeRequest) (*model.CarType, error) {
	activeCarType, err := cts.carTypeRepository.FindOne(ID)

	if err != nil {
		return nil, err
	}

	if activeCarType == nil {
		return nil, err
	}

	_, carBrandError := cts.carBrandRepository.FindByID(request.CarBrandID)

	if carBrandError != nil {
		return nil, err
	}

	patchCarType := &model.CarType{
		Name:       request.Name,
		CarBrandID: request.CarBrandID,
	}

	result, err := cts.carTypeRepository.Update(activeCarType, patchCarType)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (cts *CarTypeService) Delete(ID uuid.UUID) error {
	activeCarType, err := cts.carTypeRepository.FindOne(ID)

	if err != nil {
		return err
	}

	deleteError := cts.carTypeRepository.Delete(activeCarType)

	if deleteError != nil {
		return deleteError
	}

	return nil
}
