package book

import "strings"

type Status int

const (
	Read Status = iota
	Finish
	Archived
)

var (
	bookStatusMap = map[string]Status{
		"read":     Read,
		"finish":   Finish,
		"archived": Archived,
	}
)

func StatusFrom(valueStr string) (Status, bool) {
	status, ok := bookStatusMap[strings.ToLower(valueStr)]
	return status, ok
}
