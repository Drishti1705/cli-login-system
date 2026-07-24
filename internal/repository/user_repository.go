package repository

import (
	"database/sql"
	"errors"
	"time"

	"github.com/Drishti1705/cli-login-system/internal/database"
	"github.com/Drishti1705/cli-login-system/internal/models"
)

// CreateUser inserts a new user into the database.
func CreateUser(username, hashedPassword string) error {

	query := `
	INSERT INTO users(username, password)
	VALUES(?, ?)
	`

	_, err := database.DB.Exec(query, username, hashedPassword)
	return err
}

// GetUserByUsername fetches a user from the database.
func GetUserByUsername(username string) (*models.User, error) {

	query := `
	SELECT
		id,
		username,
		password,
		totp_secret,
		two_factor_enabled,
		failed_attempts,
		locked_until,
		created_at,
		last_login
	FROM users
	WHERE username = ?
	`

	var user models.User

	err := database.DB.QueryRow(query, username).Scan(
		&user.ID,
		&user.Username,
		&user.Password,
		&user.TOTPSecret,
		&user.TwoFactorEnabled,
		&user.FailedAttempts,
		&user.LockedUntil,
		&user.CreatedAt,
		&user.LastLogin,
	)

	if err == sql.ErrNoRows {
		return nil, errors.New("User Not Found")
	}

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func GetAllUsers() ([]models.User, error) {

	query := `
	SELECT id,
	       username,
	       failed_attempts,
	       locked_until
	FROM users
	`

	rows, err := database.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {

		var user models.User

		err := rows.Scan(
			&user.ID,
			&user.Username,
			&user.FailedAttempts,
			&user.LockedUntil,
		)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func UpdateFailedAttempts(username string, attempts int) error {

	query := `
	UPDATE users
	SET failed_attempts = ?
	WHERE username = ?
	`

	_, err := database.DB.Exec(query, attempts, username)
	return err
}

func ResetFailedAttempts(username string) error {

	query := `
	UPDATE users
	SET failed_attempts = 0,
	    locked_until = NULL
	WHERE username = ?
	`

	_, err := database.DB.Exec(query, username)
	return err
}

func LockAccount(username string, lockedUntil time.Time) error {

	query := `
	UPDATE users
	SET locked_until = ?
	WHERE username = ?
	`

	_, err := database.DB.Exec(query, lockedUntil, username)
	return err
}

func UpdatePassword(username, hashedPassword string) error {

	query := `
	UPDATE users
	SET password = ?
	WHERE username = ?
	`

	_, err := database.DB.Exec(query, hashedPassword, username)
	return err
}

func EnableTwoFactor(username, secret string) error {

	query := `
	UPDATE users
	SET totp_secret = ?,
	    two_factor_enabled = TRUE
	WHERE username = ?
	`

	_, err := database.DB.Exec(query, secret, username)
	return err
}

func DisableTwoFactor(username string) error {

	query := `
	UPDATE users
	SET totp_secret = NULL,
	    two_factor_enabled = FALSE
	WHERE username = ?
	`

	_, err := database.DB.Exec(query, username)
	return err
}

func UpdateLastLogin(username string) error {

	query := `
	UPDATE users
	SET last_login = CURRENT_TIMESTAMP
	WHERE username = ?
	`

	_, err := database.DB.Exec(query, username)
	return err
}