package main

import (
	"book-car/controller"
	"book-car/model"
	"book-car/repository"
	"book-car/service/authentication"
	"book-car/service/authentication/jwt"
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password=postgres dbname=car_book port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	} else {
		fmt.Println("Connection Opened to Database")
	}

	db.AutoMigrate(&model.User{}, &model.CarBrand{}, &model.CarType{}, &model.Car{}, &model.Agenda{})

	userRepository := repository.UserRepositoryImpl(db)
	authenticationService := authentication.AuthenticationServiceImpl(userRepository)
	authenticatonController := controller.AuthenticationControllerImpl(authenticationService)

	router := gin.Default()

	api := router.Group("/api")
	v1 := api.Group("/v1")

	v1.POST("/register", authenticatonController.Register)
	v1.POST("/login", authenticatonController.Login)

	protected := v1.Group("/")
	protected.Use(jwt.AuthMiddleware())
	{
		protected.GET("/me", authenticatonController.GetUserLoginInformation)
	}
	router.Run()
}
