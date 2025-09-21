package controller

import (
	"book-car/dto"
	"book-car/pkg/pagination"
	"book-car/service"
	"net/http"
	"strconv"

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
	pageNumber, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	sizeNumber, _ := strconv.Atoi(ctx.DefaultQuery("size", "10"))

	pagination := pagination.New(pageNumber, sizeNumber)
	carBrands, err := cb.carBrandService.FindAll(pagination)
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

func (cb *CarBrandController) Update(ctx *gin.Context) {
	var carBrandRequest dto.CarBrandRequest
	ID := ctx.Param("id")

	if err := ctx.ShouldBindJSON(&carBrandRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	result, err := cb.carBrandService.Update(ID, carBrandRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": result,
	})

}

func (cb *CarBrandController) Delete(ctx *gin.Context) {
	ID := ctx.Param("id")

	result := cb.carBrandService.Delete(ID)

	if result != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"error": result.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "deleted",
	})
}
