package user

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
)

// Error
var (
	ErrIllegalUserIDValue = errors.New("illegal UserID value")
)

type ID struct {
	value uuid.UUID
}

func NewUserID() *ID {
	value := uuid.New()
	return &ID{value: value}
}

func IDFrom(valueStr string) (ID, error) {
	if len(valueStr) == 0 {
		return ID{}, ErrIllegalUserIDValue
	}

	value, err := uuid.Parse(valueStr)
	if err != nil {
		return ID{}, fmt.Errorf("valueStr isn't UUID format. Cause: %q", err.Error())
	}

	return ID{value: value}, nil
}

func (u ID) Value() uuid.UUID {
	return u.value
}

func (u ID) String() string {
	return u.value.String()
}
