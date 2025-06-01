package service

import "testing"

func TestUALService_Validate(t *testing.T) {
	tests := []struct {
		name        string
		ual         string
		expectError bool
	}{
		{
			name:        "valid with 3 segments",
			ual:         "did:dkg:otp/0x1234567890abcdef1234567890abcdef12345678/1234",
			expectError: false,
		},
		{
			name:        "valid with 4 segments",
			ual:         "did:dkg:otp/0x1234567890abcdef1234567890abcdef12345678/1234/5678",
			expectError: false,
		},
		{
			name:        "valid with 3 segments with chainId",
			ual:         "did:dkg:otp:2043/0x1234567890abcdef1234567890abcdef12345678/1234",
			expectError: false,
		},
		{
			name:        "valid with 4 segments with chainId",
			ual:         "did:dkg:otp:2043/0x1234567890abcdef1234567890abcdef12345678/1234/5678",
			expectError: false,
		},
		{
			name:        "invalid - missing prefix",
			ual:         "invalid:otp/0x123/1234",
			expectError: true,
		},
		{
			name:        "invalid - only 2 segments",
			ual:         "did:dkg:otp/1234",
			expectError: true,
		},
		{
			name:        "invalid - too many segments",
			ual:         "did:dkg:otp/0x123/1234/5678/extra",
			expectError: true,
		},
		{
			name:        "empty string",
			ual:         "",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := UALSvc.Validate(tt.ual)
			if tt.expectError && err == nil {
				t.Errorf("Validate(%q): expected error, got nil", tt.ual)
			}
			if !tt.expectError && err != nil {
				t.Errorf("Validate(%q): unexpected error: %v", tt.ual, err)
			}
		})
	}
}
