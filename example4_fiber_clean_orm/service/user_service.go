package service

import (
	"errors"
	"example4_fiber_clean_orm/models"
	"example4_fiber_clean_orm/repository"
)

type UserService interface {
	CreateUser(user *models.User) error
	GetUsers() ([]models.User, error)
	GetUser(id uint) (*models.User, error)
	UpdateUser(id uint, updated *models.User) error
	DeleteUser(id uint) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo}
}

func (s *userService) CreateUser(user *models.User) error {

	if user.Email == "" {
		return errors.New("email required")
	}
	
	return s.repo.Create(user)
}

func (s *userService) GetUsers() ([]models.User, error) {
	return s.repo.FindAll()
}

func (s *userService) GetUser(id uint) (*models.User, error) {
	return s.repo.FindByID(id)
}

func (s *userService) UpdateUser(id uint, updated *models.User) error {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	user.Name = updated.Name
	user.Email = updated.Email

	return s.repo.Update(user)
}

func (s *userService) DeleteUser(id uint) error {
	return s.repo.Delete(id)
}
