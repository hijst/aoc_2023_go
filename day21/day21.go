package main

import (
	"aoc_2023_go/util"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Coord struct {
	row int
	col int
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
	grid := readGrid(lines)
	startCoord := findStartCoord(grid)
	grid[startCoord.row][startCoord.col] = "."
	res := findReachablePlots(grid, 64, startCoord)
	fmt.Printf("answer: %d\n", res)
}

func two(lines []string) {
	fmt.Printf("part two not yet implemented, number of lines: %d\n", len(lines))
}

func readGrid(lines []string) [][]string {
	var res [][]string
	for _, line := range lines {
		res = append(res, strings.Split(line, ""))
	}
	return res
}

func findReachablePlots(grid [][]string, steps int, start Coord) int {
	prevReachable := []Coord{{row: start.row, col: start.col}}
	for i := 0; i < steps; i++ {
		reachable := make(map[string]Coord)

		for _, c := range prevReachable {
			neighbours := getNeighbours(c, grid)
			for _, n := range neighbours {
				reachable[toString(n)] = n
			}
		}
		prevReachable = util.MapValues(reachable)
	}
	return len(prevReachable)
}

func findStartCoord(grid [][]string) Coord {
	for i, row := range grid {
		for j, col := range row {
			if col == "S" {
				return Coord{i, j}
			}
		}
	}
	panic("no start coordinate found")
}

func toString(coord Coord) string {
	return strconv.Itoa(coord.row) + "," + strconv.Itoa(coord.col)
}

func getNeighbours(c Coord, grid [][]string) []Coord {
	var res []Coord
	if c.row > 0 && grid[c.row-1][c.col] == "." {
		res = append(res, Coord{c.row - 1, c.col})
	}
	if c.row < len(grid)-1 && grid[c.row+1][c.col] == "." {
		res = append(res, Coord{c.row + 1, c.col})
	}
	if c.col > 0 && grid[c.row][c.col-1] == "." {
		res = append(res, Coord{c.row, c.col - 1})
	}
	if c.col < len(grid[0])-1 && grid[c.row][c.col+1] == "." {
		res = append(res, Coord{c.row, c.col + 1})
	}

	return res
}
