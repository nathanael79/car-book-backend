package repository

import (
	"book-car/model"
	"book-car/pkg/pagination"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AgendaRepository struct {
	db *gorm.DB
}

func AgendaRepositoryImpl(db *gorm.DB) *AgendaRepository {
	return &AgendaRepository{db: db}
}

func (ct *AgendaRepository) FindCarByStartDatendEndDate(
	carTypeID uuid.UUID,
	startDate time.Time,
	endDate time.Time,
	pagination pagination.Pagination,
) (*[]model.Agenda, error) {
	var agendas []model.Agenda

	result := ct.db.Model(&model.Agenda{}).
		Scopes(pagination.Scope()).
		Select("cars.id, cars.name").
		Joins("LEFT JOIN agendas a ON a.car_id = cars.id AND a.start_date <= ? AND a.end_date >= ?", endDate, startDate)

	if carTypeID != uuid.Nil {
		result = result.Where("cars.car_type_id = ?", carTypeID)
	}

	result = result.Where("a.id IS NULL")

	if err := result.Find(&agendas).Error; err != nil {
		return nil, err
	}

	return &agendas, nil
}

func (ct *AgendaRepository) FindOne(ID uuid.UUID) (*model.Agenda, error) {
	var agenda model.Agenda

	result := ct.db.Debug().Preload("User").Preload("Car").Where("id = ?", ID).First(&agenda)

	if result.Error != nil {
		return nil, result.Error
	}

	return &agenda, nil
}

func (ct *AgendaRepository) FindAll(pagination *pagination.Pagination) (*[]model.Agenda, error) {
	var agendas []model.Agenda

	result := ct.db.Preload("User").Preload("Car").Scopes(pagination.Scope()).Find(&agendas)

	if result.Error != nil {
		return nil, result.Error
	}

	return &agendas, nil
}

func (ct *AgendaRepository) FindAllAgendasByUserID(userID uuid.UUID, pagination *pagination.Pagination) (*[]model.Agenda, error) {
	var agendas []model.Agenda

	result := ct.db.Preload("User").Preload("Car").Scopes(pagination.Scope()).Where("user_id = ?", userID).Find(&agendas)

	if result.Error != nil {
		return nil, result.Error
	}

	return &agendas, nil
}

func (ct *AgendaRepository) Create(agenda *model.Agenda) (*model.Agenda, error) {
	result := ct.db.Create(agenda)

	if result.Error != nil {
		return nil, result.Error
	}

	var active model.Agenda
	if err := ct.db.Preload("User").Preload("Car").First(&active, "id = ?", agenda.ID).Error; err != nil {
		return nil, err
	}

	return &active, nil
}

func (ct *AgendaRepository) Update(active *model.Agenda, patch *model.Agenda) (*model.Agenda, error) {
	result := ct.db.Model(active).Updates(patch)

	if result.Error != nil {
		return nil, result.Error
	}

	if err := ct.db.First(active, "id = ?", active.ID).Error; err != nil {
		return nil, err
	}

	return active, nil
}

func (ct *AgendaRepository) Delete(agenda *model.Agenda) error {
	result := ct.db.Delete(agenda)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
