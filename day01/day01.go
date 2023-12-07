package main

import (
	"aoc_2023_go/util"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	lines, _ := util.ReadFileLinesFromSubdirectory("", "input.txt")
	start := time.Now()

	sum := 0
	sum2 := 0
	digitMap := map[string]string{
		"1":     "1",
		"2":     "2",
		"3":     "3",
		"4":     "4",
		"5":     "5",
		"6":     "6",
		"7":     "7",
		"8":     "8",
		"9":     "9",
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	for _, line := range lines {
		first := findFirstDigit(line)
		last := findLastDigit(line)
		combined := first + last
		digit, err := strconv.Atoi(combined)
		if err != nil {
			panic("tried to convert string that was not convertible to an integer")
		}
		sum += digit
	}
	fmt.Printf("The answer to question 1 is %d\n", sum)
	count := 0
	for _, line := range lines {
		first := findFirstDigitSpelled(line, digitMap)
		last := findLastDigitSpelled(line, digitMap)
		combined := first + last
		count += 1
		digit, err := strconv.Atoi(combined)
		if err != nil {
			panic("tried to convert string that was not convertible to an integer")
		}
		sum2 += digit
	}
	fmt.Printf("The answer to question 2 is %d\n", sum2)

	elapsed := time.Since(start)
	fmt.Printf("Time taken: %s\n", elapsed)
}

func findFirstDigit(line string) string {
	chars := strings.Split(line, "")
	for _, char := range chars {
		if _, err := strconv.Atoi(char); err == nil {
			return char
		}
	}
	panic("not found")
}

func findLastDigit(line string) string {
	chars := strings.Split(line, "")
	for i := range chars {
		if _, err := strconv.Atoi(chars[len(chars)-1-i]); err == nil {
			return chars[len(chars)-1-i]
		}
	}
	panic("not found")
}

func findFirstDigitSpelled(line string, digitMap map[string]string) string {
	minIndex := len(line) + 1
	var first string
	var keys []string
	for key := range digitMap {
		keys = append(keys, key)
	}

	for _, key := range keys {
		index := strings.Index(line, key)
		if index != -1 && index < minIndex {
			minIndex = index
			first = key
		}
	}
	return digitMap[first]
}

func findLastDigitSpelled(line string, digitMap map[string]string) string {
	maxIndex := -1
	var last string
	var keys []string
	for key := range digitMap {
		keys = append(keys, key)
	}

	for _, key := range keys {
		index := strings.LastIndex(line, key)
		if index > maxIndex {
			maxIndex = index
			last = key
		}
	}
	return digitMap[last]
}
