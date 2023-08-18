package user

import "strings"

type Status int

const (
	Active Status = iota
	NotActive
)

var (
	userStatusMap = map[string]Status{
		"active":    Active,
		"notactive": NotActive,
	}
)

func (s Status) String() string {
	return [...]string{"Active", "NotActive"}[s]
}

func StatusFrom(valueStr string) (Status, bool) {
	var status, ok = userStatusMap[strings.ToLower(valueStr)]
	return status, ok
}
