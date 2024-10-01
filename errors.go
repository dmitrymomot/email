package email

import "errors"

var (
	ErrInvalidEmail                         = errors.New("invalid email")
	ErrInvalidEmailFormat                   = errors.New("invalid email format")
	ErrInvalidEmailHost                     = errors.New("invalid email host")
	ErrInvalidIcanSuffix                    = errors.New("invalid ICANN suffix")
	ErrFailedDomainLookup                   = errors.New("failed finding MX and A records for the email's domain")
	ErrBlacklistedDomain                    = errors.New("blacklisted domain")
	ErrEmptyUsername                        = errors.New("username is empty")
	ErrUsernameTooLong                      = errors.New("username is too long")
	ErrUsernameCannotStartOrEndWithDot      = errors.New("username cannot start or end with a dot")
	ErrUsernameCannotContainConsecutiveDots = errors.New("username cannot contain consecutive dots")
	ErrInvalidUsernameChar                  = errors.New("invalid character in username")
)
