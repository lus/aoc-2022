package main

import (
	"fmt"
	"github.com/lus/aoc-2022/internal/x"
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
	sort.Slice(elves, func(i, j int) bool {
		return elves[j] < elves[i]
	})

	fmt.Printf("The Elf carrying the most Calories carries a total amount of %d.\n", elves[0])
	fmt.Printf("The 3 Elves carrying the most Calories carry a total amount of %d.\n", elves[0]+elves[1]+elves[2])
}
