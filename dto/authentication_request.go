package dto

type AuthenticationRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required,min=6"`
}
