package userRep

import "github.com/Dpyde/Omchu/internal/entity"

type AuthRepository interface {
	Reg(user *entity.User) (*entity.User, error)
	Log(email string) (*entity.User, error)
}
