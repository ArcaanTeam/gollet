package usecases

import (
	"errors"
	"gollet/internal/entities"
	"gollet/internal/interfaces/payloads"
	"gollet/internal/repositories"
	"gollet/internal/utils"
)

type UserService interface {
	CreateUser(input payloads.CreateUserInput) (*entities.User, error)
	GetUserByID(id string) (*entities.User, error)
	GetUsers() ([]*entities.User, error)
	UpdateUser(id string, input payloads.UpdateUserInput) (*entities.User, error)
}

type userService struct {
	r repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userService{r: userRepo}
}

func (s *userService) CreateUser(
	input payloads.CreateUserInput,
) (*entities.User, error) {
	if len(input.Password) < 6 {
		return nil, errors.New("password must be at least 6 characters long")
	}

	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		return nil, err
	}

	var role string
	if !utils.IsValidRole(role) || input.Role == "" {
		role = entities.RoleUser // Default role
	} else {
		role = input.Role
	}

	user := entities.User{
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

func (s *userService) UpdateUser(id string, input payloads.UpdateUserInput) (*entities.User, error) {
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

func (s *userService) GetUserByID(id string) (*entities.User, error) {
	user, err := s.r.FindByID(id)
	return &user, err
}

func (s *userService) GetUsers() ([]*entities.User, error) {
	users, err := s.r.FindUsers()
	if err != nil {
		return []*entities.User{}, nil
	}
	return users, nil
}
