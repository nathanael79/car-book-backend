package dto

import (
	"time"

	"github.com/google/uuid"
)

type CarFindByTimeRequest struct {
	CarTypeID uuid.UUID `json:"car_type_id" binding:"omitempty,uuid"`
	StartDate time.Time `json:"start_date" binding:"required" time_format:"datetime=2006-01-02"`
	EndDate   time.Time `json:"end_date" binding:"required" time_format:"datetime=2006-01-02"`
}
