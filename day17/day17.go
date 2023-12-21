package main

import (
	"aoc_2023_go/util"
	"fmt"
	"strconv"
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

func one(lines []string) {
	grid := readGrid(lines)
	util.PrintSlice(grid)
	fmt.Printf("part one not yet implemented, number of lines: %d\n", len(lines))
}

func two(lines []string) {
	fmt.Printf("part two not yet implemented, number of lines: %d\n", len(lines))
}

func readGrid(lines []string) [][]int {
	var res [][]int
	for _, line := range lines {
		s := strings.Split(line, "")
		r := util.Map(s, atoi)
		res = append(res, r)
	}
	return res
}

func atoi(s string) int {
	r, _ := strconv.Atoi(s)
	return r
}

func iota(i int) string {
	return strconv.Itoa(i)
}
