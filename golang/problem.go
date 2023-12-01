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
	ProblemInput InputType = "input.txt"
)

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

func (p *Problem) Solve(solvers []func(input string) int) {
	fmt.Println("Start solving problems\n=========")
	solvingStopwatch := time.Now()
	for i, solver := range solvers {
		fmt.Printf("Problem #%d\n", i+1)
		solverStopwatch := time.Now()

		result := solver(p.input)

		solvingTime := time.Now().Sub(solverStopwatch)

		fmt.Printf("Result: %d\nTime: %s\n=========\n", result, solvingTime)
	}

	solvingTime := time.Now().Sub(solvingStopwatch)
	fmt.Printf("Solved, total time: %s\n", solvingTime)
}
