package authentication

import (
	"book-car/dto"
	"book-car/model"
	"book-car/repository"
	"book-car/service/authentication/jwt"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type AuthenticationService struct {
	userRepository *repository.UserRepository
}

func AuthenticationServiceImpl(userRepository *repository.UserRepository) *AuthenticationService {
	return &AuthenticationService{userRepository: userRepository}
}

func (a *AuthenticationService) Register(authenticationRequest *dto.AuthenticationRequest) (string, error) {
	activeUser, _ := a.userRepository.FindUserByEmail(authenticationRequest.Email)

	if activeUser != nil {
		return "", errors.New("user with this email already exists")
	}

	newUser := &model.User{
		Email:    authenticationRequest.Email,
		Password: authenticationRequest.Password,
	}

	activeUser, err := a.userRepository.CreateUser(*newUser)
	if err != nil {
		return "", err
	}

	token, err := jwt.CreateToken(activeUser.Email)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (a *AuthenticationService) Login(authenticationRequest dto.AuthenticationRequest) (string, error) {
	activeUser, err := a.userRepository.FindUserByEmail(authenticationRequest.Email)

	if err != nil {
		return "", errors.New("user not found")
	}

	checkPassword := bcrypt.CompareHashAndPassword([]byte(activeUser.Password), []byte(authenticationRequest.Password))

	if checkPassword != nil {
		return "", checkPassword
	}

	token, err := jwt.CreateToken(activeUser.Email)
	if err != nil {
		return "", err
	}

	return token, nil
}
