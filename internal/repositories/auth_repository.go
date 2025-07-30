package repositories

import (
	"gollet/internal/entities"

	"gorm.io/gorm"
)

// HINT: interface is for bypassing DI
// FIXME: Move to auth service package
type AuthRepository interface {
	GetUserByEmail(email string) (*entities.User, error)
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{db: db}
}

func (r *authRepository) GetUserByEmail(email string) (*entities.User, error) {
	var user entities.User
	if err := r.db.First(&user, "email = ?", email).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
