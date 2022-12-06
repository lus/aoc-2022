package x

import (
	"io"
	"os"
	"strings"
)

func ReadInput(trim bool) string {
	input, err := io.ReadAll(os.Stdin)
	PanicOnErr(err)
	if trim {
		return strings.TrimSpace(string(input))
	}
	return string(input)
}
