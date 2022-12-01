package main

import (
	"fmt"
	"github.com/lus/aoc-2022/x"
	"strconv"
	"strings"
)

func main() {
	input := x.ReadInput()
	lines := strings.Split(input, "\n")

	cur := 0
	max := 0
	for _, line := range lines {
		if line == "" {
			if cur > max {
				max = cur
			}
			cur = 0
			continue
		}

		number, err := strconv.Atoi(line)
		x.PanicOnErr(err)
		cur += number
	}

	fmt.Printf("The answer is %d.\n", max)
}
