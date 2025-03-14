package userSer

import (
	"github.com/Dpyde/Omchu/internal/entity"
	userRep "github.com/Dpyde/Omchu/internal/repository/user"
	authSer "github.com/Dpyde/Omchu/internal/service/auth" // Ensure this path is correct and the package exists
)

// primary port
type UserService interface {
	CreateUser(user entity.User) (*entity.User, error)
	FindUsersToSwipe(id uint) (*[]entity.User, error)
	FindByID(id uint) (*entity.User, error)
	FindByUsername(username string) (*entity.User, error)
	FindByEmail(email string) (*entity.User, error)
	UpdateUser(newUser entity.User, id uint) (*entity.User, error)
	RemoveUser(user entity.User) error
}

type userServiceImpl struct {
	repo userRep.UserRepository
}

func NewUserService(repo userRep.UserRepository) UserService {
	return &userServiceImpl{repo: repo}
}

func (s *userServiceImpl) CreateUser(user entity.User) (*entity.User, error) {
	// Business logic...
	hashedPassword, err := authSer.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = hashedPassword
	createdUser, err := s.repo.Save(&user)
	if err != nil {
		return nil, err
	}
	return createdUser, nil
}
func (s *userServiceImpl) FindUsersToSwipe(id uint) (*[]entity.User, error) {
	users, err := s.repo.FindUsersToSwipe(id)
	if err != nil {
		return nil, err
	}
	return users, nil
}
func (s *userServiceImpl) FindByID(id uint) (*entity.User, error) {
	user, err := s.repo.FindByIDGORM(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (s *userServiceImpl) FindByUsername(username string) (*entity.User, error) {
	user, err := s.repo.FindByUsernameGORM(username)
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (s *userServiceImpl) FindByEmail(email string) (*entity.User, error) {
	user, err := s.repo.FindByEmailGORM(email)
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (s *userServiceImpl) UpdateUser(newUser entity.User, id uint) (*entity.User, error) {
	updatedUser, err := s.repo.Update(&newUser, id)
	if err != nil {
		return nil, err
	}
	return updatedUser, nil
}
func (s *userServiceImpl) RemoveUser(user entity.User) error {
	if err := s.repo.Remove(&user); err != nil {
		return err
	}
	return nil
}
