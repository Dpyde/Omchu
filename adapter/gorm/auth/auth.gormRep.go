package authGormRep

import (
	"errors"
	"fmt"

	"github.com/Dpyde/Omchu/internal/entity"
	authRep "github.com/Dpyde/Omchu/internal/repository/auth"
	authSer "github.com/Dpyde/Omchu/internal/service/auth"
	"gorm.io/gorm"
)

type GormAuthRepository struct {
	db *gorm.DB
}

func NewGormAuthRepository(db *gorm.DB) authRep.AuthRepository {
	return &GormAuthRepository{db: db}
}

func (repo *GormAuthRepository) Log(username string, password string) (string, error) {
	// Retrieve the user from the database
	var user entity.User
	if err := repo.db.Where("username = ?", username).First(&user).Error; err != nil {
		return "", errors.New("user not found")
	}

	// Compare the provided password with the stored hashed password
	if !authSer.ComparePassword(password, &user) {
		return "", errors.New("incorrect password")
	}

	// Generate a JWT token if the password is correct
	token, err := authSer.GenerateToken(fmt.Sprint(user.ID))
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %v", err)
	}
	return token, nil
}
func (repo *GormAuthRepository) Reg(newUser *entity.User) (*entity.User, error) {
	// Check if the username already exists

	if err := repo.db.Create(newUser).Error; err != nil {
		return nil, fmt.Errorf("failed to create user: %v", err)
	}

	return newUser, nil
}
