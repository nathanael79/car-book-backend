package controller

import (
	"book-car/dto"
	"book-car/pkg/pagination"
	"book-car/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AgendaController struct {
	agendaService *service.AgendaService
}

func AgendaControllerImpl(
	agendaService *service.AgendaService,
) *AgendaController {
	return &AgendaController{
		agendaService: agendaService,
	}
}

func (ac *AgendaController) CreateAgenda(ctx *gin.Context) {
	var agendaRequest dto.AgendaRequest

	if err := ctx.ShouldBindJSON(&agendaRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	result, err := ac.agendaService.CreateAgenda(ctx, agendaRequest)

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

func (ac *AgendaController) FindCarByStartDatendEndDate(ctx *gin.Context) {
	var carFindByTimeRequest dto.CarFindByTimeRequest

	if err := ctx.ShouldBindJSON(carFindByTimeRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	pageNumber, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	sizeNumber, _ := strconv.Atoi(ctx.DefaultQuery("size", "10"))

	pagination := pagination.New(pageNumber, sizeNumber)

	result, err := ac.agendaService.FindAgendaByTime(
		carFindByTimeRequest,
		pagination,
	)

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
