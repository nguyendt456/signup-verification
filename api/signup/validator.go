package signup

import (
	"fmt"
	"regexp"
)

func ValidateUsername(username string) error {
	minLength := 5
	maxLength := 100

	if len(username) <= minLength || len(username) >= maxLength {
		return fmt.Errorf("Username must contain from %d to %d", minLength, maxLength)
	}
	isValid := regexp.MustCompile(`^[a-z0-9_]+$`).MatchString(username)
	if !isValid {
		return fmt.Errorf("Username must not contain capital or special letter !")
	}
	return nil
}
