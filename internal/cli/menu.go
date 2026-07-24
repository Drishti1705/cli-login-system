package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Drishti1705/cli-login-system/internal/auth"
	"github.com/Drishti1705/cli-login-system/internal/repository"
)

func Start() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("========== CLI Login System ==========")

	for {

		fmt.Println()
		fmt.Println("1. Register")
		fmt.Println("2. Login")
		fmt.Println("3. Reset Password")
		fmt.Println("4. Profile")
		fmt.Println("5. Logout")
		fmt.Println("6. Enable 2FA")
		fmt.Println("7. Disable 2FA")
		fmt.Println("8. Show Users (Debug)")
		fmt.Println("9. Exit")
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
				fmt.Println("✅ User Registered Successfully")
			}

		case "2":

			fmt.Print("Username: ")
			username, _ := reader.ReadString('\n')

			fmt.Print("Password: ")
			password, _ := reader.ReadString('\n')

			username = strings.TrimSpace(username)
			password = strings.TrimSpace(password)

			user, err := auth.LoginUser(username, password)

			if err != nil {
				fmt.Println("❌", err)
			} else {
				fmt.Printf("✅ Welcome %s!\n", user.Username)
			}

		case "3":

			fmt.Print("Username: ")
			username, _ := reader.ReadString('\n')

			fmt.Print("Old Password: ")
			oldPassword, _ := reader.ReadString('\n')

			fmt.Print("New Password: ")
			newPassword, _ := reader.ReadString('\n')

			username = strings.TrimSpace(username)
			oldPassword = strings.TrimSpace(oldPassword)
			newPassword = strings.TrimSpace(newPassword)

			err := auth.ResetPassword(username, oldPassword, newPassword)

			if err != nil {
				fmt.Println("❌", err)
			} else {
				fmt.Println("✅ Password Updated Successfully")
			}

		case "4":

			if !auth.IsLoggedIn() {
				fmt.Println("❌ Please login first.")
				break
			}

			user := auth.GetCurrentUser()

			fmt.Println()
			fmt.Println("===== PROFILE =====")
			fmt.Println("User ID                :", user.ID)
			fmt.Println("Username               :", user.Username)
			fmt.Println("2FA Enabled            :", user.TwoFactorEnabled)
			fmt.Println("Registered On          :", user.CreatedAt)

			if user.LastLogin != nil {
				fmt.Println("Last Login             :", user.LastLogin)
			} else {
				fmt.Println("Last Login             : Never")
			}

		case "5":

			if !auth.IsLoggedIn() {
				fmt.Println("❌ No active session.")
				break
			}

			auth.LogoutSession()
			fmt.Println("✅ Logged out successfully.")

		case "6":

			if !auth.IsLoggedIn() {
				fmt.Println("❌ Please login first.")
				break
			}

			user := auth.GetCurrentUser()

			secret, err := auth.GenerateSecret(user.Username)
			if err != nil {
				fmt.Println("❌", err)
				break
			}

			fmt.Println()
			fmt.Println("✅ Two Factor Authentication Enabled")
			fmt.Println()
			fmt.Println("Add this secret to Google Authenticator:")
			fmt.Println(secret)

		case "7":

			if !auth.IsLoggedIn() {
				fmt.Println("❌ Please login first.")
				break
			}

			user := auth.GetCurrentUser()

			err := auth.Disable2FA(user.Username)
			if err != nil {
				fmt.Println("❌", err)
				break
			}

			fmt.Println("✅ Two Factor Authentication Disabled")

		case "8":

			users, err := repository.GetAllUsers()

			if err != nil {
				fmt.Println("❌", err)
				break
			}

			fmt.Println()
			fmt.Printf("%-5s %-15s %-10s %-25s\n", "ID", "Username", "Attempts", "Locked Until")
			fmt.Println("----------------------------------------------------------------")

			for _, user := range users {

				lock := "NULL"

				if user.LockedUntil != nil {
					lock = user.LockedUntil.Format("2006-01-02 15:04:05")
				}

				fmt.Printf(
					"%-5d %-15s %-10d %-25s\n",
					user.ID,
					user.Username,
					user.FailedAttempts,
					lock,
				)
			}

		case "9":

			fmt.Println("Goodbye!")
			return

		default:
			fmt.Println("❌ Invalid choice. Please try again.")
		}
	}
}