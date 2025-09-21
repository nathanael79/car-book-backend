package dto

type CarBrandRequest struct {
	Name string `json:"name" validate:"required"`
}
