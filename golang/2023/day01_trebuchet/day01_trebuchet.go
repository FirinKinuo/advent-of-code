package main

import (
	"fmt"
	"github.com/FirinKinuo/advent-of-code"
	"log"
	"path"
	"strconv"
	"strings"
)

var replacingLetters = map[string]string{
	"one":   "o1e",
	"two":   "t2o",
	"three": "t3e",
	"four":  "f4r",
	"five":  "f5e",
	"six":   "s6x",
	"seven": "s7n",
	"eight": "e8t",
	"nine":  "n9e",
}

type Day01 struct {
	problem      *problem.Problem
	calibrations []string
}

func NewDay(inputType problem.InputType) (*Day01, error) {
	p, err := problem.NewProblem(path.Join("2023", "day01_trebuchet", string(inputType)))
	if err != nil {
		return nil, fmt.Errorf("new problem init: %s", err)
	}

	return &Day01{problem: p}, nil
}

func (d *Day01) firstProblem(input string) int {
	d.calibrations = strings.Split(input, "\r\n")

	calibrationSum := 0

	for _, calibration := range d.calibrations {
		calibrationSum += d.findCalibrationValues(calibration)
	}

	return calibrationSum
}

func (d *Day01) secondProblem(input string) int {
	d.calibrations = strings.Split(d.replaceLettersByDigits(input), "\r\n")

	calibrationSum := 0

	for _, calibration := range d.calibrations {
		calibrationSum += d.findCalibrationValues(calibration)
	}

	return calibrationSum
}

func (d *Day01) findCalibrationValues(calibration string) int {
	calibrationValue := ""
	var lastValue rune
	for _, char := range calibration {
		if char >= '0' && char <= '9' {
			lastValue = char
			if calibrationValue == "" {
				calibrationValue = string(char)
			}
		}
	}

	convertedValue, _ := strconv.Atoi(calibrationValue + string(lastValue))

	return convertedValue
}

func (d *Day01) replaceLettersByDigits(data string) string {
	for letter, replaced := range replacingLetters {
		data = strings.ReplaceAll(data, letter, replaced)
	}

	return data
}

func (d *Day01) SolveProblems() {
	solvers := []func(input string) int{
		d.firstProblem,
		d.secondProblem,
	}

	d.problem.Solve(solvers)
}

func main() {
	day, err := NewDay(problem.ProblemInput)
	if err != nil {
		log.Fatalf("new day: %s", err)
	}

	day.SolveProblems()
}
