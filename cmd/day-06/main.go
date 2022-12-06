package main

import (
	"fmt"
	"github.com/lus/aoc-2022/internal/x"
	"strings"
)

func main() {
	input := x.ReadInput(true)

	packetMarkerPosition := -1
	for i := 3; i < len(input); i++ {
		potentialMarker := input[i-3 : i+1]
		seen := ""
		for _, char := range potentialMarker {
			if strings.ContainsRune(seen, char) {
				break
			}
			seen += string(char)
		}
		if len(seen) == 4 {
			packetMarkerPosition = i + 1
			break
		}
	}

	messageMarkerPosition := -1
	for i := 13; i < len(input); i++ {
		potentialMarker := input[i-13 : i+1]
		seen := ""
		for _, char := range potentialMarker {
			if strings.ContainsRune(seen, char) {
				break
			}
			seen += string(char)
		}
		if len(seen) == 14 {
			messageMarkerPosition = i + 1
			break
		}
	}

	fmt.Printf("The first packet marker appears at position %d.\n", packetMarkerPosition)
	fmt.Printf("The first message marker appears at position %d.\n", messageMarkerPosition)
}
