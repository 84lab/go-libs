package password

import (
	"errors"
	"unicode"
)

var (
	ErrInvalidPasswordLength    = errors.New("invalid password: length must not be less than 8")
	ErrInvalidPasswordNoLetters = errors.New("invalid password: at least 1 letter required")
	ErrInvalidPasswordNoNumbers = errors.New("invalid password: at least 1 number required")
	ErrInvalidPasswordNoSpecial = errors.New("invalid password: at least 1 special character required")
)

func ValidatePassword(s string) error {
	var (
		hasLetter  = false
		hasMinLen  = false
		hasNumber  = false
		hasSpecial = false
	)

	if len(s) >= 8 {
		hasMinLen = true
	}

	for _, char := range s {
		switch {
		case unicode.IsLetter(char):
			hasLetter = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	if !hasMinLen {
		return ErrInvalidPasswordLength
	}

	if !hasLetter {
		return ErrInvalidPasswordNoLetters
	}

	if !hasNumber {
		return ErrInvalidPasswordNoNumbers
	}

	if !hasSpecial {
		return ErrInvalidPasswordNoSpecial
	}

	return nil
}
