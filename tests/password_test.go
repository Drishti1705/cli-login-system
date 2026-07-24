package tests

import (
	"testing"

	"github.com/Drishti1705/cli-login-system/internal/auth"
)

func TestHashPassword(t *testing.T) {
	password := "Password@123"

	hash, err := auth.HashPassword(password)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if hash == "" {
		t.Fatal("expected hash to be generated")
	}

	if hash == password {
		t.Fatal("hash should not be equal to the original password")
	}
}

func TestCheckPasswordSuccess(t *testing.T) {
	password := "Password@123"

	hash, err := auth.HashPassword(password)
	if err != nil {
		t.Fatalf("failed to hash password: %v", err)
	}

	err = auth.CheckPassword(password, hash)

	if err != nil {
		t.Fatalf("expected password to match, got %v", err)
	}
}

func TestCheckPasswordFailure(t *testing.T) {
	password := "Password@123"

	hash, err := auth.HashPassword(password)
	if err != nil {
		t.Fatalf("failed to hash password: %v", err)
	}

	err = auth.CheckPassword("WrongPassword", hash)

	if err == nil {
		t.Fatal("expected password mismatch error")
	}
}