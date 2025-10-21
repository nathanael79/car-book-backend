package service

import (
	"book-car/dto"
	"book-car/model"
	"book-car/pkg/pagination"
	"book-car/pkg/utils"
	"book-car/repository"
	"book-car/service/authentication/jwt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AgendaService struct {
	userRepository    *repository.UserRepository
	carTypeRepository *repository.CarTypeRepository
	carRepository     *repository.CarRepository
	agendaRepository  *repository.AgendaRepository
}

func AgendaServiceImpl(
	userRepository *repository.UserRepository,
	carTypeRepository *repository.CarTypeRepository,
	carRepository *repository.CarRepository,
	agendaRepository *repository.AgendaRepository,
) *AgendaService {
	return &AgendaService{
		userRepository:    userRepository,
		carTypeRepository: carTypeRepository,
		carRepository:     carRepository,
		agendaRepository:  agendaRepository,
	}
}

func (as *AgendaService) CreateAgenda(ctx *gin.Context, agendaRequest dto.AgendaRequest) (*model.Agenda, error) {
	claims := ctx.MustGet(jwt.ContextClaimsKey).(*jwt.UserClaims)

	userID := uuid.MustParse(claims.ID)

	parsedStartDate, err := utils.ConvertStringToDate(agendaRequest.StartDate)
	
	if err != nil {
		return nil, err
	}
	parsedEndDate, err := utils.ConvertStringToDate(agendaRequest.EndDate)
	
	if err != nil {
		return nil, err
	}

	newAgenda := &model.Agenda{
		UserID:    userID,
		CarID:     agendaRequest.CarID,
		StartDate: parsedStartDate,
		EndDate:   parsedEndDate,
	}

	result, err := as.agendaRepository.Create(newAgenda)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (as *AgendaService) FindAgendaByTime(carFindByTimeRequest dto.CarFindByTimeRequest, pagination pagination.Pagination) (*[]model.Agenda, error) {
	activeCarType, err := as.carTypeRepository.FindOne(carFindByTimeRequest.CarTypeID)

	if err != nil {
		return nil, err
	}

	result, err := as.agendaRepository.FindCarByStartDatendEndDate(activeCarType.ID, carFindByTimeRequest.StartDate, carFindByTimeRequest.EndDate, pagination)

	if err != nil {
		return nil, err
	}

	return result, nil

}

func (as *AgendaService) FindAllAgendasByUserID(ctx *gin.Context, pagination *pagination.Pagination) (*[]model.Agenda, error) {
	claims := ctx.MustGet(jwt.ContextClaimsKey).(*jwt.UserClaims)

	userID := uuid.MustParse(claims.ID)

	result, err := as.agendaRepository.FindAllAgendasByUserID(userID, pagination)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (asc *AgendaService) FindOneAgendaByID(agendaID uuid.UUID) (*model.Agenda, error) {
	result, err := asc.agendaRepository.FindOne(agendaID)

	if err != nil {
		return nil, err
	}

	return result, nil
}
