package main

import (
	"aoc_2023_go/util"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	lines, _ := util.ReadFileLines("")
	one(lines)
	two(lines)
}

func one(lines []string) {
	sum := 0

	for i, line := range lines {
		split := prepare(line)
		satisfies := satisfies(split)
		if satisfies {
			sum += 1 + i
		}
	}

	fmt.Printf("the answer to question 1 is %d\n", sum)
}

func two(lines []string) {
	sum := 0

	for _, line := range lines {
		sum += findPower(line)
	}

	fmt.Printf("the answer to question 2 is %d\n", sum)
}

func prepare(line string) []string {
	draws := cutAtColon(line)
	replaced := strings.ReplaceAll(draws, ";", ",")
	replacedcommas := strings.ReplaceAll(replaced, ",", "")
	split := strings.Split(replacedcommas, " ")
	return split
}

func cutAtColon(line string) string {
	for i, char := range line {
		if char == ':' {
			return line[i+2:]
		}
	}
	panic("line did not contain :")
}

func satisfies(split []string) bool {
	amountPerColor := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	for j := range split {
		if j%2 == 0 {
			count, _ := strconv.Atoi(split[j])
			if count > amountPerColor[split[j+1]] {
				return false
			}
		}
	}
	return true
}

func findPower(line string) int {
	minPerColor := map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}
	split := prepare(line)

	for j := range split {
		if j%2 == 0 {
			count, _ := strconv.Atoi(split[j])
			if count > minPerColor[split[j+1]] {
				minPerColor[split[j+1]] = count
			}
		}
	}
	return minPerColor["red"] * minPerColor["green"] * minPerColor["blue"]
}
