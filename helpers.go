package email

import (
	"errors"
	"fmt"
	"net"
	"net/mail"
	"strings"
)

// parseEmail will parse the input and validate the email locally. If you want to validate the host of
// this email address remotely call the ValidateHost method.
func parseEmail(email string) (string, string, error) {
	addr, err := mail.ParseAddress(email)
	if err != nil {
		return "", "", ErrInvalidEmailFormat
	}
	parts := strings.Split(addr.Address, "@")
	if len(parts) != 2 {
		return "", "", ErrInvalidEmailFormat
	}
	return parts[0], parts[1], nil
}

// lookupHost first checks if any MX records are available and if not, it will check
// if A records are available as they can resolve email server hosts. An error indicates
// that non of the A or MX records are available.
func lookupHost(domain string) (string, error) {
	if mx, err := net.LookupMX(domain); err == nil {
		return mx[0].Host, nil
	}
	if ips, err := net.LookupIP(domain); err == nil {
		return ips[0].String(), nil // randomly returns IPv4 or IPv6 (when available)
	}
	return "", errors.Join(ErrFailedDomainLookup, fmt.Errorf("invalid domain: %s", domain))
}

// removeNonASCII removes non-ASCII characters from a string
func removeNonASCII(s string) string {
	result := make([]rune, 0, len(s))
	for _, r := range s {
		if r <= 127 {
			result = append(result, r)
		}
	}
	return string(result)
}
