package types

import (
	"errors"
	"regexp"
	"strings"
)

var contractAddressPattern = regexp.MustCompile(`^0x[0-9a-fA-F]{40}$`)

type ContractAddress string

func NewContractAddress(s string) (ContractAddress, error) {
	if err := validateContractAddress(s); err != nil {
		return "", err
	}
	return ContractAddress(strings.ToLower(s)), nil
}

func validateContractAddress(s string) error {
	if len(s) != 42 {
		return errors.New("contract address must be 42 charactes long")
	}
	if !contractAddressPattern.MatchString(string(s)) {
		return errors.New("contract address is not in valid hex format")
	}
	return nil
}
