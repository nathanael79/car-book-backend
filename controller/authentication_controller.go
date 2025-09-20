package controller

import (
	"book-car/dto"
	"book-car/service/authentication"
	"book-car/service/authentication/jwt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthenticationController struct {
	authenticationService *authentication.AuthenticationService
}

func AuthenticationControllerImpl(authenticationService *authentication.AuthenticationService) *AuthenticationController {
	return &AuthenticationController{authenticationService: authenticationService}
}

func (ac *AuthenticationController) Register(ctx *gin.Context) {
	var authenticationRequest dto.AuthenticationRequest

	if err := ctx.ShouldBindJSON(&authenticationRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	token, err := ac.authenticationService.Register(&authenticationRequest)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func (ac *AuthenticationController) Login(ctx *gin.Context) {
	var authenticationRequest dto.AuthenticationRequest

	if err := ctx.ShouldBindJSON(&authenticationRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	token, err := ac.authenticationService.Login(authenticationRequest)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func (ac *AuthenticationController) GetUserLoginInformation(ctx *gin.Context) {
	claims := ctx.MustGet(jwt.ContextClaimsKey).(*jwt.UserClaims)
	ctx.JSON(http.StatusOK, gin.H{
		"email": claims.Email,
		"exp":   claims.ExpiresAt.Time,
	})
}
