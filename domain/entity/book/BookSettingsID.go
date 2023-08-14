package book

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
)

// Errors
var (
	ErrIllegalBookSettingsIDValue = errors.New("illegal UserID value")
)

type SettingsID struct {
	value uuid.UUID
}

func NewBookSettingsID() *SettingsID {
	value := uuid.New()
	return &SettingsID{value: value}
}

func (u *SettingsID) From(valueStr string) (*SettingsID, error) {
	if len(valueStr) == 0 {
		return nil, ErrIllegalBookSettingsIDValue
	}

	value, err := uuid.Parse(valueStr)
	if err != nil {
		return nil, fmt.Errorf("valueStr isn't UUID format. Cause: %q", err.Error())
	}

	return &SettingsID{value: value}, nil
}

func (u *SettingsID) Value() uuid.UUID {
	return u.value
}

func (u *SettingsID) String() string {
	return u.value.String()
}
