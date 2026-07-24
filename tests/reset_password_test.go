package tests

import (
	"fmt"
	"testing"
	"time"

	"github.com/Drishti1705/cli-login-system/internal/auth"
	"github.com/Drishti1705/cli-login-system/internal/database"
)

func init() {
	database.Connect("../data/login.db")
}

func TestResetPasswordSuccess(t *testing.T) {

	username := fmt.Sprintf("reset_%d", time.Now().UnixNano())

	oldPassword := "Password@123"
	newPassword := "NewPassword@456"

	// Register user
	err := auth.RegisterUser(username, oldPassword)
	if err != nil {
		t.Fatalf("failed to register user: %v", err)
	}

	// Reset password
	err = auth.ResetPassword(username, oldPassword, newPassword)
	if err != nil {
		t.Fatalf("password reset failed: %v", err)
	}

	// Old password should fail
	_, err = auth.LoginUser(username, oldPassword)
	if err == nil {
		t.Fatal("expected old password to fail")
	}

	// New password should work
	_, err = auth.LoginUser(username, newPassword)
	if err != nil {
		t.Fatalf("expected new password to work: %v", err)
	}
}

func TestResetPasswordWrongOldPassword(t *testing.T) {

	username := fmt.Sprintf("reset_wrong_%d", time.Now().UnixNano())

	password := "Password@123"

	err := auth.RegisterUser(username, password)
	if err != nil {
		t.Fatalf("failed to register user: %v", err)
	}

	err = auth.ResetPassword(
		username,
		"WrongPassword",
		"NewPassword@456",
	)

	if err == nil {
		t.Fatal("expected password reset to fail")
	}
}