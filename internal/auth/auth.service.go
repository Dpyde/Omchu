package auth

import (
	"errors"
	"os"
	"regexp"
	"time"

	"github.com/Dpyde/Omchu/internal/entity"

	jwt "github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Register(username string, email string, password string, age uint) (*entity.User, error)
	Login(email string, password string) (*entity.User, error)
	// TokenToId(token string) (string, error)
}

type authServiceImpl struct {
	repo AuthRepository
}

func NewAuthService(repo AuthRepository) AuthService {
	return &authServiceImpl{repo: repo}
}

func constraintCheck(email string, password string, age uint) error {
	var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

	// Password rules: At least 8 characters, one uppercase, one lowercase, one number
	var passwordRegex = regexp.MustCompile(`^(?=.*[a-z])(?=.*[A-Z])(?=.*\d).{8,}$`)
	// Validate email
	if !emailRegex.MatchString(email) {
		return errors.New("invalid email format")
	}

	// Validate password
	if !passwordRegex.MatchString(password) {
		return errors.New("password must be at least 8 characters, include one uppercase, one lowercase, and one number")
	}
	if age < 18 {
		return errors.New("PM might hungry")
	}

	return nil
}
func (s *authServiceImpl) Login(email string, password string) (*entity.User, error) {
	user, err := s.repo.Log(email)
	if err != nil {
		return nil, errors.New("user not found")
	}
	if !ComparePassword(password, user) {
		return nil, errors.New("Invalid credentials")
	}
	return user, nil
}
func (s *authServiceImpl) Register(username string, email string, password string, age uint) (*entity.User, error) {
	_, err := s.repo.Log(email)
	if err == nil {
		return nil, errors.New("email already taken")
	}
	if err := constraintCheck(email, password, age); err != nil {
		return nil, err
	}
	hashedPassword, err := HashPassword(password)
	if err != nil {
		return nil, err
	}
	newUser := &entity.User{
		Name:     username,
		Email:    email,
		Password: hashedPassword,
		Age:      age,
	}
	if _, err := s.repo.Reg(newUser); err != nil {
		return nil, err
	}
	return newUser, nil
}

func TokenToId(token string) (string, error) {
	secret := os.Getenv("JWT_SECRET")
	t, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return "", err
	}
	claims, ok := t.Claims.(jwt.MapClaims)
	if !ok || !t.Valid {
		return "", err
	}
	return claims["id"].(string), nil

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
