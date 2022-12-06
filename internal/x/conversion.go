package x

import "strconv"

func MustInt(raw string) int {
	parsed, err := strconv.Atoi(raw)
	PanicOnErr(err)
	return parsed
}
