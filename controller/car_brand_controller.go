package controller

import (
	"book-car/dto"
	"book-car/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type CarBrandController struct {
	carBrandService *service.CarBrandService
}

func CarBrandControllerImpl(carBrandService *service.CarBrandService) *CarBrandController {
	return &CarBrandController{carBrandService: carBrandService}
}

func (cb *CarBrandController) FindAll(ctx *gin.Context) {
	carBrands, err := cb.carBrandService.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": carBrands,
	})
}

func (cb *CarBrandController) FindOneByID(ctx *gin.Context) {
	ID := ctx.Param("id")

	result, err := cb.carBrandService.FindOneByID(ID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "data not found",
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}

func (cb *CarBrandController) Create(ctx *gin.Context) {
	var carBrandRequest dto.CarBrandRequest

	if err := ctx.ShouldBindJSON(&carBrandRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	validate := validator.New()
	err := validate.Struct(carBrandRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	newCarBrand, err := cb.carBrandService.Create(&carBrandRequest)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": newCarBrand,
	})
}
