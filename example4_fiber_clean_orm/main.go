package main

import (
	"log"

	"example4_fiber_clean_orm/config"
	"example4_fiber_clean_orm/controller"
	"example4_fiber_clean_orm/middleware"
	"example4_fiber_clean_orm/models"
	"example4_fiber_clean_orm/repository"
	"example4_fiber_clean_orm/routes"
	"example4_fiber_clean_orm/service"
	"example4_fiber_clean_orm/utils"

	"github.com/gofiber/fiber/v3"
)

func main() {
	utils.InitLogger()

	app := fiber.New()

	app.Use(middleware.RequestID())
	app.Use(middleware.LoggerMiddleware())
	app.Use(middleware.ErrorLogger())

	utils.Log.Info().Msg("Server starting...")

	config.ConnectDB()

	config.DB.AutoMigrate(&models.User{})

	userRepo := repository.NewUserRepository(config.DB)

	userService := service.NewUserService(userRepo)

	userController := controller.NewUserController(userService)

	routes.SetupRoutes(app, userController)

	log.Fatal(app.Listen(":3000"))
}
