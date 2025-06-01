package types

import (
	"errors"
	"fmt"
	"strconv"
)

type KcId int

func NewKcID(s string) (KcId, error) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("invalid kcId %q: %w", s, err)
	}
	if i < 0 {
		return 0, errors.New("KcId has to be positive integer")
	}
	return KcId(i), nil
}

type KaId int

func NewKaId(s string) (KaId, error) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("invalid kaId %q: %w", s, err)
	}
	if i < 0 {
		return 0, errors.New("KaId has to be positive integer")
	}
	return KaId(i), nil
}
