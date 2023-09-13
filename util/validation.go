package util

import (
	"regexp"
	"unicode"
)

func IsValidPassword(s string) bool {
	var hasNumber, hasUpperCase, hasLowercase, hasSpecial bool

	for _, c := range s {
		switch {
		case unicode.IsNumber(c):
			hasNumber = true
		case unicode.IsUpper(c):
			hasUpperCase = true
		case unicode.IsLower(c):
			hasLowercase = true
		case c == '#' || c == '|':
			return false
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			hasSpecial = true
		}
	}

	return hasNumber && hasUpperCase && hasLowercase && hasSpecial
}

func IsValidUsername(s string) bool {
	var validUsername = regexp.MustCompile(`^[a-zA-z]{4,}$`)

	return validUsername.MatchString(s)
}
