package main

import (
	"book-car/controller"
	"book-car/model"
	"book-car/repository"
	"book-car/service"
	"book-car/service/authentication"
	"book-car/service/authentication/jwt"
	"fmt"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func LoadEnv(env string) {
	file := ".env." + env
	if err := godotenv.Load(file); err != nil {
		log.Printf("Warning: gagal load %s, pakai OS env", file)
	}
}

func main() {
	LoadEnv("dev")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	} else {
		fmt.Println("Connection Opened to Database")
	}

	db.AutoMigrate(&model.User{}, &model.CarBrand{}, &model.CarType{}, &model.Car{}, &model.Agenda{})

	userRepository := repository.UserRepositoryImpl(db)
	carBrandRepository := repository.CarBrandRepositoryImpl(db)
	authenticationService := authentication.AuthenticationServiceImpl(userRepository)
	carBrandService := service.CarBrandServiceImpl(carBrandRepository)
	authenticatonController := controller.AuthenticationControllerImpl(authenticationService)
	carBrandController := controller.CarBrandControllerImpl(carBrandService)

	router := gin.Default()

	router.Use(cors.Default())

	api := router.Group("/api")
	v1 := api.Group("/v1")

	v1.POST("/register", authenticatonController.Register)
	v1.POST("/login", authenticatonController.Login)

	protected := v1.Group("/")
	protected.Use(jwt.AuthMiddleware())
	{
		protected.GET("/me", authenticatonController.GetUserLoginInformation)
	}

	newCarBrandRoute := protected.Group("/car-brand")
	{
		newCarBrandRoute.GET("", carBrandController.FindAll)
		newCarBrandRoute.POST("/create", carBrandController.Create)
		newCarBrandRoute.PATCH("/update/:id", carBrandController.Update)
		newCarBrandRoute.GET("/:id", carBrandController.FindOneByID)
	}
	router.Run()
}
