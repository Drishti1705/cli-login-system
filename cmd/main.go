package main

import (
	"github.com/Drishti1705/cli-login-system/internal/cli"
	"github.com/Drishti1705/cli-login-system/internal/database"
)

func main() {

	database.Connect()
	database.Migrate()

	cli.Start()
}