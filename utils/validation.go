package utils

import "regexp"

func IsValidEmail(email string) bool   {
	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	regulierExpr := regexp.MustCompile(regex)

	return regulierExpr.MatchString(email)
}

func IsValidUsername(username string) bool {
	regex := `^[a-zA-Z0-9_-]{5,}$`

	regulierExpr := regexp.MustCompile(regex)

	return regulierExpr.MatchString(username)
}

