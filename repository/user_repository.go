package repository

import (
	"book-car/model"
	"log"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func UserRepositoryImpl(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (u *UserRepository) FindUserByEmail(email string) (*model.User, error) {
	var user model.User
	result := u.db.Where("email = ?", email).First(&user)

	if result.Error != nil {
		log.Println("cannot find user with email: ", email)
		return nil, result.Error
	}

	return &user, nil
}

func (u *UserRepository) CreateUser(newUser model.User) (*model.User, error) {
	result := u.db.Create(&newUser)

	if result.Error != nil {
		log.Println("cannot create user with error: ", result.Error)
		return nil, result.Error
	}

	return &newUser, nil
}
