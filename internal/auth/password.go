package auth

import "golang.org/x/crypto/bcrypt"

// HashPassword hashes a plain-text password.
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)
	return string(bytes), err
}

// CheckPassword compares a plain password with a hashed password.
func CheckPassword(password, hash string) error {
	return bcrypt.CompareHashAndPassword(
		[]byte(hash),
		[]byte(password),
	)
}