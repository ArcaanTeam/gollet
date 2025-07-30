package services

import (
	"errors"
	"gollet/api/payloads"
	"gollet/models"
	"gollet/repositories"
	"gollet/utils"
)

type UserService interface {
	CreateUser(input payloads.CreateUserInput) (*models.User, error)
	GetUserByID(id string) (*models.User, error)
	GetUsers() ([]*models.User, error)
	UpdateUser(id string, input payloads.UpdateUserInput) (*models.User, error)
}

type userService struct {
	r repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userService{r: userRepo}
}

func (s *userService) CreateUser(
	input payloads.CreateUserInput,
) (*models.User, error) {
	if len(input.Password) < 6 {
		return nil, errors.New("password must be at least 6 characters long")
	}

	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		return nil, err
	}

	var role string
	if !utils.IsValidRole(role) || input.Role == "" {
		role = models.RoleUser // Default role
	} else {
		role = input.Role
	}

	user := models.User{
		Name:         input.Name,
		Email:        input.Email,
		PasswordHash: hashedPassword,
		Role:         role,
	}

	if err := s.r.Create(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *userService) UpdateUser(id string, input payloads.UpdateUserInput) (*models.User, error) {
	userToUpdate, err := s.r.FindByID(id)
	if err != nil {
		return nil, err
	}

	if input.Name != "" {
		userToUpdate.Name = input.Name
	}
	if input.Email != "" {
		userToUpdate.Email = input.Email
	}
	if input.Role != "" && utils.IsValidRole(input.Role) {
		userToUpdate.Role = input.Role
	}
	if input.Password != "" && len(input.Password) > 5 {
		hashedNewPassword, err := utils.HashPassword(input.Password)
		if err != nil {
			return nil, err
		}
		userToUpdate.PasswordHash = hashedNewPassword
	}

	if err := s.r.Update(&userToUpdate); err != nil {
		return nil, err
	}
	return &userToUpdate, nil
}

func (s *userService) GetUserByID(id string) (*models.User, error) {
	user, err := s.r.FindByID(id)
	return &user, err
}

func (s *userService) GetUsers() ([]*models.User, error) {
	users, err := s.r.FindUsers()
	if err != nil {
		return []*models.User{}, nil
	}
	return users, nil
}
