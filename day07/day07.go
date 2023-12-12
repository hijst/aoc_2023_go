package main

import (
	"aoc_2023_go/util"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Hand struct {
	Cards string
	Bet   int64
	Score string
}

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

func one(lines []string) {
	hands := readHands(lines)
	solvedHands := solveHands(hands)
	orderedHands := orderHands(solvedHands)
	score := calculateScore(orderedHands)
	fmt.Printf("answer: %d\n", score)
}

func two(lines []string) {
	hands := readHands(lines)
	solvedHands := solveHandsWithJokers(hands)
	orderedHands := orderHands(solvedHands)
	score := calculateScore(orderedHands)
	fmt.Printf("answer: %d\n", score)
}

func readHands(lines []string) []Hand {
	var hands []Hand
	for _, line := range lines {
		parts := strings.Fields(line)
		bet, _ := strconv.Atoi(parts[1])
		hand := Hand{
			Cards: parts[0],
			Bet:   int64(bet),
			Score: "",
		}
		hands = append(hands, hand)
	}
	return hands
}

func solveHands(hands []Hand) []Hand {
	for i := range hands {
		hands[i].Score += findHandResult(hands[i].Cards)
		hands[i].Score += mapCardsToScore(hands[i].Cards)
	}
	return hands
}

func solveHandsWithJokers(hands []Hand) []Hand {
	for i := range hands {
		hands[i].Score += findHandResultJokers(hands[i].Cards)
		hands[i].Score += mapCardsToScoreJokers(hands[i].Cards)
	}
	return hands
}

func orderHands(hands []Hand) []Hand {
	sort.Slice(hands, func(i, j int) bool {
		return compareHexScores(hands[i].Score, hands[j].Score)
	})
	return hands
}

func findHandResult(cards string) string {
	countsMap := countCharacters(cards)
	counts := util.MapValues(countsMap)
	sort.Sort(sort.Reverse(sort.IntSlice(counts)))
	if counts[0] == 5 {
		return "F"
	}
	if counts[0] == 4 {
		return "E"
	}
	if counts[0] == 3 {
		if counts[1] == 2 {
			return "D"
		}
		return "C"
	}
	if counts[0] == 2 {
		if counts[1] == 2 {
			return "B"
		}
		return "A"
	}
	return ""
}

func findHandResultJokers(cards string) string {
	jokerCount := 0
	cardsWithoutJokers := ""
	for _, char := range cards {
		if char == 'J' {
			jokerCount += 1
		} else {
			cardsWithoutJokers = cardsWithoutJokers + string(char)
		}
	}
	if cardsWithoutJokers == "" {
		return "F"
	}
	countsMap := countCharacters(cardsWithoutJokers)
	counts := util.MapValues(countsMap)
	sort.Sort(sort.Reverse(sort.IntSlice(counts)))
	counts[0] += jokerCount
	if counts[0] == 5 {
		return "F"
	}
	if counts[0] == 4 {
		return "E"
	}
	if counts[0] == 3 {
		if counts[1] == 2 {
			return "D"
		}
		return "C"
	}
	if counts[0] == 2 {
		if counts[1] == 2 {
			return "B"
		}
		return "A"
	}
	return ""
}

func calculateScore(hands []Hand) int64 {
	sum := int64(0)
	for i, hand := range hands {
		sum += int64(i+1) * hand.Bet
	}
	return sum
}

func mapCardsToScore(cards string) string {
	replacements := map[rune]rune{
		'A': 'F',
		'K': 'E',
		'Q': 'D',
		'J': 'C',
		'T': 'B',
	}
	res := replaceCharacters(cards, replacements)
	return res
}

func mapCardsToScoreJokers(cards string) string {
	replacements := map[rune]rune{
		'A': 'F',
		'K': 'E',
		'Q': 'D',
		'J': '1',
		'T': 'B',
	}
	res := replaceCharacters(cards, replacements)
	return res
}

func replaceCharacters(input string, replacements map[rune]rune) string {
	for oldChar, newChar := range replacements {
		input = strings.Replace(input, string(oldChar), string(newChar), -1)
	}
	return input
}

func countCharacters(input string) map[rune]int {
	characterCounts := make(map[rune]int)

	for _, char := range input {
		characterCounts[char]++
	}

	return characterCounts
}

func compareHexScores(a, b string) bool {
	intA, _ := strconv.ParseInt(a, 16, 64)
	intB, _ := strconv.ParseInt(b, 16, 64)
	return intA < intB
}
