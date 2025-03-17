package userSer

import (
	"github.com/Dpyde/Omchu/internal/entity"
	userRep "github.com/Dpyde/Omchu/internal/repository/user" // Ensure this path is correct and the package exists
	authSer "github.com/Dpyde/Omchu/internal/service/auth"    // Ensure this path is correct and the package exists
)

// primary port
// primary port
type UserService interface {
	CreateUser(user entity.User) (*entity.User, error)
	FindUsersToSwipe(id uint) (*[]entity.User, error)
	FindByID(id uint) (*entity.User, error)
	FindByUsername(username string) (*entity.User, error)
	FindByEmail(email string) (*entity.User, error)
	UpdateUser(newUser entity.User, id uint) (*entity.User, error)
	RemoveUser(id uint) error
}

type userServiceImpl struct {
	repo userRep.UserRepository
}

func NewUserService(repo userRep.UserRepository) UserService {
}

func NewUserService(repo userRep.UserRepository) UserService {
	return &userServiceImpl{repo: repo}
}

func (s *userServiceImpl) CreateUser(user entity.User) (*entity.User, error) {
	hashedPassword, err := authSer.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = hashedPassword
	newUser, err := s.repo.CreateUser(&user)
	if err != nil {
		return nil, err
	}
	// fmt.Println("checkPoint3")
	return newUser, nil
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
func (s *userServiceImpl) RemoveUser(id uint) error {
	if err := s.repo.Remove(id); err != nil {
		return err
	}
	return nil
}

}
