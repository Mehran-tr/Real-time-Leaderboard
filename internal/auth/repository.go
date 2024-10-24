package auth

import (
	"database/sql"
	"errors"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

// CreateUser inserts a new user into the database
func (r *UserRepository) CreateUser(username, email, passwordHash string) error {
	query := `INSERT INTO users (username, email, password_hash) VALUES ($1, $2, $3)`
	_, err := r.db.Exec(query, username, email, passwordHash)
	if err != nil {
		return err
	}
	return nil
}

// FindUserByEmail finds a user by email and returns the ID and hashed password
func (r *UserRepository) FindUserByEmail(email string) (int, string, error) {
	var id int
	var passwordHash string

	query := `SELECT id, password_hash FROM users WHERE email = $1`
	err := r.db.QueryRow(query, email).Scan(&id, &passwordHash)
	if err == sql.ErrNoRows {
		return 0, "", errors.New("user not found")
	}
	if err != nil {
		return 0, "", err
	}
	return id, passwordHash, nil
}
