package auth

import (
	"errors"

	"github.com/Drishti1705/cli-login-system/internal/database"
)

// RegisterUser registers a new user.
func RegisterUser(username, password string) error {

	// Hash the password
	hashedPassword, err := HashPassword(password)
	if err != nil {
		return err
	}

	query := `
	INSERT INTO users(username, password)
	VALUES(?, ?)
	`

	_, err = database.DB.Exec(query, username, hashedPassword)

	if err != nil {
		return errors.New("Username Already Exists")
	}

	return nil
}