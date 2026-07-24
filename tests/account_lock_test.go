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

func TestAccountLockout(t *testing.T) {

	username := fmt.Sprintf("lock_%d", time.Now().UnixNano())
	password := "Password@123"

	// Register user
	err := auth.RegisterUser(username, password)
	if err != nil {
		t.Fatalf("failed to register user: %v", err)
	}

	// Three failed login attempts
	for i := 0; i < 3; i++ {
		_, _ = auth.LoginUser(username, "WrongPassword")
	}

	// Now try the correct password
	_, err = auth.LoginUser(username, password)

	if err == nil {
		t.Fatal("expected account to be locked")
	}

	t.Logf("Account correctly locked: %v", err)
}