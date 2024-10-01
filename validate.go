package email

import (
	"errors"
	"strings"
	"unicode"

	"golang.org/x/net/publicsuffix"
)

// ValidatorFunc defines a function type for validating email addresses.
type ValidatorFunc func(username, address string) error

// ValidateEmail validates the email address locally. If you want to validate the host of
// this email address remotely call the ValidateHost method.
func ValidateEmail(emailAddr string, validators ...ValidatorFunc) error {
	username, domain, err := parseEmail(emailAddr)
	if err != nil {
		return err
	}

	// Run the custom validators
	for _, validator := range validators {
		if err := validator(username, domain); err != nil {
			return errors.Join(ErrInvalidEmail, err)
		}
	}

	return nil
}

// ValidateHost will test if the email address is actually reachable. It will first try to resolve
// the host and then start a mail transaction.
func ValidateHost(_, address string) error {
	if _, err := lookupHost(address); err != nil {
		return errors.Join(ErrInvalidEmailHost, err)
	}
	return nil
}

// ValidateIcanSuffix will test if the public suffix of the domain is managed by ICANN using
// the golang.org/x/net/publicsuffix package. If not it will return an error. Note that if this
// method returns an error it does not necessarily mean that the email address is invalid. Also the
// suffix list in the standard package is embedded and thereby not up to date.
func ValidateIcanSuffix(_, address string) error {
	d := strings.ToLower(address)
	if _, icann := publicsuffix.PublicSuffix(d); !icann {
		return ErrInvalidIcanSuffix
	}
	return nil
}

// IsAddressBlacklisted checks if the domain of the email address is blacklisted.
func IsAddressBlacklisted(_, address string) error {
	if IsDomainBlacklisted(address) {
		return ErrBlacklistedDomain
	}
	return nil
}

func ValidateUsernameFormat(username, _ string) error {
	// Check length of the username
	if len(username) == 0 {
		return ErrEmptyUsername
	}
	if len(username) > 64 {
		return ErrUsernameTooLong
	}

	// Check for leading or trailing dot
	if username[0] == '.' || username[len(username)-1] == '.' {
		return ErrUsernameCannotStartOrEndWithDot
	}

	// Check for consecutive dots
	if strings.Contains(username, "..") {
		return ErrUsernameCannotContainConsecutiveDots
	}

	// Validate allowed characters in username
	for _, r := range username {
		if !isAllowedUsernameChar(r) {
			return errors.Join(ErrInvalidUsernameChar, errors.New(string(r)))
		}
	}

	// Username is valid
	return nil
}

// isAllowedUsernameChar checks if a rune is allowed in the username
func isAllowedUsernameChar(r rune) bool {
	// Allow only letters, digits, underscores, periods, and hyphens
	if unicode.IsLetter(r) || unicode.IsDigit(r) {
		return true
	}

	switch r {
	case '_', '.', '-':
		return true
	default:
		return false
	}
}
