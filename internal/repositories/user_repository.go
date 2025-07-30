package repositories

import (
	"gollet/internal/entities"

	"gorm.io/gorm"
)

type UpdateUserInput struct{}

type UserRepository interface {
	Create(user *entities.User) error
	FindByID(ID string) (entities.User, error)
	FindUsers() ([]*entities.User, error)
	FindByEmail(email string) (entities.User, error)
	Update(user *entities.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *entities.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) FindByID(ID string) (entities.User, error) {
	var user entities.User
	err := r.db.First(&user, ID).Error
	return user, err
}

func (r *userRepository) FindByEmail(email string) (entities.User, error) {
	var user entities.User
	err := r.db.First(&user, "email = ?", email).Error
	return user, err
}

func (r *userRepository) Update(user *entities.User) error {
	return r.db.Save(user).Error
}

func (r *userRepository) FindUsers() ([]*entities.User, error) {
	var users []*entities.User
	if err := r.db.Find(&users).Error; err != nil {
		return []*entities.User{}, nil
	}
	return users, nil
}
