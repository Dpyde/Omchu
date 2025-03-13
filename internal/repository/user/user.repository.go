package userRep

import (
	"github.com/Dpyde/Omchu/internal/entity"
)

type UserRepository interface {
	Save(user entity.User) error
	FindByID(id uint) (*entity.User, error)
	FindByUsername(username string) (*entity.User, error)
}
