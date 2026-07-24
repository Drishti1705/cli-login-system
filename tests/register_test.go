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

func TestRegisterUserSuccess(t *testing.T) {

	username := fmt.Sprintf("user_%d", time.Now().UnixNano())
	password := "Password@123"

	err := auth.RegisterUser(username, password)

	if err != nil {
		t.Fatalf("expected successful registration, got %v", err)
	}
}

func TestRegisterDuplicateUser(t *testing.T) {

	username := fmt.Sprintf("duplicate_%d", time.Now().UnixNano())
	password := "Password@123"

	err := auth.RegisterUser(username, password)
	if err != nil {
		t.Fatalf("first registration failed: %v", err)
	}

	err = auth.RegisterUser(username, password)

	if err == nil {
		t.Fatal("expected duplicate username error")
	}
}