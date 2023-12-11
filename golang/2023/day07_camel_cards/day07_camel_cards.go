package main

import (
	"fmt"
	"github.com/FirinKinuo/advent-of-code"
	"log"
	"slices"
	"strconv"
	"strings"
)

type game struct {
	cards []int
	bet   int
	rank  int
}

const (
	HighCard = iota
	OnePair
	TwoPairs
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

func (g *game) countCards() map[int]int {
	counts := make(map[int]int)
	for _, card := range g.cards {
		if card != 1 {
			counts[card]++
		}
	}

	for k, v := range counts {
		if v == 1 {
			delete(counts, k)
		}
	}

	return counts
}

func (g *game) GetRank(withJokers bool) int {
	cardCounts := g.countCards()

	counts := make([]int, 0, len(cardCounts))
	for _, value := range cardCounts {
		counts = append(counts, value)
	}

	jokersCount := 0
	if withJokers {
		// if JJJJJ
		if slices.Equal(g.cards, []int{1, 1, 1, 1, 1}) {
			return FiveOfAKind
		}

		for _, card := range g.cards {
			if card == 1 {
				jokersCount++
			}
		}
	}

	switch len(counts) {
	case 1:
		switch counts[0] + jokersCount {
		case 2:
			return OnePair
		case 3:
			return ThreeOfAKind
		case 4:
			return FourOfAKind
		case 5:
			return FiveOfAKind
		}
	case 2:
		switch sum := counts[0] + counts[1]; sum {
		case 5:
			return FullHouse
		case 4:
			return TwoPairs
		}
	}

	return HighCard
}

type Day struct {
	*problem.DayTemplate
	games           []game
	gamesWithJokers []game
}

func NewDay(inputType problem.InputType) (*Day, error) {
	template, err := problem.NewDayTemplate("2023", "day07_camel_cards", inputType)
	if err != nil {
		return nil, fmt.Errorf("new day template: %s", err)
	}
	return &Day{DayTemplate: template}, nil
}

func (d *Day) PrepareInput(input string) {
	for _, line := range strings.Split(input, "\r\n") {
		split := strings.Split(line, " ")
		bet, _ := strconv.Atoi(split[1])
		d.games = append(d.games, game{cards: d.convertCardsToInts(split[0], false), bet: bet})
		d.gamesWithJokers = append(d.gamesWithJokers, game{cards: d.convertCardsToInts(split[0], true), bet: bet})
	}
}

func (d *Day) convertCardsToInts(cardString string, withJokers bool) []int {
	cards := make([]int, 0, len(cardString))
	for _, card := range cardString {
		switch card {
		case 'A':
			cards = append(cards, 14)
		case 'K':
			cards = append(cards, 13)
		case 'Q':
			cards = append(cards, 12)
		case 'J':
			if withJokers {
				cards = append(cards, 1)
			} else {
				cards = append(cards, 11)

			}
		case 'T':
			cards = append(cards, 10)
		default:
			convertedCard, _ := strconv.Atoi(string(card))
			cards = append(cards, convertedCard)
		}
	}

	return cards
}

func (d *Day) sortByRank(a, b game) int {
	switch {
	case a.rank < b.rank:
		return -1
	case a.rank == b.rank:
		for i := 0; i < len(a.cards); i++ {
			switch {
			case a.cards[i] < b.cards[i]:
				return -1
			case a.cards[i] > b.cards[i]:
				return 1
			}
		}
	}

	return 1
}

func (d *Day) FirstProblem() int {
	totalWin := 0
	gamesWithRanks := make([]game, 0, len(d.games))

	for _, g := range d.games {
		g.rank = g.GetRank(false)
		gamesWithRanks = append(gamesWithRanks, g)
	}

	slices.SortStableFunc(gamesWithRanks, d.sortByRank)

	for i, g := range gamesWithRanks {
		totalWin += g.bet * (i + 1)
	}

	return totalWin
}

func (d *Day) SecondProblem() int {
	totalWin := 0
	gamesWithRanks := make([]game, 0, len(d.gamesWithJokers))

	for _, g := range d.gamesWithJokers {
		g.rank = g.GetRank(true)
		gamesWithRanks = append(gamesWithRanks, g)
	}
	slices.SortStableFunc(gamesWithRanks, d.sortByRank)

	for i, g := range gamesWithRanks {
		totalWin += g.bet * (i + 1)
	}

	return totalWin
}

func main() {
	day, err := NewDay(problem.ProblemInput)
	if err != nil {
		log.Fatalf("new day: %s", err)
	}

	day.Problem.Solve(day)
}
