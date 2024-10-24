package auth

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

type mockUserRepository struct {
	shouldFail bool
}

func (m *mockUserRepository) CreateUser(username, email, passwordHash string) error {
	if m.shouldFail {
		return errors.New("mock create user error")
	}
	return nil
}

func (m *mockUserRepository) FindUserByEmail(email string) (int, string, error) {
	if email == "nonexistent@example.com" {
		return 0, "", errors.New("user not found")
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
	return 1, string(hashedPassword), nil
}

func TestRegisterSuccess(t *testing.T) {
	userRepo := &mockUserRepository{}
	authService := NewAuthService(userRepo)

	err := authService.Register("john_doe", "john@example.com", "password123")
	assert.NoError(t, err)
}

func TestRegisterFail(t *testing.T) {
	userRepo := &mockUserRepository{shouldFail: true}
	authService := NewAuthService(userRepo)

	err := authService.Register("john_doe", "john@example.com", "password123")
	assert.Error(t, err)
	assert.Equal(t, "mock create user error", err.Error())
}

func TestLoginSuccess(t *testing.T) {
	userRepo := &mockUserRepository{}
	authService := NewAuthService(userRepo)

	token, err := authService.Login("test@example.com", "password123")
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}

func TestLoginFailUserNotFound(t *testing.T) {
	userRepo := &mockUserRepository{}
	authService := NewAuthService(userRepo)

	token, err := authService.Login("nonexistent@example.com", "password123")
	assert.Error(t, err)
	assert.Equal(t, "invalid credentials", err.Error())
	assert.Empty(t, token)
}

func TestLoginFailWrongPassword(t *testing.T) {
	userRepo := &mockUserRepository{}
	authService := NewAuthService(userRepo)

	token, err := authService.Login("test@example.com", "wrongpassword")
	assert.Error(t, err)
	assert.Equal(t, "invalid credentials", err.Error())
	assert.Empty(t, token)
}
