package main

import (
	"aoc_2023_go/util"
	"fmt"
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
	grid := util.ReadGrid(lines)
	expanded := markEmpty(grid)
	res := calculateMinDistanceSum(expanded, int64(2))
	fmt.Printf("answer: %d\n", res)
}

func two(lines []string) {
	grid := util.ReadGrid(lines)
	expanded := markEmpty(grid)
	res := calculateMinDistanceSum(expanded, int64(1000000))
	fmt.Printf("answer: %d\n", res)
}

func calculateMinDistanceSum(grid [][]string, evalue int64) int64 {
	galaxies := findGalaxies(grid)
	var res int64
	for i := 0; i < len(galaxies)-1; i++ {
		for j := i + 1; j < len(galaxies); j++ {
			res += shortestPath(Coord{galaxies[i].row, galaxies[i].col}, Coord{galaxies[j].row, galaxies[j].col}, grid, evalue)
		}
	}
	return res
}

func shortestPath(source, target Coord, grid [][]string, evalue int64) int64 {
	var res int64
	b, e := min(source.row, target.row)
	for i := b + 1; i <= e; i++ {
		if grid[i][0] == "E" {
			res += evalue
		} else {
			res += 1
		}
	}
	bc, ec := min(source.col, target.col)
	for j := bc + 1; j <= ec; j++ {
		if grid[0][j] == "E" {
			res += evalue
		} else {
			res += 1
		}
	}
	return res
}

func min(a, b int) (int, int) {
	if a < b {
		return a, b
	}
	return b, a
}

func markEmpty(grid [][]string) [][]string {
	var temp [][]string
	for _, row := range grid {
		if isEmptySpace(row) {
			temp = append(temp, toEmpty(row))
		} else {
			temp = append(temp, row)
		}
	}
	temp = util.Transpose(temp)
	var res [][]string
	for _, row := range temp {
		if isEmptySpace(row) {
			res = append(res, toEmpty(row))
		} else {
			res = append(res, row)
		}
	}
	return util.Transpose(res)
}

func toEmpty(row []string) []string {
	var res []string
	for range row {
		res = append(res, "E")
	}
	return res
}

func findGalaxies(grid [][]string) []Coord {
	var galaxies []Coord

	for i, row := range grid {
		for j, col := range row {
			if col == "#" {
				galaxies = append(galaxies, Coord{i, j})
			}
		}
	}
	return galaxies
}

func isEmptySpace(row []string) bool {
	for _, char := range row {
		if char == "#" {
			return false
		}
	}
	return true
}
