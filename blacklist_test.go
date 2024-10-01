package email_test

import (
	"reflect"
	"testing"

	"github.com/dmitrymomot/email"
)

func TestIsDomainBlacklisted(t *testing.T) {
	tests := []struct {
		name    string
		domain  string
		want    bool
		setup   func()
		cleanup func()
	}{
		{
			name:   "ExistingDomain",
			domain: "bad.com",
			want:   true,
			setup: func() {
				email.SetBlacklist("bad.com")
			},
			cleanup: func() {
				email.RemoveBlacklist("bad.com")
			},
		},
		{
			name:   "NonExistingDomain",
			domain: "good.com",
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				tt.setup()
			}
			if got := email.IsDomainBlacklisted(tt.domain); got != tt.want {
				t.Errorf("IsDomainBlacklisted() = %v, want %v", got, tt.want)
			}
			if tt.cleanup != nil {
				tt.cleanup()
			}
		})
	}
}

func TestBlacklist(t *testing.T) {
	tests := []struct {
		name    string
		want    map[string]bool
		setup   func()
		cleanup func()
	}{
		{
			name: "GetBlacklist",
			want: map[string]bool{
				"bad.com": true,
			},
			setup: func() {
				email.SetBlacklist("bad.com")
			},
			cleanup: func() {
				email.RemoveBlacklist("bad.com")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				tt.setup()
			}
			if got := email.Blacklist(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Blacklist() = %v, want %v", got, tt.want)
			}
			if tt.cleanup != nil {
				tt.cleanup()
			}
		})
	}
}

func TestSetBlacklist(t *testing.T) {
	tests := []struct {
		name    string
		domains []string
		want    map[string]bool
		cleanup func()
	}{
		{
			name:    "SetDomains",
			domains: []string{"bad.com", "worse.com"},
			want: map[string]bool{
				"bad.com":   true,
				"worse.com": true,
			},
			cleanup: func() {
				email.RemoveBlacklist("bad.com")
				email.RemoveBlacklist("worse.com")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			email.SetBlacklist(tt.domains...)
			if got := email.Blacklist(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Blacklist() = %v, want %v", got, tt.want)
			}
			if tt.cleanup != nil {
				tt.cleanup()
			}
		})
	}
}

func TestRemoveBlacklist(t *testing.T) {
	tests := []struct {
		name   string
		domain string
		setup  func()
	}{
		{
			name:   "RemoveDomain",
			domain: "bad.com",
			setup: func() {
				email.SetBlacklist("bad.com")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				tt.setup()
			}
			email.RemoveBlacklist(tt.domain)
			if email.IsDomainBlacklisted(tt.domain) {
				t.Errorf("RemoveBlacklist() = %v, want %v", true, false)
			}
		})
	}
}

func TestAddBlacklist(t *testing.T) {
	tests := []struct {
		name    string
		domain  string
		want    bool
		cleanup func()
	}{
		{
			name:   "AddDomain",
			domain: "bad.com",
			want:   true,
			cleanup: func() {
				email.RemoveBlacklist("bad.com")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			email.AddBlacklist(tt.domain)
			if got := email.IsDomainBlacklisted(tt.domain); got != tt.want {
				t.Errorf("AddBlacklist() = %v, want %v", got, tt.want)
			}
			if tt.cleanup != nil {
				tt.cleanup()
			}
		})
	}
}
