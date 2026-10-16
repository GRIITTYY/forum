package validator

import (
	"net/mail"
	"strings"
	"unicode/utf8"
)

func IsEmpty(value string) bool {
	return strings.TrimSpace(value) == ""
}

func IsEmptyStruct(data RegValidationResult) bool {
	var empty = RegValidationErrors{}
	return data.Errors == empty
}

func ContainsSpaces(value string) bool {
	return strings.Contains(value, " ")
}

func Matches(value1, value2 string) bool {
	return value1 == value2
}

func HasMinLength(value string, length int) bool {
	return utf8.RuneCountInString(value) >= length
}

func IsValidEmail(value string) bool {
	if !(utf8.RuneCountInString(value) >= 6 && utf8.RuneCountInString(value) <= 60) {
		return false
	}

	addr, err := mail.ParseAddress(value)
	if err != nil {
		return false
	}

	if addr.Address != value {
		return false
	}

	parts := strings.Split(addr.Address, "@")
	if len(parts) != 2 {
		return false
	}

	secondPart := parts[1]
	if strings.HasSuffix(secondPart, ".") {
		return false
	}

	secondParts := strings.Split(secondPart, ".")
	if len(secondParts) < 2 {
		return false
	}
	return true
}
