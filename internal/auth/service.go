package auth

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

var jwtKey []byte

func init() {
	// Load JWT secret from environment variable
	jwtKey = []byte(os.Getenv("JWT_SECRET"))
}

type AuthService struct {
	userRepo *UserRepository
}

func NewAuthService(userRepo *UserRepository) *AuthService {
	return &AuthService{userRepo: userRepo}
}

func (s *AuthService) Register(username, email, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	return s.userRepo.CreateUser(username, email, string(hashedPassword))
}

func (s *AuthService) Login(email, password string) (string, error) {
	userID, hashedPassword, err := s.userRepo.FindUserByEmail(email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		return "", errors.New("invalid credentials")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
