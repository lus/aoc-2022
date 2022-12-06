package main

import (
	"fmt"
	"github.com/lus/aoc-2022/internal/x"
	"strings"
)

func main() {
	input := x.ReadInput(true)
	lines := strings.Split(input, "\n")

	fullyContainingPairs := 0
	overlappingPairs := 0
	for _, line := range lines {
		pair := strings.Split(line, ",")

		sections := [2][2]int{
			{
				x.MustInt(strings.Split(pair[0], "-")[0]),
				x.MustInt(strings.Split(pair[0], "-")[1]),
			},
			{
				x.MustInt(strings.Split(pair[1], "-")[0]),
				x.MustInt(strings.Split(pair[1], "-")[1]),
			},
		}

		if (sections[0][0] >= sections[1][0] && sections[0][1] <= sections[1][1]) ||
			(sections[1][0] >= sections[0][0] && sections[1][1] <= sections[0][1]) {
			fullyContainingPairs++
		}

		if (sections[1][0] >= sections[0][0] && sections[1][0] <= sections[0][1]) ||
			(sections[0][0] >= sections[1][0] && sections[0][0] <= sections[1][1]) {
			overlappingPairs++
		}
	}

	fmt.Printf("The amount of pairs where one range fully contains the other is %d.\n", fullyContainingPairs)
	fmt.Printf("The amount of pairs where the ranges overlap at all is %d.\n", overlappingPairs)
}
