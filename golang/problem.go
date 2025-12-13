package problem

import (
	"fmt"
	"io"
	"os"
	"time"
)

type InputType string

const (
	TestInput    InputType = "test.txt"
	Test2Input   InputType = "test2.txt" // Test2Input for some days, where we have 2 different test inputs
	ProblemInput InputType = "input.txt"
)

type Solver interface {
	Year() string
	DayName() string
	PrepareInput(input string)
	FirstProblem() int
	SecondProblem() int
}

type Problem struct {
	input string
}

func NewProblem(inputPath string) (*Problem, error) {
	inputFile, err := os.OpenFile(inputPath, os.O_RDONLY, 0)
	if err != nil {
		return nil, fmt.Errorf("open input data: %w", err)
	}

	readInput, err := io.ReadAll(inputFile)
	if err != nil {
		return nil, fmt.Errorf("read input: %w", err)
	}

	return &Problem{
		input: string(readInput),
	}, nil
}

func (p *Problem) Solve(solver Solver) {
	fmt.Printf("\nStart solving %s: %s\n=========\n", solver.Year(), solver.DayName())
	solver.PrepareInput(p.input)
	p.solveAndPrint(solver.FirstProblem, 1)
	p.solveAndPrint(solver.SecondProblem, 2)
}

func (p *Problem) solveAndPrint(solution func() int, problemNumber int) {
	fmt.Printf("Problem #%d\n", problemNumber)
	solverStopwatch := time.Now()
	result := solution()
	solvingTime := time.Now().Sub(solverStopwatch)
	fmt.Printf("Result: %d\nTime: %s\n=========\n", result, solvingTime)
}
