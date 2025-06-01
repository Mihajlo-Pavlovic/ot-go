package types

import (
	"strings"
	"testing"
)

func TestNewContractAddress(t *testing.T) {
	tests := []struct {
		input       string
		expectError bool
	}{
		{"0x1234567890abcdef1234567890abcdef12345678", false}, // valid
		{"0xABCDEFabcdefABCDEFabcdefABCDEFabcdefABCD", false}, // valid uppercase
		{"0x123", true}, // too short
		{"1234567890abcdef1234567890abcdef12345678", true},     // missing "0x"
		{"0xZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZ", true},   // invalid hex chars
		{"0x1234567890abcdef1234567890abcdef1234567890", true}, // too long
		{"", true}, // empty
	}

	for _, tt := range tests {
		addr, err := NewContractAddress(tt.input)
		if tt.expectError && err == nil {
			t.Errorf("NewContractAddress(%q): expected error, got nil", tt.input)
		}
		if !tt.expectError && err != nil {
			t.Errorf("NewContractAddress(%q): unexpected error: %v", tt.input, err)
		}
		if !tt.expectError && string(addr) != strings.ToLower(tt.input) {
			t.Errorf("NewContractAddress(%q): expected address %q, got %q", tt.input, tt.input, addr)
		}
	}
}
