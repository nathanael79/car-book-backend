package model

import (
	"time"

	"github.com/google/uuid"
)

type Agenda struct {
	BaseModel
	UserID    uuid.UUID `gorm:"type:uuid;not null"`
	User      User      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CarID     uuid.UUID `grom:"type:uuid;not null"`
	Car       Car       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Qty       uint8     `gorm:"not null"`
	StartDate time.Time `gorm:"not null"`
	EndDate   time.Time `gorm:"not null"`
	Status    string    `gorm:"not null"`
}
