package main

import (
	"fmt"
	"github.com/FirinKinuo/advent-of-code"
	"log"
	"slices"
	"strconv"
	"strings"
)

type Day struct {
	*problem.DayTemplate
	cards [][][]int
}

func NewDay(inputType problem.InputType) (*Day, error) {
	template, err := problem.NewDayTemplate("2023", "day04_scratchcards", inputType)
	if err != nil {
		return nil, fmt.Errorf("new day template: %s", err)
	}
	return &Day{DayTemplate: template}, nil
}

func (d *Day) FirstProblem() int {
	points := 0

	for _, card := range d.cards {
		firstWin := false
		cardPoints := 0
		for _, numberToCheck := range card[0] {
			if slices.Contains(card[1], numberToCheck) {
				if !firstWin {
					cardPoints = 1
					firstWin = true
				} else {
					cardPoints *= 2
				}
			}
		}

		points += cardPoints
	}

	return points
}

func (d *Day) SecondProblem() int {
	cardsInHand := make(map[int]int)
	totalCards := 0

	for i, card := range d.cards {
		cardsInHand[i+1]++
		winCount := 1
		for _, numberToCheck := range card[0] {
			if slices.Contains(card[1], numberToCheck) {
				cardsInHand[i+1+winCount] += cardsInHand[i+1]
				winCount++
			}
		}

	}

	for _, cardCount := range cardsInHand {
		totalCards += cardCount
	}

	return totalCards
}

func (d *Day) PrepareInput(input string) {
	inputCards := strings.Split(input, "\r\n")
	d.cards = make([][][]int, len(inputCards))

	for i, card := range inputCards {
		cardSides := strings.Split(card, "|")
		d.cards[i] = make([][]int, len(cardSides))

		for j, side := range cardSides {
			d.cards[i][j] = make([]int, 0, len(side))
			side = strings.TrimPrefix(side, fmt.Sprintf("Card %d:", i+1))

			for _, value := range strings.Split(side, " ") {
				parsedValue, err := strconv.Atoi(value)
				if err != nil {
					continue
				}

				d.cards[i][j] = append(d.cards[i][j], parsedValue)
			}
		}
	}
}

func main() {
	day, err := NewDay(problem.ProblemInput)
	if err != nil {
		log.Fatalf("new day: %s", err)
	}
	day.Problem.Solve(day)
}
