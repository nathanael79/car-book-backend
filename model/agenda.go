package model

import (
	"time"

	"github.com/google/uuid"
)

type Agenda struct {
	BaseModel
	UserID    uuid.UUID `gorm:"type:uuid;not null"`
	User      User      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CarID     uuid.UUID `grom:"type:uuid;not null"`
	Car       Car       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Qty       uint8     `gorm:"not null"`
	StartDate time.Time `gorm:"not null"`
	EndDate   time.Time `gorm:not null`
}
