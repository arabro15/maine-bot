package book

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
)

// Errors
var (
	ErrIllegalBookIDValue = errors.New("illegal BookID value")
)

type ID struct {
	value uuid.UUID
}

func NewBookID() *ID {
	value := uuid.New()
	return &ID{value: value}
}

func IDFrom(valueStr string) (*ID, error) {
	if len(valueStr) == 0 {
		return nil, ErrIllegalBookIDValue
	}

	value, err := uuid.Parse(valueStr)
	if err != nil {
		return nil, fmt.Errorf("valueStr isn`t UUID format. Cause: %q", err.Error())
	}

	return &ID{value: value}, nil
}

func (b ID) Value() uuid.UUID {
	return b.value
}

func (b ID) String() string {
	return b.value.String()
}
