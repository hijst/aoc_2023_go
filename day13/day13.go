package main

import (
	"aoc_2023_go/util"
	"fmt"
	"reflect"
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
	grids := readGrids(lines)
	res := 0
	for _, grid := range grids {
		res += solveGrid(grid, 0)
	}
	fmt.Printf("answer: %d\n", res)
}

func two(lines []string) {
	grids := readGrids(lines)
	res := 0
	for _, grid := range grids {
		res += solveSmudge(grid)
	}
	fmt.Printf("answer: %d\n", res)
}

func readGrids(lines []string) [][][]string {
	var res [][][]string
	var grid [][]string
	count := 0
	for _, line := range lines {
		if line == "" {
			count += 1
			res = append(res, grid)
			grid = nil
		} else {
			grid = append(grid, strings.Split(line, ""))
		}
	}
	res = append(res, grid)
	return res
}

func solveSmudge(grid [][]string) int {
	orig := solveGrid(grid, 0)
	derived := orig
	if derived >= 100 {
		derived = derived / 100
	}
	for i, row := range grid {
		for j := range row {
			res := trySmudge(i, j, grid, derived)
			if res > 0 && res != orig {
				return res
			}
		}
	}
	// if no new value found, it must be the same nr, but row <-> col
	if orig >= 100 {
		return orig / 100
	}
	return orig * 100
}

func trySmudge(i, j int, grid [][]string, orig int) int {
	dup := duplicate(grid)
	if dup[i][j] == "#" {
		dup[i][j] = "."
	} else {
		dup[i][j] = "#"
	}
	return solveGrid(dup, orig)
}

func solveGrid(grid [][]string, orig int) int {
	res1 := solveByRows(grid, orig) * 100
	cols := util.Transpose(grid)
	res2 := solveByRows(cols, orig)
	return max(res1, res2)
}

func solveByRows(grid [][]string, orig int) int {
	for i := 0; i < len(grid)-1; i++ {
		if reflect.DeepEqual(grid[i], grid[i+1]) {
			if isMirror(i, grid) && i+1 != orig {
				return i + 1
			}
		}
	}
	return -1
}

func isMirror(i int, grid [][]string) bool {
	closest := min(i, len(grid)-2-i)
	for j := 1; j <= closest; j++ {
		if !reflect.DeepEqual(grid[i-j], grid[i+1+j]) {
			return false
		}
	}
	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func duplicate(matrix [][]string) [][]string {
	duplicate := make([][]string, len(matrix))
	for i := range matrix {
		duplicate[i] = make([]string, len(matrix[i]))
		copy(duplicate[i], matrix[i])
	}
	return duplicate
}
