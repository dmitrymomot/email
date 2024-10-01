package email

import (
	"strings"
)

// SanitizerFunc defines a function type for sanitizing email addresses,
// returning the sanitized email or an error if any.
type SanitizerFunc func(email string) string

// SanitizeEmail sanitizes the email address to avoid scam emails
func SanitizeEmail(email string, sanitizers ...SanitizerFunc) string {
	for _, sanitizer := range sanitizers {
		email = sanitizer(email)
	}
	return email
}

// RemoveNonASCII removes non-ASCII characters from a string
func RemoveNonASCII(s string) string {
	return removeNonASCII(s)
}

// ToLower converts a string to lowercase
func ToLower(s string) string {
	return strings.ToLower(s)
}

// TrimSpace removes leading and trailing whitespace from a string
func TrimSpace(s string) string {
	return strings.TrimSpace(s)
}

// RemoveAfterPlus removes the plus part from an email address.
// For example, name+test@gmail.com will be converted to name@gmail.com
func RemoveAfterPlus(email string) string {
	emailParts := strings.Split(email, "@")
	if len(emailParts) != 2 {
		return email
	}
	username := emailParts[0]
	domain := emailParts[1]

	// Remove everything after the first plus sign
	usernameParts := strings.Split(username, "+")
	if len(usernameParts) > 1 {
		username = usernameParts[0]
	}

	return username + "@" + domain
}
