package controller

import (
	"strconv"

	"example4_fiber_clean_orm/models"
	"example4_fiber_clean_orm/service"

	"github.com/gofiber/fiber/v3"
)

type UserController struct {
	service service.UserService
}

func NewUserController(s service.UserService) *UserController {
	return &UserController{s}
}

func (uc *UserController) CreateUser(c fiber.Ctx) error {
	user := new(models.User)

	if err := c.Bind().Body(user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	if err := uc.service.CreateUser(user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(user)
}

func (uc *UserController) GetUsers(c fiber.Ctx) error {
	users, _ := uc.service.GetUsers()

	return c.JSON(users)
}

func (uc *UserController) GetUser(c fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	user, err := uc.service.GetUser(uint(id))

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "not found"})
	}

	return c.JSON(user)
}

func (uc *UserController) UpdateUser(c fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	var input models.User

	if err := c.Bind().Body(&input); err != nil {
		return err
	}

	uc.service.UpdateUser(uint(id), &input)

	return c.JSON(fiber.Map{"message": "updated"})
}

func (uc *UserController) DeleteUser(c fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	uc.service.DeleteUser(uint(id))

	return c.JSON(fiber.Map{"message": "deleted"})
}
