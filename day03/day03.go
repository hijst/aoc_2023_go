package main

import (
	"aoc_2023_go/util"
	"fmt"
	"regexp"
	"strconv"
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
	sum := 0
	grid := linesToGrid(lines)

	for i := 0; i < len(grid[0]); i++ {
		for j := 0; j < len(grid); {
			if isNumeric(grid[i][j]) {
				numericString := findNumberAsString(j, grid[i])
				number, _ := strconv.Atoi(numericString)
				if anyDigitHasCorrectNeighbour(i, j, len(numericString), grid) {
					sum += number
				}
				j += len(numericString)
			} else {
				j++
			}
		}
	}
	fmt.Printf("answer: %d\n", sum)
}

func two(lines []string) {
	grid := linesToGrid(lines)
	gears := make([][][]int64, len(grid))
	for i := range gears {
		gears[i] = make([][]int64, len(grid))
		for j := range gears[i] {
			gears[i][j] = make([]int64, 0)
		}
	}

	for i := 0; i < len(grid[0]); i++ {
		for j := 0; j < len(grid); {
			if isNumeric(grid[i][j]) {
				numericString := findNumberAsString(j, grid[i])
				number32, _ := strconv.Atoi(numericString)
				number := int64(number32)
				gears = updateGears(gears, grid, number, i, j, len(numericString))
				j += len(numericString)
			} else {
				j++
			}
		}
	}
	sum := calculateGearRatiosSum(gears)
	fmt.Printf("answer: %d\n", sum)
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

func anyDigitHasCorrectNeighbour(row int, col int, length int, grid [][]rune) bool {
	for c := col; c < col+length; c++ {
		if hasCorrectNeighbour(row, c, grid) {
			return true
		}
	}
	return false
}

func hasCorrectNeighbour(row int, col int, grid [][]rune) bool {
	neighbours := getNeighbours(row, col, grid)
	for _, n := range neighbours {
		if !isNumeric(n) && n != '.' {
			return true
		}
	}
	return false
}

func getNeighbours(row int, col int, grid [][]rune) []rune {
	nrows := len(grid)
	ncols := len(grid[0])

	var neighbours []rune

	srow := max(0, row-1)
	erow := min(nrows-1, row+1)

	scol := max(0, col-1)
	ecol := min(ncols-1, col+1)

	for i := srow; i <= erow; i++ {
		for j := scol; j <= ecol; j++ {
			if i == row && j == col {
				continue
			} else {
				neighbours = append(neighbours, grid[i][j])
			}
		}
	}
	return neighbours
}

func getNeighbourGearCoords(row int, col int, grid [][]rune) []Coord {
	nrows := len(grid)
	ncols := len(grid[0])

	var neighbours []Coord

	srow := max(0, row-1)
	erow := min(nrows-1, row+1)

	scol := max(0, col-1)
	ecol := min(ncols-1, col+1)

	for i := srow; i <= erow; i++ {
		for j := scol; j <= ecol; j++ {
			if i == row && j == col {
				continue
			} else {
				if grid[i][j] == '*' {
					neighbours = append(neighbours, Coord{Row: i, Col: j})
				}
			}
		}
	}
	return neighbours
}

type Coord struct {
	Row int
	Col int
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func deduplicateSplice(splice []Coord) []Coord {
	uniqueElements := make(map[Coord]struct{})
	for _, c := range splice {
		uniqueElements[c] = struct{}{}
	}
	var res []Coord
	for c := range uniqueElements {
		res = append(res, c)
	}

	return res
}

func updateGears(gears [][][]int64, grid [][]rune, number int64, row int, col int, length int) [][][]int64 {
	var gearNeighbourCoords []Coord
	for c := col; c < col+length; c++ {
		gearNeighbourCoords = append(gearNeighbourCoords, getNeighbourGearCoords(row, c, grid)...)
	}
	deduped := deduplicateSplice(gearNeighbourCoords)

	for _, d := range deduped {
		gears[d.Row][d.Col] = append(gears[d.Row][d.Col], number)
	}

	return gears
}

func calculateGearRatiosSum(gears [][][]int64) int64 {
	res := int64(0)
	for i, row := range gears {
		for j := range row {
			if len(gears[i][j]) == 2 {
				res += int64(gears[i][j][0] * gears[i][j][1])
			}
		}
	}
	return res
}
