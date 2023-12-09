package main

import (
	"aoc_2023_go/util"
	"fmt"
	"regexp"
	"slices"
	"strings"
	"time"
)

func main() {
	lines, _ := util.ReadFileLines("")
	fmt.Println("______part_1______")
	start := time.Now()
	one(lines)
	fmt.Printf("took %s\n", time.Since(start))

	fmt.Println("______part_2______")
	start = time.Now()
	two(lines)
	fmt.Printf("took %s\n", time.Since(start))
}

type Card struct {
	WinningNumbers []string
	DrawnNumbers   []string
	count          int
}

func one(lines []string) {
	sum := 0
	cards := getCards(lines)
	for _, card := range cards {
		sum += findScore(card)
	}
	fmt.Printf("answer: %d\n", sum)
}

func two(lines []string) {
	cards := getCards(lines)
	result := resolveCards(cards)
	fmt.Printf("answer: %d\n", result)
}

func findScore(card Card) int {
	sum := -1
	for _, draw := range card.DrawnNumbers {
		if slices.Contains(card.WinningNumbers, draw) {
			sum += 1
		}
	}
	if sum >= 0 {
		return util.IntPow(2, sum)
	}
	return 0
}

func getCards(lines []string) []Card {
	var cards []Card
	for _, game := range lines {
		winningNumbers := regexp.MustCompile(`:.*\|`).FindString(game)
		winningNumbers = winningNumbers[1 : len(winningNumbers)-1]
		drawnNumbers := regexp.MustCompile(`\|.*`).FindString(game)[1:]
		source := strings.Fields(winningNumbers)
		target := strings.Fields(drawnNumbers)
		cards = append(cards, Card{source, target, 1})
	}
	return cards
}

func resolveCards(cards []Card) int {
	for nr, card := range cards {
		winningDigitCount := countWinningDigits(card)
		for i := 1; i <= winningDigitCount; i++ {
			cards[nr+i].count += card.count
		}
	}
	count := 0
	for _, card := range cards {
		count += card.count
	}
	return count
}

func countWinningDigits(card Card) int {
	res := 0
	for _, draw := range card.DrawnNumbers {
		if slices.Contains(card.WinningNumbers, draw) {
			res += 1
		}
	}
	return res
}
