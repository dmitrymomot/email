package email

import (
	"sync"
)

// List of known disposable or scam email domains
var blacklistedDomains = map[string]bool{
	"mailinator.com":         true,
	"tempmail.com":           true,
	"10minutemail.com":       true,
	"guerrillamail.com":      true,
	"sharklasers.com":        true,
	"guerrillamailblock.com": true,
	"blacklisted.com":        true, // Used for testing
	// Add more known scam or disposable domains as needed
}

var blacklistMutex sync.RWMutex

// IsDomainBlacklisted checks if the domain of the email address is blacklisted.
func IsDomainBlacklisted(domain string) bool {
	blacklistMutex.RLock()
	defer blacklistMutex.RUnlock()
	return blacklistedDomains[domain]
}

// Blacklist represents a list of blacklisted domains
func Blacklist() map[string]bool {
	blacklistMutex.RLock()
	defer blacklistMutex.RUnlock()
	return blacklistedDomains
}

// SetBlacklist sets the list of blacklisted domains
func SetBlacklist(domains ...string) {
	blacklistMutex.Lock()
	defer blacklistMutex.Unlock()
	blacklistedDomains = make(map[string]bool, len(domains))
	for _, domain := range domains {
		blacklistedDomains[domain] = true
	}
}

// RemoveBlacklist removes a domain from the blacklist
func RemoveBlacklist(domain string) {
	blacklistMutex.Lock()
	defer blacklistMutex.Unlock()
	delete(blacklistedDomains, domain)
}

// AddBlacklist adds a domain to the blacklist
func AddBlacklist(domain string) {
	blacklistMutex.Lock()
	defer blacklistMutex.Unlock()
	blacklistedDomains[domain] = true
}
