package entity

import "errors"

type Name struct {
	value string
}

// Errors
var (
	ErrIllegalNameValue = errors.New("illegal name value")
)

func NewName(value string) (*Name, error) {
	if len(value) == 0 {
		return nil, ErrIllegalNameValue
	}

	return &Name{value: value}, nil
}

func (n *Name) Value() string {
	return n.value
}
