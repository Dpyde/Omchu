package service

import (
	"errors"
	"os"
	"time"

	"github.com/Dpyde/Omchu/internal/entity"
	authRep "github.com/Dpyde/Omchu/internal/repository/auth"

	jwt "github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Register(username string, email string, password string) (*entity.User, error)
	Login(email string, password string) error
}

type authServiceImpl struct {
	repo authRep.AuthRepository
}

func NewAuthService(repo authRep.AuthRepository) AuthService {
	return &authServiceImpl{repo: repo}
}

func (s *authServiceImpl) Login(email string, password string) error {
	user, err := s.repo.Log(email)
	if err != nil {
		return errors.New("user not found")
	}
	if !ComparePassword(password, user) {
		return errors.New("incorrect password")
	}
	return nil
}
func (s *authServiceImpl) Register(username string, email string, password string) (*entity.User, error) {
	_, err := s.repo.Log(email)
	if err == nil {
		return nil, errors.New("email already taken")
	}
	hashedPassword, err := HashPassword(password)
	if err != nil {
		return nil, err
	}
	newUser := &entity.User{
		Name:     username,
		Email:    email,
		Password: hashedPassword,
	}
	if _, err := s.repo.Reg(newUser); err != nil {
		return nil, err
	}
	return newUser, nil
}

// NOTE: The following functions are not used in the current implementation
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), err
}

func GenerateToken(userId string) (string, error) {
	secret := os.Getenv("JWT_SECRET")
	claims := jwt.MapClaims{
		"id":  userId,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}
func ComparePassword(password string, u *entity.User) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
