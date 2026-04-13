package routes

import (
	"example4_fiber_clean_orm/controller"

	"github.com/gofiber/fiber/v3"
)

func SetupRoutes(app *fiber.App, uc *controller.UserController) {
	api := app.Group("/api")

	api.Post("/users", uc.CreateUser)
	api.Get("/users", uc.GetUsers)
	api.Get("/users/:id", uc.GetUser)
	api.Put("/users/:id", uc.UpdateUser)
	api.Delete("/users/:id", uc.DeleteUser)
}
