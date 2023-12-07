package main

import (
	"aoc_2023_go/util"
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	lines, _ := util.ReadFileLines("")
	one(lines)
	two(lines)
}

func one(lines []string) {
	sum := 0
	grid := linesToGrid(lines)

	for i := 0; i < len(grid[0]); i++ {
		for j := 0; j < len(grid); {
			if isNumeric(grid[i][j]) {
				numericString := findNumberAsString(j, grid[i])
				number, _ := strconv.Atoi(numericString)
				fmt.Printf("number is: %d\n", number)
				sum += number
				j += len(numericString)
			} else {
				j++
			}
		}
	}
	fmt.Printf("answer to part 1: %d\n", sum)
}

func two(lines []string) {
	fmt.Printf("not yet implemented, number of lines: %d\n", len(lines))
}

func linesToGrid(lines []string) [][]rune {
	var res [][]rune
	for _, line := range lines {
		res = append(res, []rune(line))
	}
	return res
}

func isNumeric(character rune) bool {
	return regexp.MustCompile(`\d`).MatchString(string(character))
}

func findNumberAsString(col int, row []rune) string {
	var res []rune
	for k := col; k < len(row); k++ {
		if isNumeric(row[k]) {
			res = append(res, row[k])
		} else {
			break
		}
	}
	return string(res)
}
