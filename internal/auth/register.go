package auth

import (
	"errors"
	"strings"

	"github.com/Drishti1705/cli-login-system/internal/repository"
)

func RegisterUser(username, password string) error {

	// Basic validation
	username = strings.TrimSpace(username)
	password = strings.TrimSpace(password)

	if username == "" {
		return errors.New("username cannot be empty")
	}

	if password == "" {
		return errors.New("password cannot be empty")
	}

	hashedPassword, err := HashPassword(password)
	if err != nil {
		return err
	}

	err = repository.CreateUser(username, hashedPassword)
	if err != nil {
		return errors.New("username already exists")
	}

	return nil
}