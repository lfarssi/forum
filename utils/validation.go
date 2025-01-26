package utils

import (
	"net/http"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

func IsValidEmail(email string) bool {
	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	regulierExpr := regexp.MustCompile(regex)

	return regulierExpr.MatchString(email)
}

func IsValidUsername(username string) bool {
	regex := `^[a-zA-Z0-9_-]{3,}$`

	regulierExpr := regexp.MustCompile(regex)

	return regulierExpr.MatchString(username)
}

func IsValidPassword(password string) bool {
	regex := `^[a-zA-Z0-9._%+-]{8,}$`

	regulierExpr := regexp.MustCompile(regex)

	return regulierExpr.MatchString(password)
}

func HashPassword(password string) string {
	HashPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(HashPassword)
}

func IsLoggedIn(r *http.Request) bool {
	cookie, err := r.Cookie("token")
	if err != nil || cookie.Value == "" {
		return false
	}

	return true
}
