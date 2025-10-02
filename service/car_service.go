package service

import (
	"book-car/dto"
	"book-car/model"
	"book-car/pkg/pagination"
	"book-car/pkg/utils"
	"book-car/repository"

	"github.com/google/uuid"
)

type CarService struct {
	carRepository     *repository.CarRepository
	carTypeRepository *repository.CarTypeRepository
}

func CarServiceImpl(carRepo *repository.CarRepository, carTypeRepo *repository.CarTypeRepository) *CarService {
	return &CarService{
		carRepository:     carRepo,
		carTypeRepository: carTypeRepo,
	}
}

func (cs *CarService) FindOne(id uuid.UUID) (*model.Car, error) {
	result, err := cs.carRepository.FindOne(id)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (cs *CarService) FindAll(pagination pagination.Pagination) ([]model.Car, error) {
	result, err := cs.carRepository.FindAll(pagination)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (cs *CarService) Create(carRequest *dto.CarRequest) (*model.Car, error) {
	_, err := cs.carTypeRepository.FindOne(carRequest.CarTypeID)

	if err != nil {
		return nil, err
	}

	licenseNumberExpiredAt, err := utils.ConvertStringToDateTime(carRequest.LicenseNumberExpired)

	if err != nil {
		return nil, err
	}

	newCar := &model.Car{
		CarTypeID:              carRequest.CarTypeID,
		LicenseNumber:          carRequest.LicenseNumber,
		MachineFrameNumber:     carRequest.MachineFrameNumber,
		Color:                  carRequest.Color,
		LicenseNumberExpiredAt: licenseNumberExpiredAt,
	}

	result, err := cs.carRepository.Create(newCar)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (cs *CarService) Update(id uuid.UUID, carRequest *dto.CarRequest) (*model.Car, error) {
	activeCar, err := cs.carRepository.FindOne(id)

	if err != nil {
		return nil, err
	}

	if activeCar == nil {
		return nil, err
	}

	activeCarType, err := cs.carTypeRepository.FindOne(carRequest.CarTypeID)

	if err != nil {
		return nil, err
	}

	if activeCarType == nil {
		return nil, err
	}

	updateCar := &model.Car{
		CarTypeID:          carRequest.CarTypeID,
		LicenseNumber:      carRequest.LicenseNumber,
		MachineFrameNumber: carRequest.MachineFrameNumber,
		Color:              carRequest.Color,
	}

	licenseNumberExpiredAt, err := utils.ConvertStringToDateTime(carRequest.LicenseNumberExpired)

	if err != nil {
		return nil, err
	}

	updateCar.LicenseNumberExpiredAt = licenseNumberExpiredAt

	result, err := cs.carRepository.Update(activeCar, updateCar)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (cs *CarService) Delete(id uuid.UUID) error {
	activeCar, err := cs.carRepository.FindOne(id)

	if err != nil {
		return nil
	}

	if activeCar == nil {
		return nil
	}

	err = cs.carRepository.Delete(activeCar)

	if err != nil {
		return err
	}

	return nil
}
