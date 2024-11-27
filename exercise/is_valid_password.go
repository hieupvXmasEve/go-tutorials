package exercise

import (
	"fmt"
)

// As part of improving security, Textio wants to enforce a new password policy. A valid password must meet the following criteria:
// At least 5 characters long but no more than 12 characters.
// Contains at least one uppercase letter.
// Contains at least one digit.

func isValidPassword(password string) bool {
	// Check min and Max length of password
	if len(password) > 12 || len(password) < 5 {
		return false
	}
	// Flags to track uppercase and digit requirements
	hasUppercase := false
	hasDigit := false

	for _, char := range password {
		if char >= 'A' && char <= 'Z' {
			hasUppercase = true
		}
		// Check for digit
		if char >= '0' && char <= '9' {
			hasDigit = true
		}
	}
	return hasUppercase && hasDigit
}
func main() {
	password := "Password123"
	fmt.Println(isValidPassword(password))
	password = "testting"
	fmt.Println(isValidPassword(password))
}
