package user

import (
	"github.com/Dpyde/Omchu/internal/entity"
)

type UserRepository interface {
	CreateUser(user *entity.User) (*entity.User, error)
	Update(newUser *entity.User, id uint) (*entity.User, error)
	GetUsersToSwipe(id uint) (*[]entity.User, error)
	FindByIDGORM(id uint) (*entity.User, error)
	Remove(id uint) error
	// FindByUsernameGORM(username string) (*entity.User, error)
	// FindByEmailGORM(email string) (*entity.User, error)
}
