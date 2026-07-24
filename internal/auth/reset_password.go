package auth

import (
	"errors"

	"github.com/Drishti1705/cli-login-system/internal/repository"
)

func ResetPassword(username, oldPassword, newPassword string) error {

	user, err := repository.GetUserByUsername(username)
	if err != nil {
		return err
	}

	// Verify old password
	err = CheckPassword(oldPassword, user.Password)
	if err != nil {
		return errors.New("Old Password is Incorrect")
	}

	// Hash the new password
	hashedPassword, err := HashPassword(newPassword)
	if err != nil {
		return err
	}

	// Update password
	return repository.UpdatePassword(username, hashedPassword)
}