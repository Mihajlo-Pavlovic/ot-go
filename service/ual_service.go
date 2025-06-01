package service

import (
	"errors"
	"strings"
)

type UALService struct{}

// Singleton instance
var UALSvc = &UALService{}

func (s *UALService) Validate(ual string) error {
	if !strings.HasPrefix(ual, "did:dkg") {
		return errors.New("UAL must start with 'did:dkg'")
	}
	parts := strings.Split(ual, "/")
	if len(parts) != 3 && len(parts) != 4 {
		return errors.New("UAL must contain 3 or 4 segments after prefix")
	}

	return nil
}
