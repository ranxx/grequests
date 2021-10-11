package grequests

import (
	"strings"
)

// IsStringEmpty method tells whether given string is empty or not
func IsStringEmpty(str string) bool {
	return len(strings.TrimSpace(str)) == 0
}

// IsJSONContentType json
func IsJSONContentType(str string) bool {
	if str == "application/json" {
		return true
	}
	return false
}
