package main

import (
	"fmt"
	"github.com/lus/aoc-2022/internal/x"
	"strings"
)

func main() {
	input := x.ReadInput(true)
	lines := strings.Split(input, "\n")

	grid := make([][]int, len(lines))
	for i, line := range lines {
		row := make([]int, len(line))
		for j, tree := range line {
			row[j] = x.MustInt(string(tree))
		}
		grid[i] = row
	}

	visibleCount := 0
	maxScenicScore := 0
	for i, row := range grid {
		for j, tree := range row {
			visible := false
			scenicScore := 0

			if i == 0 || i == len(grid)-1 || j == 0 || j == len(row)-1 {
				visible = true
			}

			leftScore := 0
			for k := j - 1; k >= 0; k-- {
				leftScore++
				if row[k] >= tree {
					break
				}
				if k == 0 {
					visible = true
				}
			}

			rightScore := 0
			for k := j + 1; k < len(row); k++ {
				rightScore++
				if row[k] >= tree {
					break
				}
				if k == len(row)-1 {
					visible = true
				}
			}

			upScore := 0
			for k := i - 1; k >= 0; k-- {
				upScore++
				if grid[k][j] >= tree {
					break
				}
				if k == 0 {
					visible = true
				}
			}

			downScore := 0
			for k := i + 1; k < len(grid); k++ {
				downScore++
				if grid[k][j] >= tree {
					break
				}
				if k == len(grid)-1 {
					visible = true
				}
			}

			if visible {
				visibleCount++
			}

			scenicScore = leftScore * rightScore * upScore * downScore
			if scenicScore > maxScenicScore {
				maxScenicScore = scenicScore
			}
		}
	}

	fmt.Printf("The amount of visible trees is %d.\n", visibleCount)
	fmt.Printf("The highest scenic score reached in the grid is %d.\n", maxScenicScore)
}
