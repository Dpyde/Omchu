package authGormRep

import (
	"errors"

	"github.com/Dpyde/Omchu/internal/entity"
	authRep "github.com/Dpyde/Omchu/internal/repository/auth"
	"gorm.io/gorm"
)

type GormAuthRepository struct {
	db *gorm.DB
}

func NewGormAuthRepository(db *gorm.DB) authRep.AuthRepository {
	return &GormAuthRepository{db: db}
}

func (r *GormAuthRepository) Reg(newUser *entity.User) (*entity.User, error) {
	// Check if the username already exists
	if err := r.db.Create(newUser).Error; err != nil {
		return nil, errors.New("failed to create user")
	}
	return newUser, nil
}

func (r *GormAuthRepository) Log(email string) (*entity.User, error) {
	var user entity.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, errors.New("user not found")
	}

	return &user, nil
}
