package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	BaseModel
	Email    string `gorm:"uniqueIndex;not null"`
	Password string `gorm:"not null"`
}

func (u *User) BeforeSave(db *gorm.DB) (err error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hashed)
	return nil
}
