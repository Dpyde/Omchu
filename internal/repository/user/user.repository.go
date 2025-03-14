package userRep

import (
	"github.com/Dpyde/Omchu/internal/entity"
)

type UserRepository interface {
	Save(user *entity.User) error
	Update(newUser *entity.User, id uint) (*entity.User, error)
	FindByIDGORM(id uint) (*entity.User, error)
	FindByUsernameGORM(username string) (*entity.User, error)
	FindByEmailGORM(email string) (*entity.User, error)
	Remove(user *entity.User) error
}
