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

type Beam struct {
	pos Coord
	dir Coord
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
	firstBeam := Beam{Coord{0, -1}, Coord{0, 1}}
	res := resolve(grid, firstBeam)
	fmt.Printf("answer: %d\n", res)
}

func two(lines []string) {
	grid := readGrid(lines)
	res := tryAllStartingPositions(grid)
	fmt.Printf("answer: %d\n", res)
}

func readGrid(lines []string) [][]string {
	var res [][]string
	for _, line := range lines {
		res = append(res, strings.Split(line, ""))
	}
	return res
}

func tryAllStartingPositions(grid [][]string) int {
	var res int
	for i := 0; i < len(grid); i++ {
		r1 := resolve(grid, Beam{Coord{i, -1}, Coord{0, 1}})
		r2 := resolve(grid, Beam{Coord{i, len(grid)}, Coord{0, -1}})
		r := max(r1, r2)
		res = max(res, r)
	}
	for i := 0; i < len(grid[0]); i++ {
		r1 := resolve(grid, Beam{Coord{-1, i}, Coord{1, 0}})
		r2 := resolve(grid, Beam{Coord{len(grid[0]), i}, Coord{-1, 0}})
		r := max(r1, r2)
		res = max(res, r)
	}
	return res
}

func resolve(grid [][]string, start Beam) int {
	energized := make(map[string]bool)
	beams := make(map[string]Beam)
	seen := make(map[string]Beam)
	beams[toString(start)] = start
	newStates := true
	for newStates {
		newStates = false
		for hash, beam := range beams {
			energized[cString(beam.pos)] = true
			newPos := addCoords(beam.pos, beam.dir)
			if outOfBounds(newPos, grid) || exists(seen, hash) {
				delete(beams, hash)
			} else {
				newStates = true
				seen[hash] = beam
				switch grid[newPos.row][newPos.col] {
				case ".":
					newBeam := Beam{newPos, beam.dir}
					beams[toString(newBeam)] = newBeam
				case "-":
					if beam.dir.row == 0 {
						newBeam := Beam{newPos, beam.dir}
						beams[toString(newBeam)] = newBeam
					} else {
						newBeam1 := Beam{newPos, Coord{0, -1}}
						newBeam2 := Beam{newPos, Coord{0, 1}}
						beams[toString(newBeam1)] = newBeam1
						beams[toString(newBeam2)] = newBeam2
					}
				case "|":
					if beam.dir.col == 0 {
						newBeam := Beam{newPos, beam.dir}
						beams[toString(newBeam)] = newBeam
					} else {
						newBeam1 := Beam{newPos, Coord{1, 0}}
						newBeam2 := Beam{newPos, Coord{-1, 0}}
						beams[toString(newBeam1)] = newBeam1
						beams[toString(newBeam2)] = newBeam2
					}
				case "/":
					if beam.dir.row == 0 {
						newDir := Coord{-1 * beam.dir.col, 0}
						beams[toString(Beam{newPos, newDir})] = Beam{newPos, newDir}
					} else {
						newDir := Coord{0, -1 * beam.dir.row}
						beams[toString(Beam{newPos, newDir})] = Beam{newPos, newDir}
					}
				default:
					if beam.dir.row == 0 {
						newDir := Coord{beam.dir.col, 0}
						beams[toString(Beam{newPos, newDir})] = Beam{newPos, newDir}
					} else {
						newDir := Coord{0, beam.dir.row}
						beams[toString(Beam{newPos, newDir})] = Beam{newPos, newDir}
					}
				}
			}
		}
	}
	return len(energized) - 1
}

func outOfBounds(coord Coord, grid [][]string) bool {
	return coord.col < 0 || coord.row < 0 || coord.row > len(grid)-1 || coord.col > len(grid[0])-1
}

func cString(coord Coord) string {
	return strconv.Itoa(coord.row) + "," + strconv.Itoa(coord.col)
}

func toString(beam Beam) string {
	return strconv.Itoa(beam.pos.row) + "," + strconv.Itoa(beam.pos.col) + "," + strconv.Itoa(beam.dir.row) + "," + strconv.Itoa(beam.dir.col)
}

func addCoords(c1, c2 Coord) Coord {
	return Coord{c1.row + c2.row, c1.col + c2.col}
}

func exists(m map[string]Beam, k string) bool {
	_, exists := m[k]
	return exists
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
