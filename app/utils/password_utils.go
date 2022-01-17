package utils

import "golang.org/x/crypto/bcrypt"

// Hash Password ( Generate Password )
func GeneratePassword(password string) (string, error) {
	// Hash Password with Bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)

	// Check if there is error when generate password
	if err != nil {
		return "", err
	}

	// Return Hashed Password
	return string(hashedPassword), nil
}

// Compare Password from Request with from Database
func ComparePassword(passwordFromDatabase string, passwordFromRequest string) bool {
	// Match Password
	err := bcrypt.CompareHashAndPassword([]byte(passwordFromDatabase), []byte(passwordFromRequest))

	// Return when password match
	return err == nil
}
