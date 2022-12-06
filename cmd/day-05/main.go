package main

import (
	"fmt"
	"github.com/lus/aoc-2022/internal/x"
	"strings"
)

func main() {
	input := x.ReadInput(false)
	lines := strings.Split(input, "\n")

	crates := parseCrateDrawing(lines)
	crates2 := parseCrateDrawing(lines)

	for _, line := range lines {
		instruction := strings.Split(line, " ")
		if instruction[0] != "move" {
			continue
		}

		amount := x.MustInt(instruction[1])
		from := x.MustInt(instruction[3]) - 1
		to := x.MustInt(instruction[5]) - 1

		for i := 0; i < amount; i++ {
			crates[to] = append(crates[to], crates[from][len(crates[from])-1-i])
		}
		crates[from] = crates[from][:len(crates[from])-amount]

		toMove := crates2[from][len(crates2[from])-amount:]
		crates2[to] = append(crates2[to], toMove...)
		crates2[from] = crates2[from][:len(crates2[from])-amount]
	}

	result := ""
	for _, crate := range crates {
		result += crate[len(crate)-1]
	}
	fmt.Printf("The final order of the top crates will be '%s'.\n", result)

	result2 := ""
	for _, crate := range crates2 {
		result2 += crate[len(crate)-1]
	}
	fmt.Printf("The final order of the top crates when using the CrateMover 9001 will be '%s'.\n", result2)
}

func parseCrateDrawing(lines []string) [9][]string {
	var rawCrates [][]string
	for _, line := range lines {
		if strings.Replace(line, " ", "", -1)[0] == '1' {
			break
		}

		var curCrates []string
		for i := 0; i < len(line); i += 4 {
			curCrates = append(curCrates, string(line[i+1]))
		}
		rawCrates = append(rawCrates, curCrates)
	}

	var crates [9][]string
	for i := len(rawCrates) - 1; i >= 0; i-- {
		for j := 0; j < 9; j++ {
			if rawCrates[i][j] != " " {
				crates[j] = append(crates[j], rawCrates[i][j])
			}
		}
	}

	return crates
}
