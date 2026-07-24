package auth

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/Drishti1705/cli-login-system/internal/models"
	"github.com/Drishti1705/cli-login-system/internal/repository"
)

const (
	MaxFailedAttempts = 3
	LockDuration      = 15 * time.Minute
)

// LoginUser authenticates a user.
func LoginUser(username, password string) (*models.User, error) {

	user, err := repository.GetUserByUsername(username)
	if err != nil {
		return nil, err
	}

	// Check if account is locked
	if user.LockedUntil != nil && time.Now().Before(*user.LockedUntil) {
		return nil, errors.New("account is locked. Please try again later")
	}

	// Verify password
	err = CheckPassword(password, user.Password)
	if err != nil {

		// Increment failed attempts
		attempts := user.FailedAttempts + 1
		_ = repository.UpdateFailedAttempts(username, attempts)

		// Lock account if limit reached
		if attempts >= MaxFailedAttempts {

			lockTime := time.Now().Add(LockDuration)

			err = repository.LockAccount(username, lockTime)
			if err != nil {
				return nil, err
			}

			return nil, errors.New("account locked for 15 minutes")
		}

		return nil, errors.New("invalid password")
	}

	// Reset failed attempts
	_ = repository.ResetFailedAttempts(username)

	// Verify OTP if 2FA is enabled
	if user.TwoFactorEnabled {

		if user.TOTPSecret == nil {
			return nil, errors.New("2FA secret not configured")
		}

		reader := bufio.NewReader(os.Stdin)

		fmt.Print("Enter 6-digit OTP: ")
		code, _ := reader.ReadString('\n')
		code = strings.TrimSpace(code)

		if !VerifyOTP(*user.TOTPSecret, code) {
			return nil, errors.New("invalid OTP")
		}
	}

	// Update last login
	_ = repository.UpdateLastLogin(username)

	// Start session
	LoginSession(user)

	return user, nil
}