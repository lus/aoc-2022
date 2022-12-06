package main

import (
	"fmt"
	"github.com/lus/aoc-2022/internal/x"
	"strings"
)

const (
	RockA     = "A"
	PaperA    = "B"
	ScissorsA = "C"

	RockB     = "X"
	PaperB    = "Y"
	ScissorsB = "Z"

	ScoreLoss = 0
	ScoreDraw = 3
	ScoreWin  = 6
)

var Scores = map[string]int{
	RockB:     1,
	PaperB:    2,
	ScissorsB: 3,
}

var Siblings = map[string]string{
	RockA:     RockB,
	PaperA:    PaperB,
	ScissorsA: ScissorsB,
}

var Counterparts = map[string]string{
	RockA:     PaperB,
	PaperA:    ScissorsB,
	ScissorsA: RockB,
}

var LosingCounterparts = map[string]string{
	RockA:     ScissorsB,
	PaperA:    RockB,
	ScissorsA: PaperB,
}

func main() {
	input := x.ReadInput(true)
	lines := strings.Split(input, "\n")

	total1 := 0
	total2 := 0
	for _, line := range lines {
		game := strings.Split(line, " ")

		total1 += Scores[game[1]]

		if Siblings[game[0]] == game[1] {
			total1 += ScoreDraw
		} else if Counterparts[game[0]] == game[1] {
			total1 += ScoreWin
		} else {
			total1 += ScoreLoss
		}

		switch game[1] {
		case "X":
			total2 += ScoreLoss
			total2 += Scores[LosingCounterparts[game[0]]]
			break
		case "Y":
			total2 += ScoreDraw
			total2 += Scores[Siblings[game[0]]]
			break
		case "Z":
			total2 += ScoreWin
			total2 += Scores[Counterparts[game[0]]]
			break
		}
	}

	fmt.Printf("The total score when following the strategy guide using the initial interpretation would be %d.\n", total1)
	fmt.Printf("The total score when following the strategy guide using the correct interpretation would be %d.\n", total2)
}
