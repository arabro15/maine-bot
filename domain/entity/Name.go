package entity

import "errors"

type Name struct {
	value string
}

// Errors
var (
	ErrIllegalNameValue = errors.New("illegal name value")
)

func NameFrom(value string) (Name, error) {
	if len(value) == 0 {
		return Name{}, ErrIllegalNameValue
	}

	return Name{value: value}, nil
}

func (n Name) Value() string {
	return n.value
}
