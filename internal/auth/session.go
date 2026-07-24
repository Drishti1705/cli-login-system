package auth

import "github.com/Drishti1705/cli-login-system/internal/models"

var CurrentUser *models.User

func LoginSession(user *models.User) {
	CurrentUser = user
}

func LogoutSession() {
	CurrentUser = nil
}

func IsLoggedIn() bool {
	return CurrentUser != nil
}

func GetCurrentUser() *models.User {
	return CurrentUser
}