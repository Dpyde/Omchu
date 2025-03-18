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
	var users []entity.User
	swiped := r.db.Model(&entity.Swipe{}).
		Select("swiped_id").
		Where("swiper_id = ?", id)
	err := r.db.
		Where("id != ?", id).
		Where("id NOT IN (?)", swiped).
		Find(&users).Error

	if err != nil {
		return nil, err
	}
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
func (r *GormUserRepository) Update(newUser *entity.User, id uint) (*entity.User, error) {
	var user entity.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	// Map to store only non-zero values
	updates := make(map[string]interface{})

	if newUser.Name != "" {
		updates["name"] = newUser.Name
	}
	if newUser.Age != 0 { // Ensure Age is not 0 before updating
		updates["age"] = newUser.Age
	}
	if newUser.Email != "" {
		updates["email"] = newUser.Email
	}
	if newUser.Color != "" {
		updates["color"] = newUser.Color
	}
	if newUser.Password != "" {
		updates["password"] = newUser.Password
	}

	if len(updates) > 0 {
		if err := r.db.Model(&user).Omit("Chats", "Swipes").Updates(updates).Error; err != nil {
			return nil, err
		}
	}

	if len(newUser.Pictures) > 0 {
		if err := r.db.Model(&user).Association("Pictures").Replace(newUser.Pictures); err != nil {
			return nil, err
		}
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

// func (r *GormUserRepository) FindByUsernameGORM(username string) (*entity.User, error) {
// 	var user entity.User
// 	if err := r.db.Where("username = ?", username).First(&user).Error; err != nil {
// 		return nil, errors.New("user not found")
// 	}
// 	return &user, nil
// }
// func (r *GormUserRepository) FindByEmailGORM(email string) (*entity.User, error) {
// 	var user entity.User
// 	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
// 		return nil, errors.New("user not found")
// 	}
// 	return &user, nil
// }
