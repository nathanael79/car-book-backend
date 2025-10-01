package model

import (
	"time"

	"github.com/google/uuid"
)

type Car struct {
	BaseModel
	CarTypeID              uuid.UUID `gorm:"type:uuid;not null"`
	CarType                CarType   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	LicenseNumber          string    `gorm:"uniqueIndex;not null"`
	MachineFrameNumber     string    `gorm:"uniqueIndex;not null"`
	Color                  string    `gorm:"not null"`
	LicenseNumberExpiredAt time.Time `gorm:"not null"`
}
