package main

import (
	"errors"
	"fmt"
	"regexp"
	"unicode"
)

func IsValidEmail(email string) (bool, error) {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if matches := re.MatchString(email); !matches {
		return false, fmt.Errorf("Email %s is invalid", email)
	}

	return true, nil
}

func IsValidPassword(password string) (bool, error) {
	characters := 0
	var hasNumber bool
	var hasUpper bool
	var hasLower bool
	var hasSpecial bool

	for _, c := range password {
		switch {
		case unicode.IsLower(c):
			hasLower = true
		case unicode.IsNumber(c):
			hasNumber = true
		case unicode.IsUpper(c):
			hasUpper = true
			characters++
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			hasSpecial = true
		}

		characters++
	}

	if characters < 7 {
		return false, errors.New("Invalid password, less than 7 characters")
	}

	if !hasUpper {
		return false, errors.New("Invalid password, missing uppercase letter")
	}

	if !hasLower {
		return false, errors.New("Invalid password, missing lowercase letter")
	}

	if !hasSpecial {
		return false, errors.New("Invalid password, missing special character")
	}

	if !hasNumber {
		return false, errors.New("Invalid password, missing number")
	}

	return true, nil
}
