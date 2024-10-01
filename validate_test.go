package email

import (
	"errors"
	"testing"
)

func TestValidateEmail(t *testing.T) {
	type testEmailCase struct {
		name        string
		email       string
		validators  []ValidatorFunc
		expectError error
	}
	var testEmailCases = []testEmailCase{
		{
			name:        "Valid",
			email:       "user@example.com",
			validators:  []ValidatorFunc{ValidateHost, ValidateIcanSuffix, IsAddressBlacklisted, ValidateUsernameFormat},
			expectError: nil,
		},
		{
			name:        "InvalidHost",
			email:       "user@exampl",
			validators:  []ValidatorFunc{ValidateHost},
			expectError: ErrInvalidEmailHost,
		},
		{
			name:        "InvalidSuffix",
			email:       "user@example.comm",
			validators:  []ValidatorFunc{ValidateIcanSuffix},
			expectError: ErrInvalidIcanSuffix,
		},
		{
			name:        "Blacklisted",
			email:       "user@blacklisted.com",
			validators:  []ValidatorFunc{IsAddressBlacklisted},
			expectError: ErrBlacklistedDomain,
		},
		{
			name:        "InvalidUsername",
			email:       "user name@example.com",
			validators:  []ValidatorFunc{ValidateUsernameFormat},
			expectError: ErrInvalidEmailFormat,
		},
		{
			name:        "EmptyUsername",
			email:       "@example.com",
			validators:  []ValidatorFunc{ValidateUsernameFormat},
			expectError: ErrInvalidEmailFormat,
		},
		{
			name:        "UsernameTooLong",
			email:       "UsernameTooLongUsernameTooLongUsernameTooLongUsernameTooLong12345@example.com",
			validators:  []ValidatorFunc{ValidateUsernameFormat},
			expectError: ErrUsernameTooLong,
		},
		{
			name:        "UsernameStartWithDot",
			email:       ".name@example.com",
			validators:  []ValidatorFunc{ValidateUsernameFormat},
			expectError: ErrInvalidEmailFormat,
		},
		{
			name:        "UsernameEndWithDot",
			email:       "name.@example.com",
			validators:  []ValidatorFunc{ValidateUsernameFormat},
			expectError: ErrInvalidEmailFormat,
		},
		{
			name:        "UsernameConsecutiveDots",
			email:       "name..user@example.com",
			validators:  []ValidatorFunc{ValidateUsernameFormat},
			expectError: ErrInvalidEmailFormat,
		},
		{
			name:        "InvalidChar",
			email:       "nam$e@example.com",
			validators:  []ValidatorFunc{ValidateUsernameFormat},
			expectError: ErrInvalidUsernameChar,
		},
	}
	for _, tt := range testEmailCases {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateEmail(tt.email, tt.validators...)
			if !errors.Is(err, tt.expectError) {
				t.Errorf("expected %v, but got %v", tt.expectError, err)
			}
		})
	}
}
