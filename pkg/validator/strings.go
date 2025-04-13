package validator

import (
	"net/mail"
	"strings"
	"unicode"
)

// IsEmail checks if the given string is a valid email address
// using the mail.ParseAddress function
func IsEmail(email string) bool {
	emailAddress, err := mail.ParseAddress(email)
	return err == nil && emailAddress.Address == email && strings.Contains(emailAddress.Address, ".")
}

// IsPasswordStrong checks if a given plain-text password follows streght rules.
// Minimum of 8 characters, with upper and lowercase and contains a number or a symbol.
func IsPasswordStrong(password string) bool {
	// Check minimum length
	if len(password) < 8 {
		return false
	}

	// Check for uppercase letter
	hasUppercase := false
	// Check for lowercase letter
	hasLowercase := false
	// Check for digit or symbol
	hasDigitOrSymbol := false

	for _, char := range password {
		if unicode.IsUpper(char) {
			hasUppercase = true
		} else if unicode.IsLower(char) {
			hasLowercase = true
		} else if unicode.IsDigit(char) || strings.ContainsAny(string(char), "!@#$%^&*()_+{}[]:;<>,.?/~") {
			hasDigitOrSymbol = true
		}
	}

	return hasUppercase && hasLowercase && hasDigitOrSymbol
}

// IsAnyEmpty checks if any of the provided strings are empty.
func IsAnyEmpty(values ...string) bool {
	for _, v := range values {
		if v == "" {
			return true
		}
	}
	return false
}
