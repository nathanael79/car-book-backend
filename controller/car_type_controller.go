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

type CarTypeController struct {
	carTypeService *service.CarTypeService
}

func CarTypeControllerImpl(carTypeService *service.CarTypeService) *CarTypeController {
	return &CarTypeController{carTypeService: carTypeService}
}

func (controller *CarTypeController) FindAll(ctx *gin.Context) {
	pageNumber, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	sizeNumber, _ := strconv.Atoi(ctx.DefaultQuery("size", "10"))

	pagination := pagination.New(pageNumber, sizeNumber)
	carTypes, err := controller.carTypeService.FindAll(&pagination)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": carTypes})
}

func (controller *CarTypeController) FindOneByID(ctx *gin.Context) {
	id := ctx.Param("id")
	uuid, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID format"})
		return
	}

	carType, err := controller.carTypeService.FindOneByID(uuid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": carType})
}

func (controller *CarTypeController) Create(ctx *gin.Context) {
	var carTypeRequest dto.CarTypeRequest

	if err := ctx.ShouldBindJSON(&carTypeRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	validate := validator.New()
	err := validate.Struct(&carTypeRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	newCarType, err := controller.carTypeService.Create(&carTypeRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": newCarType,
	})
}

func (controller *CarTypeController) Update(ctx *gin.Context) {
	var carTypeRequest dto.CarTypeRequest
	id := ctx.Param("id")
	uuid, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID format"})
		return
	}

	if err := ctx.ShouldBindJSON(&carTypeRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	validate := validator.New()
	validationError := validate.Struct(&carTypeRequest)
	if validationError != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": validationError.Error(),
		})

		return
	}

	updatedCarType, err := controller.carTypeService.Update(uuid, &carTypeRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": updatedCarType,
	})
}

func (controller *CarTypeController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	uuid, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID format"})
		return
	}

	deleteError := controller.carTypeService.Delete(uuid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": deleteError.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": "Car type deleted successfully",
	})
}
