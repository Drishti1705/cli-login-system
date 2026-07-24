package tests

import "github.com/Drishti1705/cli-login-system/internal/database"

func init() {
	database.Connect("../data/login.db")
}