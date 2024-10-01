package email

import (
	"testing"
)

func TestSanitizeEmail(t *testing.T) {
	tests := []struct {
		name       string
		email      string
		sanitizers []SanitizerFunc
		expected   string
	}{
		{
			name:       "Lowercase Email",
			email:      "EXAMPLE@EXAMPLE.COM",
			sanitizers: []SanitizerFunc{ToLower},
			expected:   "example@example.com",
		},
		{
			name:       "Trim Spaces",
			email:      "   example@example.com   ",
			sanitizers: []SanitizerFunc{TrimSpace},
			expected:   "example@example.com",
		},
		{
			name:       "Remove After Plus",
			email:      "example+spam@example.com",
			sanitizers: []SanitizerFunc{RemoveAfterPlus},
			expected:   "example@example.com",
		},
		{
			name:       "Remove Non ASCII",
			email:      "ëxamplë@example.com",
			sanitizers: []SanitizerFunc{RemoveNonASCII},
			expected:   "xampl@example.com",
		},
		{
			name:       "Multiple Sanitizers",
			email:      "   ëxamplë+spam@EXAmplë.COM   ",
			sanitizers: []SanitizerFunc{TrimSpace, RemoveNonASCII, RemoveAfterPlus, ToLower},
			expected:   "xampl@exampl.com",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SanitizeEmail(tt.email, tt.sanitizers...)
			if result != tt.expected {
				t.Errorf("expected %v, but got %v", tt.expected, result)
			}
		})
	}
}
