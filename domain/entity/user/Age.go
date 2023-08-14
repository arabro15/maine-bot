package user

import "errors"

var (
	ErrIllegalAgeValue = errors.New("illegal Age value")
)

type Age struct {
	value string
}

func NewAge(value string) (*Age, error) {
	if len(value) == 0 {
		return nil, ErrIllegalAgeValue
	}

	return &Age{value: value}, nil
}

func (a *Age) Value() string {
	return a.value
}
