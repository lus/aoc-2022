package main

import (
	"fmt"
	"github.com/lus/aoc-2022/x"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input := x.ReadInput()
	lines := strings.Split(input, "\n")

	var elves []int
	cur := 0
	for _, line := range lines {
		if line == "" {
			elves = append(elves, cur)
			cur = 0
			continue
		}

		number, err := strconv.Atoi(line)
		x.PanicOnErr(err)
		cur += number
	}
	sort.Ints(elves)
	elves = elves[len(elves)-3:]

	fmt.Printf("The answer is %d.\n", elves[0]+elves[1]+elves[2])
}
