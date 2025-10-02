package controller

import (
	"book-car/dto"
	"book-car/pkg/pagination"
	"book-car/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type CarController struct {
	carService *service.CarService
}

func CarControllerImpl(carService *service.CarService) *CarController {
	return &CarController{carService: carService}
}

func (cc *CarController) FindOneByID(ctx *gin.Context) {
	id := ctx.Param("id")
	uuid, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID format"})
		return
	}

	result, err := cc.carService.FindOne(uuid)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": result})
}

func (cc *CarController) FindAll(ctx *gin.Context) {
	page := ctx.DefaultQuery("page", "1")
	size := ctx.DefaultQuery("size", "10")

	pageNumber, _ := strconv.Atoi(page)
	sizeNumber, _ := strconv.Atoi(size)

	pagination := pagination.New(pageNumber, sizeNumber)

	result, err := cc.carService.FindAll(pagination)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": result})
}

func (cc *CarController) Create(ctx *gin.Context) {
	var carRequest dto.CarRequest

	if err := ctx.ShouldBindJSON(&carRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	validate := validator.New()
	err := validate.Struct(&carRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	result, err := cc.carService.Create(&carRequest)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": result})
}

func (cc *CarController) Update(ctx *gin.Context) {
	var carRequest dto.CarRequest
	id := ctx.Param("id")
	uuid, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID format"})
		return
	}

	if err := ctx.ShouldBindJSON(&carRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	validate := validator.New()
	validationError := validate.Struct(&carRequest)
	if validationError != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": validationError.Error(),
		})

		return
	}

	updatedCar, err := cc.carService.Update(uuid, &carRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": updatedCar,
	})
}

func (cc *CarController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	uuid, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID format"})
		return
	}

	err = cc.carService.Delete(uuid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": "Car deleted successfully"})
}
