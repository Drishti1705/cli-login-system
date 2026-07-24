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

func TestSessionLogin(t *testing.T) {

	username := fmt.Sprintf("session_%d", time.Now().UnixNano())
	password := "Password@123"

	err := auth.RegisterUser(username, password)
	if err != nil {
		t.Fatalf("registration failed: %v", err)
	}

	_, err = auth.LoginUser(username, password)
	if err != nil {
		t.Fatalf("login failed: %v", err)
	}

	if !auth.IsLoggedIn() {
		t.Fatal("expected user to be logged in")
	}

	user := auth.GetCurrentUser()

	if user == nil {
		t.Fatal("expected current user, got nil")
	}

	if user.Username != username {
		t.Fatalf("expected %s, got %s", username, user.Username)
	}
}

func TestSessionLogout(t *testing.T) {

	auth.LogoutSession()

	if auth.IsLoggedIn() {
		t.Fatal("expected no active session")
	}

	if auth.GetCurrentUser() != nil {
		t.Fatal("expected current user to be nil")
	}
}