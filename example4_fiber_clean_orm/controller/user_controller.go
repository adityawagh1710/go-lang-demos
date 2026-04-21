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

// parseID parses a uint ID from a route param, returning 0 and false on invalid/negative input.
func parseID(c fiber.Ctx) (uint, bool) {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil || id <= 0 {
		return 0, false
	}
	return uint(id), true //nolint:gosec // id is validated positive before conversion
}

func (uc *UserController) GetUser(c fiber.Ctx) error {
	id, ok := parseID(c)
	if !ok {
		return c.Status(400).JSON(fiber.Map{"error": "invalid id"})
	}

	user, err := uc.service.GetUser(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "not found"})
	}

	return c.JSON(user)
}

func (uc *UserController) UpdateUser(c fiber.Ctx) error {
	id, ok := parseID(c)
	if !ok {
		return c.Status(400).JSON(fiber.Map{"error": "invalid id"})
	}

	var input models.User
	if err := c.Bind().Body(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	if err := uc.service.UpdateUser(id, &input); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "updated"})
}

func (uc *UserController) DeleteUser(c fiber.Ctx) error {
	id, ok := parseID(c)
	if !ok {
		return c.Status(400).JSON(fiber.Map{"error": "invalid id"})
	}

	if err := uc.service.DeleteUser(id); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "deleted"})
}
