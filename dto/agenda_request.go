package dto

import (
	"github.com/google/uuid"
)

type AgendaRequest struct {
	CarID     uuid.UUID `json:"car_id" binding:"required,uuid"`
	StartDate string    `json:"start_date" binding:"required,datetime=2006-01-02"`
	EndDate   string    `json:"end_date" binding:"required,datetime=2006-01-02"`
}
