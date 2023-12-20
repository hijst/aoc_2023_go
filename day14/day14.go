package main

import (
	"aoc_2023_go/util"
	"fmt"
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
	tiltNorth(grid)
	res := calculateLoad(grid)
	fmt.Printf("answer: %d\n", res)
}

func two(lines []string) {
	grid := readGrid(lines)
	finalState := findFinalState(grid)
	res := calculateLoad(finalState)
	fmt.Printf("answer: %d\n", res)
}

func readGrid(lines []string) [][]string {
	var res [][]string
	for _, line := range lines {
		res = append(res, strings.Split(line, ""))
	}
	return res
}

func findFinalState(grid [][]string) [][]string {
	cache := make(map[string]int)
	for i := 0; i <= 1000; i++ {
		gs := util.ToString(grid)
		_, exists := cache[gs]
		if exists {
			cl := i - cache[gs]
			numberOfCyclesToRun := (1000000000 - i) % cl
			for j := 0; j < numberOfCyclesToRun; j++ {
				cycle(grid)
			}
			return (grid)
		} else {
			cache[gs] = i
			cycle(grid)
		}
	}
	panic("NO CYCLE FOUND")
}

func calculateLoad(grid [][]string) int {
	var res int
	for i, row := range grid {
		for _, col := range row {
			if col == "O" {
				res += len(grid) - i
			}
		}
	}
	return res
}

func cycle(grid [][]string) {
	tiltNorth(grid)
	tiltWest(grid)
	tiltSouth(grid)
	tiltEast(grid)
}

func tiltNorth(grid [][]string) [][]string {
	for i, row := range grid {
		for j, col := range row {
			if col == "O" {
				k := i - 1
				for k >= 0 && grid[k][j] == "." {
					grid[k+1][j] = "."
					grid[k][j] = "O"
					k--
				}
			}
		}
	}
	return grid
}

func tiltWest(grid [][]string) [][]string {
	for i, row := range grid {
		for j, col := range row {
			if col == "O" {
				k := j - 1
				for k >= 0 && grid[i][k] == "." {
					grid[i][k+1] = "."
					grid[i][k] = "O"
					k--
				}
			}
		}
	}
	return grid
}

func tiltSouth(grid [][]string) [][]string {
	l := len(grid) - 1
	for i, row := range grid {
		for j := range row {
			if grid[l-i][j] == "O" {
				k := i - 1
				for k >= 0 && grid[l-k][j] == "." {
					grid[l-k-1][j] = "."
					grid[l-k][j] = "O"
					k--
				}
			}
		}
	}
	return grid
}

func tiltEast(grid [][]string) [][]string {
	l := len(grid[0]) - 1
	for i, row := range grid {
		for j := range row {
			if grid[i][l-j] == "O" {
				k := j - 1
				for k >= 0 && grid[i][l-k] == "." {
					grid[i][l-k-1] = "."
					grid[i][l-k] = "O"
					k--
				}
			}
		}
	}
	return grid
}
