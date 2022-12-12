package main

import (
	"fmt"
	"github.com/lus/aoc-2022/internal/x"
	"sort"
	"strings"
)

type Monkey struct {
	InitialItems     []int
	Items            []int
	Operation        func(int) int
	TestDivisor      int
	TrueDestination  int
	FalseDestination int
}

func main() {
	input := x.ReadInput(true)

	rawMonkeys := strings.Split(input, "\n\n")
	monkeys := make([]*Monkey, len(rawMonkeys))
	for i, rawMonkey := range rawMonkeys {
		lines := strings.Split(rawMonkey, "\n")[1:]

		monkey := new(Monkey)

		for _, rawItem := range strings.Split(strings.TrimPrefix(lines[0], "  Starting items: "), ", ") {
			item := x.MustInt(rawItem)
			monkey.InitialItems = append(monkey.InitialItems, item)
			monkey.Items = append(monkey.Items, item)
		}

		rawOperation := strings.Split(strings.TrimPrefix(lines[1], "  Operation: new = old "), " ")
		switch rawOperation[0] {
		case "*":
			if rawOperation[1] == "old" {
				monkey.Operation = func(i int) int {
					return i * i
				}
			} else {
				value := x.MustInt(rawOperation[1])
				monkey.Operation = func(i int) int {
					return i * value
				}
			}
			break
		case "+":
			if rawOperation[1] == "old" {
				monkey.Operation = func(i int) int {
					return i + i
				}
			} else {
				value := x.MustInt(rawOperation[1])
				monkey.Operation = func(i int) int {
					return i + value
				}
			}
			break
		default:
			break
		}

		monkey.TestDivisor = x.MustInt(strings.TrimPrefix(lines[2], "  Test: divisible by "))
		monkey.TrueDestination = x.MustInt(strings.TrimPrefix(lines[3], "    If true: throw to monkey "))
		monkey.FalseDestination = x.MustInt(strings.TrimPrefix(lines[4], "    If false: throw to monkey "))

		monkeys[i] = monkey
	}

	result1 := calculateMonkeyInspectionCounts(monkeys, 20, true)
	for _, monkey := range monkeys {
		monkey.Items = monkey.InitialItems
	}
	result2 := calculateMonkeyInspectionCounts(monkeys, 10000, false)

	fmt.Printf("The level of monkey business after 20 rounds is %d.\n", result1[0]*result1[1])
	fmt.Printf("The level of monkey business after 10000 rounds is %d.\n", result2[0]*result2[1])
}

func calculateMonkeyInspectionCounts(monkeys []*Monkey, rounds int, relief bool) []int {
	lcd := 1
	for _, monkey := range monkeys {
		lcd *= monkey.TestDivisor
	}

	monkeyInspectionCounts := make([]int, len(monkeys))
	for i := 0; i < rounds; i++ {
		for j, monkey := range monkeys {
			for _, item := range monkey.Items {
				monkeyInspectionCounts[j]++

				worry := item
				worry = monkey.Operation(worry)
				if relief {
					worry /= 3
				}
				worry %= lcd
				if worry%monkey.TestDivisor == 0 {
					monkeys[monkey.TrueDestination].Items = append(monkeys[monkey.TrueDestination].Items, worry)
				} else {
					monkeys[monkey.FalseDestination].Items = append(monkeys[monkey.FalseDestination].Items, worry)
				}
				monkey.Items = monkey.Items[1:]
			}
		}
	}
	sort.Slice(monkeyInspectionCounts, func(i, j int) bool {
		return monkeyInspectionCounts[j] < monkeyInspectionCounts[i]
	})
	return monkeyInspectionCounts
}
