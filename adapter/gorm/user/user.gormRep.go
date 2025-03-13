package userGormRep

import (
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

func (r *GormUserRepository) Save(user entity.User) error {
	if result := r.db.Create(&user); result.Error != nil {
		// Handle database errors
		return result.Error
	}
	return nil
}
