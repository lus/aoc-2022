package main

import (
	"fmt"
	"github.com/lus/aoc-2022/internal/x"
	"strings"
)

type CPU struct {
	cycle     int
	xRegister int

	cycleListeners []func(int, int)
}

func (cpu *CPU) Initialize() {
	cpu.cycle = 0
	cpu.xRegister = 1
}

func (cpu *CPU) Execute(instruction string) {
	split := strings.Split(instruction, " ")
	switch split[0] {
	case "noop":
		cpu.advanceCycles(1)
		break
	case "addx":
		val := x.MustInt(split[1])
		cpu.advanceCycles(2)
		cpu.xRegister += val
		break
	default:
		break
	}
}

func (cpu *CPU) advanceCycles(n int) {
	for i := 0; i < n; i++ {
		cpu.cycle++
		for _, listener := range cpu.cycleListeners {
			listener(cpu.cycle, cpu.xRegister)
		}
	}
}

func (cpu *CPU) RegisterCycleListener(listener func(int, int)) {
	cpu.cycleListeners = append(cpu.cycleListeners, listener)
}

func main() {
	input := x.ReadInput(true)
	lines := strings.Split(input, "\n")

	cpu := new(CPU)
	cpu.Initialize()

	signalStrengthSum := 0
	cpu.RegisterCycleListener(func(cycle, xRegister int) {
		if (cycle-20)%40 == 0 {
			signalStrengthSum += cycle * xRegister
		}
	})

	output := ""
	cpu.RegisterCycleListener(func(cycle int, xRegister int) {
		xPos := (cycle - 1) % 40
		if intAbs(xPos-xRegister) <= 1 {
			output += "#"
		} else {
			output += "."
		}
	})

	for _, line := range lines {
		cpu.Execute(line)
	}

	fmt.Printf("The sum of the signal strengths is %d.\n", signalStrengthSum)
	fmt.Printf("The output of the screen is:\n")
	for i := 0; i < len(output)/40; i++ {
		fmt.Println(output[i*40 : i*40+40])
	}
}

func intAbs(val int) int {
	if val < 0 {
		return -val
	}
	return val
}
