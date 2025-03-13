package service

import (
	"errors"

	"github.com/Dpyde/Omchu/internal/entity"
)

// Register a new user
func (s *AuthService) Register(username, password string) (*entity.User, error) {
	hashedPassword, err := HashPassword(password)
	if err != nil {
		return nil, err
	}

	user := &entity.User{Username: username, Password: hashedPassword}
	err = s.UserRepo.Create(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// Login checks credentials and returns JWT token
func (s *AuthService) Login(username, password string) (string, error) {
	user, err := s.UserRepo.FindByUsername(username)
	if err != nil || !ComparePassword(user.Password, password) {
		return "", errors.New("invalid credentials")
	}

	return GenerateToken(user.ID)
}
