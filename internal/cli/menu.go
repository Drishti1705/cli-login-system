package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Drishti1705/cli-login-system/internal/auth"
)

func Start() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("========== CLI Login System ==========")

	for {

		fmt.Println()
		fmt.Println("1. Register")
		fmt.Println("2. Exit")
		fmt.Print("Choose option: ")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {

		case "1":

			fmt.Print("Username: ")
			username, _ := reader.ReadString('\n')

			fmt.Print("Password: ")
			password, _ := reader.ReadString('\n')

			username = strings.TrimSpace(username)
			password = strings.TrimSpace(password)

			err := auth.RegisterUser(username, password)

			if err != nil {
				fmt.Println("❌", err)
			} else {
				fmt.Println("User Registered Successfully")
			}

		case "2":
			fmt.Println("Goodbye!")
			return

		default:
			fmt.Println("Invalid Choice")
		}
	}
}