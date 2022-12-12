package main

import (
	"fmt"
	"github.com/lus/aoc-2022/internal/x"
	"strings"
)

type Position [2]int

func main() {
	input := x.ReadInput(false)
	lines := strings.Split(input, "\n")

	result1 := removeDuplicates(simulateRopeTail(0, lines))
	result2 := removeDuplicates(simulateRopeTail(8, lines))

	fmt.Printf("The amount of unique locations visited by the tail with 0 additional knots is %d.\n", len(result1))
	fmt.Printf("The amount of unique locations visited by the tail with 8 additional knots is %d.\n", len(result2))
}

func simulateRopeTail(additionalKnots int, moves []string) []Position {
	tailPosHistory := make([]Position, 0, len(moves))

	knotPositions := make([]Position, additionalKnots+2)
	tailPosHistory = append(tailPosHistory, knotPositions[len(knotPositions)-1])

	for _, line := range moves {
		split := strings.Split(line, " ")
		move := split[0]
		steps := x.MustInt(split[1])

		deltaX := 0
		deltaY := 0

		switch move {
		case "U":
			deltaY = 1
			break
		case "D":
			deltaY = -1
			break
		case "R":
			deltaX = 1
			break
		case "L":
			deltaX = -1
			break
		default:
			break
		}

		for i := 0; i < steps; i++ {
			knotPositions[0] = Position{knotPositions[0][0] + deltaX, knotPositions[0][1] + deltaY}
			for j := 0; j < additionalKnots+1; j++ {
				knotPositions[j+1] = simulateKnot(knotPositions[j], knotPositions[j+1])
			}
			tailPosHistory = append(tailPosHistory, knotPositions[len(knotPositions)-1])
		}
	}
	return tailPosHistory
}

func simulateKnot(parent, cur Position) (result Position) {
	parentX := parent[0]
	parentY := parent[1]
	curX := cur[0]
	curY := cur[1]

	result = Position{curX, curY}

	if intAbs(parentX-curX) <= 1 && intAbs(parentY-curY) <= 1 {
		return
	} else if parentX == curX {
		if parentY > curY {
			result[1] = parentY - 1
		} else {
			result[1] = parentY + 1
		}
	} else if parentY == curY {
		if parentX > curX {
			result[0] = parentX - 1
		} else {
			result[0] = parentX + 1
		}
	} else {
		if parentX > curX {
			if parentY > curY {
				result[0]++
				result[1]++
			} else {
				result[0]++
				result[1]--
			}
		} else {
			if parentY > curY {
				result[0]--
				result[1]++
			} else {
				result[0]--
				result[1]--
			}
		}
	}
	return
}

func intAbs(val int) int {
	if val < 0 {
		return -val
	}
	return val
}

func removeDuplicates[T comparable](raw []T) []T {
	processed := make(map[T]struct{})
	clean := make([]T, 0, len(raw))
	for _, elem := range raw {
		if _, dup := processed[elem]; dup {
			continue
		}
		clean = append(clean, elem)
		processed[elem] = struct{}{}
	}
	return clean
}
