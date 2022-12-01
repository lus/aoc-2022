package x

import (
	"io"
	"os"
	"strings"
)

func ReadInput() string {
	input, err := io.ReadAll(os.Stdin)
	PanicOnErr(err)
	return strings.TrimSpace(string(input))
}
