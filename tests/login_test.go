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

func TestLoginSuccess(t *testing.T) {

	username := fmt.Sprintf("login_user_%d", time.Now().UnixNano())
	password := "Password@123"

	// Register a new user
	err := auth.RegisterUser(username, password)
	if err != nil {
		t.Fatalf("failed to register user: %v", err)
	}

	// Attempt login
	user, err := auth.LoginUser(username, password)

	if err != nil {
		t.Fatalf("expected successful login, got %v", err)
	}

	if user.Username != username {
		t.Fatalf("expected username %s, got %s", username, user.Username)
	}
}

func TestLoginWrongPassword(t *testing.T) {

	username := fmt.Sprintf("wrong_login_%d", time.Now().UnixNano())
	password := "Password@123"

	err := auth.RegisterUser(username, password)
	if err != nil {
		t.Fatalf("failed to register user: %v", err)
	}

	_, err = auth.LoginUser(username, "WrongPassword")

	if err == nil {
		t.Fatal("expected login to fail with wrong password")
	}
}

func TestLoginUnknownUser(t *testing.T) {

	_, err := auth.LoginUser("user_does_not_exist", "Password@123")

	if err == nil {
		t.Fatal("expected login to fail for unknown user")
	}
}