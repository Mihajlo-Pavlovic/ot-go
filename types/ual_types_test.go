package types

import (
	"testing"
)

func TestNewKcID(t *testing.T) {
	tests := []struct {
		input       string
		expectError bool
		expectedVal KcId
	}{
		{"42", false, 42},
		{"0", false, 0},
		{"-1", true, 0},
		{"abc", true, 0},
		{"", true, 0},
	}

	for _, tt := range tests {
		got, err := NewKcID(tt.input)
		if (err != nil) != tt.expectError {
			t.Errorf("NewKcID(%q): expected error: %v, got: %v", tt.input, tt.expectError, err)
		}
		if err == nil && got != tt.expectedVal {
			t.Errorf("NewKcID(%q): expected value: %v, got: %v", tt.input, tt.expectedVal, got)
		}
	}
}

func TestNewKaID(t *testing.T) {
	tests := []struct {
		input       string
		expectError bool
		expectedVal KaId
	}{
		{"99", false, 99},
		{"0", false, 0},
		{"-10", true, 0},
		{"xyz", true, 0},
		{"", true, 0},
	}

	for _, tt := range tests {
		got, err := NewKaId(tt.input)
		if (err != nil) != tt.expectError {
			t.Errorf("NewKaID(%q): expected error: %v, got: %v", tt.input, tt.expectError, err)
		}
		if err == nil && got != tt.expectedVal {
			t.Errorf("NewKaID(%q): expected value: %v, got: %v", tt.input, tt.expectedVal, got)
		}
	}
}
