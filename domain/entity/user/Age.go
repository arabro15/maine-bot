package user

import "errors"

var (
	ErrIllegalAgeValue = errors.New("illegal Age value")
)

type Age struct {
	value string
}

func AgeFrom(value string) (Age, error) {
	if len(value) == 0 {
		return Age{}, ErrIllegalAgeValue
	}

	return Age{value: value}, nil
}

func (a Age) Value() string {
	return a.value
}
