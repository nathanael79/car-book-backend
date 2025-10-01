package dto

import "github.com/google/uuid"

type CarTypeRequest struct {
	Name       string    `json:"name" validate:"required"`
	CarBrandID uuid.UUID `json:"car_brand_id" validate:"required"`
}
