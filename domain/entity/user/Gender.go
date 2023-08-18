package user

import "strings"

type Gender int

const (
	Male Gender = iota
	Female
	Other
)

var (
	gendersMap = map[string]Gender{
		"male":   Male,
		"female": Female,
		"other":  Other,
	}
)

func (g Gender) String() string {
	return [...]string{"Male", "Female", "Other"}[g]
}

func GenderFrom(valueStr string) (Gender, bool) {
	gender, ok := gendersMap[strings.ToLower(valueStr)]
	return gender, ok
}
