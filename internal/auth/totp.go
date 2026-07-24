package auth

import (
	"github.com/Drishti1705/cli-login-system/internal/repository"
	"github.com/pquerna/otp/totp"
)

// GenerateSecret creates a new TOTP secret and enables 2FA.
func GenerateSecret(username string) (string, error) {

	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "CLI Login System",
		AccountName: username,
	})

	if err != nil {
		return "", err
	}

	secret := key.Secret()

	err = repository.EnableTwoFactor(username, secret)
	if err != nil {
		return "", err
	}

	return secret, nil
}

// VerifyOTP validates the OTP entered by the user.
func VerifyOTP(secret, code string) bool {
	return totp.Validate(code, secret)
}

// Disable2FA disables 2FA for a user.
func Disable2FA(username string) error {
	return repository.DisableTwoFactor(username)
}