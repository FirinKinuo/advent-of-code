package main

import (
	"fmt"
	"log"
	"strings"
	"unicode"

	"github.com/FirinKinuo/advent-of-code"
	"github.com/FirinKinuo/advent-of-code/tools"
)

type Day struct {
	*problem.DayTemplate

	instruction string
}

func NewDay(inputType problem.InputType) (*Day, error) {
	template, err := problem.NewDayTemplate("2024", "day03_mull_it_over", inputType)
	if err != nil {
		return nil, fmt.Errorf("new day template: %w", err)
	}
	return &Day{DayTemplate: template}, nil
}

func (d *Day) PrepareInput(input string) {
	d.instruction = strings.ReplaceAll(input, "\n", "")
}

func (d *Day) findCorrectMuls(memory string, withInstructions bool) [][2]int {
	var muls [][2]int
	memoryLen := len([]rune(memory))
	mulEnabled := true

	for i := 0; i < memoryLen; i++ {
		foundMulsIndex := strings.Index(memory[i:], "mul(")
		if foundMulsIndex == -1 {
			return muls
		}
		foundMulsIndex += i + 4

		if withInstructions {
			if strings.Contains(memory[i:foundMulsIndex], "do()") {
				mulEnabled = true
			}
			if strings.Contains(memory[i:foundMulsIndex], "don't()") {
				mulEnabled = false
			}
		}

		if withInstructions && !mulEnabled {
			i = foundMulsIndex
			continue
		}

		end := foundMulsIndex + 8
		if end >= memoryLen {
			end = memoryLen
		}

		if !strings.Contains(memory[foundMulsIndex:end], ",") ||
			!strings.Contains(memory[foundMulsIndex:end], ")") {
			i = foundMulsIndex
			continue
		}

		for offset, symbol := range memory[foundMulsIndex:end] {
			if !unicode.IsDigit(symbol) && symbol == ')' {
				closeBracketIndex := foundMulsIndex + offset

				multipliers := strings.Split(memory[foundMulsIndex:closeBracketIndex], ",")
				muls = append(muls, [2]int{tools.MustAtoi(multipliers[0]), tools.MustAtoi(multipliers[1])})
				i = closeBracketIndex
				break
			}
		}
	}

	return muls
}

func (d *Day) FirstProblem() int {
	var result int

	validMuls := d.findCorrectMuls(d.instruction, false)
	for _, mul := range validMuls {
		result += mul[0] * mul[1]
	}

	return result
}

func (d *Day) SecondProblem() int {
	var result int

	validMuls := d.findCorrectMuls(d.instruction, true)
	for _, mul := range validMuls {
		result += mul[0] * mul[1]
	}

	return result
}

func main() {
	day, err := NewDay(problem.ProblemInput)
	if err != nil {
		log.Fatalf("new day: %s", err)
	}

	day.Problem.Solve(day)
}
