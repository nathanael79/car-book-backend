package model

import "github.com/google/uuid"

type CarType struct {
	BaseModel
	Name       string    `gorm:"uniqueIndex;not null"`
	CarBrandID uuid.UUID `gorm:"type:uuid;not null"`
	CarBrand   CarBrand  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
