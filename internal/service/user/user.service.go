package userSer

import (
	"errors"

	"github.com/Dpyde/Omchu/internal/entity"
	userRep "github.com/Dpyde/Omchu/internal/repository/user"
)

//primary port
type UserService interface {
	CreateUser(user entity.User) error
}

type userServiceImpl struct {
	repo userRep.UserRepository
  }
  
  func NewUserService(repo userRep.UserRepository) UserService {
	return &userServiceImpl{repo: repo}
  }
  
  func (s *userServiceImpl) CreateUser(user entity.User) error {
	if user.Age <= 18 {
	  return errors.New("PM might hungry")
	}
	// Business logic...
	if err := s.repo.Save(user); err != nil {
	  return err
	}
	return nil
  }