package main

import (
	"aoc_2023_go/util"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Instruction struct {
	dir   Direction
	steps int
}

type Direction struct {
	v int
	h int
}

type Position struct {
	row int
	col int
}

var (
	up    = Direction{-1, 0}
	down  = Direction{1, 0}
	left  = Direction{0, -1}
	right = Direction{0, 1}
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
	instructions := readInstructions(lines)
	grid := make([][]string, len(instructions)*2)
	for i := range grid {
		grid[i] = make([]string, len(instructions)*2)
		for j := range grid[i] {
			grid[i][j] = "."
		}
	}
	dig(grid, instructions)
	res := countIn(grid, instructions)
	fmt.Printf("answer: %d\n", res)
}

func two(lines []string) {
	instructions := readInstructions2(lines)
	res := trapezoid(instructions)
	fmt.Printf("answer: %d\n", res)
}

func readInstructions(lines []string) []Instruction {
	var res []Instruction
	dirMap := map[string]Direction{"U": up, "D": down, "L": left, "R": right}
	for _, line := range lines {
		parts := strings.Fields(line)
		res = append(res, Instruction{dirMap[parts[0]], atoi(parts[1])})
	}
	return res
}

func readInstructions2(lines []string) []Instruction {
	var res []Instruction
	dirMap := map[string]Direction{"0": right, "1": down, "2": left, "3": up}
	for _, line := range lines {
		in := strings.Fields(line)[2]
		steps := in[2 : len(in)-2]
		dir := dirMap[in[len(in)-2:len(in)-1]]
		stepsC, _ := strconv.ParseInt(steps, 16, 0)
		res = append(res, Instruction{dir, int(stepsC)})

	}
	return res
}

func dig(grid [][]string, instructions []Instruction) [][]string {
	row := len(instructions)
	col := row
	grid[row][col] = "#"
	for _, ins := range instructions {
		for i := 0; i < ins.steps; i++ {
			row += ins.dir.v
			col += ins.dir.h

			grid[row][col] = "#"
		}
	}
	return grid
}

func countIn(grid [][]string, instructions []Instruction) int {
	visited := make(map[string]bool)
	row := len(instructions)
	col := row
	rDirs := []Direction{up, right, down, left}
	grid[row][col] = "#"
	for _, ins := range instructions { // walk the outline again
		for i := 0; i < ins.steps; i++ {
			row += ins.dir.v
			col += ins.dir.h
			visited[hash(Position{row, col})] = true // mark outline visited
			d := nextDir(rDirs, ins.dir)             // check the inside (along the right of the "wall")
			rPos := Position{row + d.v, col + d.h}
			if grid[rPos.row][rPos.col] == "." { // bfs
				s := []Position{rPos}
				for len(s) != 0 {
					cur := s[0]
					if !isVisited(visited, cur) {
						visited[hash(cur)] = true // mark inside visited
						for _, d := range rDirs {
							next := Position{cur.row + d.v, cur.col + d.h}
							if grid[next.row][next.col] == "." {
								s = append(s, next) // push
							}
						}
					}
					s = s[1:] // pop
				}
			}
		}
	}

	return len(visited)
}

func trapezoid(ins []Instruction) int64 { // applying the trapezoid formula https://en.wikipedia.org/wiki/Shoelace_formula
	cur := Position{0, 0}
	prev := Position{0, 0}
	var res int64
	var sb int64
	sb += 1 // we will also add 1 side of the boundary (including the starting position) to the result because the trench also has volume
	for i := 0; i < len(ins); i++ {
		prev.row = cur.row
		prev.col = cur.col
		in := ins[i]
		switch in.dir {
		case up:
			cur.row -= in.steps
		case down:
			sb += int64(in.steps)
			cur.row += in.steps
		case right:
			sb += int64(in.steps)
			cur.col += in.steps
		default:
			cur.col -= in.steps
		}
		res += (int64(cur.row) + int64(prev.row)) * (int64(prev.col) - int64(cur.col))

	}

	return sb + res/2
}

func nextDir(s []Direction, t Direction) Direction {
	for i, v := range s {
		if v.v == t.v && v.h == t.h {
			return s[(i+1)%len(s)]
		}
	}
	panic("invalid dir")
}

func isVisited(visited map[string]bool, p Position) bool {
	_, e := visited[hash(p)]
	return e
}

func hash(d Position) string {
	return itoa(d.row) + "," + itoa(d.col)
}

func itoa(i int) string {
	return strconv.Itoa(i)
}

func atoi(s string) int {
	v, _ := strconv.Atoi(s)
	return v
}
