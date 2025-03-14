package userGormRep

import (
	"errors"

	"github.com/Dpyde/Omchu/internal/entity"
	userRep "github.com/Dpyde/Omchu/internal/repository/user"
	"gorm.io/gorm"
)

// Secondary adapter
type GormUserRepository struct {
	db *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) userRep.UserRepository {
	return &GormUserRepository{db: db}
}

func (r *GormUserRepository) Save(user *entity.User) error {
	if result := r.db.Create(&user); result.Error != nil {
		// Handle database errors
		return result.Error
	}
	return nil
}

// func (r *GormUserRepository) FindByIDGORM(id uint) (*entity.User, error) {
func (r *GormUserRepository) FindByIDGORM(id uint) (*entity.User, error) {
	var user entity.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
func (r *GormUserRepository) FindByUsernameGORM(username string) (*entity.User, error) {
	var user entity.User
	if err := r.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, errors.New("user not found")
	}
	return &user, nil
}
func (r *GormUserRepository) FindByEmailGORM(email string) (*entity.User, error) {
	var user entity.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, errors.New("user not found")
	}
	return &user, nil
}
func (r *GormUserRepository) Update(newUser *entity.User, id uint) (*entity.User, error) {
	var user entity.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	if err := r.db.Model(&user).Updates(newUser).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
func (r *GormUserRepository) Remove(user *entity.User) error {
	if err := r.db.Delete(&user).Error; err != nil {
		return err
	}
	return nil
}
