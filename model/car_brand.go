package model

type CarBrand struct {
	BaseModel
	Name string `gorm:"uniqueIndex;not null"`
}
