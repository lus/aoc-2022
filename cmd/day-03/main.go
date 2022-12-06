package main

import (
	"fmt"
	"github.com/lus/aoc-2022/internal/x"
	"strings"
	"unicode"
)

func calculatePriority(item rune) int {
	if unicode.IsLower(item) {
		return int(item) - 96
	}
	return int(item) - 38
}

func main() {
	input := x.ReadInput(true)
	lines := strings.Split(input, "\n")

	misplacedTypesPriorities := 0
	for _, backpack := range lines {
		compartment1 := backpack[:len(backpack)/2]
		compartment2 := backpack[len(backpack)/2:]

		for _, item := range []rune(compartment1) {
			if strings.Contains(compartment2, string(item)) {
				misplacedTypesPriorities += calculatePriority(item)
				break
			}
		}
	}

	badgeTypesPriorities := 0
	for i := 0; i < len(lines); i += 3 {
		backpacks := lines[i : i+3]
		for _, item := range []rune(backpacks[0]) {
			if strings.Contains(backpacks[1], string(item)) && strings.Contains(backpacks[2], string(item)) {
				badgeTypesPriorities += calculatePriority(item)
				break
			}
		}
	}

	fmt.Printf("The priorities of the misplaced item types are %d.\n", misplacedTypesPriorities)
	fmt.Printf("The priorities of the badge item types are %d.\n", badgeTypesPriorities)
}
