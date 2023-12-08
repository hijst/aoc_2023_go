package main

import (
	"aoc_2023_go/util"
	"fmt"
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
	fmt.Printf("part one not yet implemented, number of lines: %d\n", len(lines))
}

func two(lines []string) {
	fmt.Printf("part two not yet implemented, number of lines: %d\n", len(lines))
}
