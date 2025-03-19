package user

import (
	"errors"

	"github.com/Dpyde/Omchu/internal/entity"
	"gorm.io/gorm"
)

// Secondary adapter

// Secondary adapter
type GormUserRepository struct {
	db *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) UserRepository {
	return &GormUserRepository{db: db}
}

func (r *GormUserRepository) CreateUser(user *entity.User) (*entity.User, error) {
	if err := r.db.Create(user).Error; err != nil {
		return nil, errors.New("failed to create user")
	}
	// fmt.Println("checkPoint4")
	return user, nil
}
func (r *GormUserRepository) FindUsersToSwipe(id uint) (*[]entity.User, error) {
	var user entity.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	swiped := r.db.Model(&entity.Swipe{}).
		Select("swiped_id").
		Where("wiper_id = ?", id)
	var users []entity.User
	r.db.Where("id != ?", id).
		Where("id NOT IN (?)", swiped).
		Find(&users)

	return &users, nil
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
func (r *GormUserRepository) Remove(id uint) error {
	var user entity.User
	if err := r.db.First(&user, id).Error; err != nil {
		return err
	}
	if err := r.db.Delete(&user).Error; err != nil {
		return err
	}
	return nil
}
