package main

import (
	"fmt"
	"github.com/FirinKinuo/advent-of-code"
	"log"
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

type Day struct {
	*problem.DayTemplate
	calibrations                    []string
	calibrationsWithReplacedLetters []string
}

func NewDay(inputType problem.InputType) (*Day, error) {
	template, err := problem.NewDayTemplate("2023", "day01_trebuchet", inputType)
	if err != nil {
		return nil, fmt.Errorf("new day template: %s", err)
	}
	return &Day{DayTemplate: template}, nil
}

func (d *Day) PrepareInput(input string) {
	d.calibrations = strings.Split(input, "\r\n")
	d.calibrationsWithReplacedLetters = strings.Split(d.replaceLettersByDigits(input), "\r\n")
}

func (d *Day) FirstProblem() int {
	calibrationSum := 0

	for _, calibration := range d.calibrations {
		calibrationSum += d.findCalibrationValues(calibration)
	}

	return calibrationSum
}

func (d *Day) SecondProblem() int {
	calibrationSum := 0

	for _, calibration := range d.calibrationsWithReplacedLetters {
		calibrationSum += d.findCalibrationValues(calibration)
	}

	return calibrationSum
}

func (d *Day) findCalibrationValues(calibration string) int {
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

func (d *Day) replaceLettersByDigits(data string) string {
	for letter, replaced := range replacingLetters {
		data = strings.ReplaceAll(data, letter, replaced)
	}

	return data
}

func main() {
	day, err := NewDay(problem.ProblemInput)
	if err != nil {
		log.Fatalf("new day: %s", err)
	}

	day.Problem.Solve(day)
}
