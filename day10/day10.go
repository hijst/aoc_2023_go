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
	pipelines := readPipe(lines)
	res := findFurthestLocation(pipelines)
	fmt.Printf("answer: %d\n", res)
}

func two(lines []string) {
	fmt.Printf("part two not yet implemented, number of lines: %d\n", len(lines))
}

func readPipe(lines []string) [][]string {
	var res [][]string
	for _, line := range lines {
		res = append(res, strings.Split(line, ""))
	}
	return res
}

func findFurthestLocation(pipelines [][]string) int {
	prev := findStartCoord(pipelines)
	cur := findFirstPipe(pipelines, prev)
	visited := []util.Coord{prev}
	for pipelines[cur.X][cur.Y] != "S" {
		visited = append(visited, cur)
		next := findNextPipe(cur, prev, pipelines)
		prev = cur
		cur = next
	}

	return len(visited) / 2
}

type Tile struct {
	coord util.Coord
	value string
}

func findArea(pipelines [][]string) int {
	prev := findStartCoord(pipelines)
	cur := findFirstPipe(pipelines, prev)
	visited := []Tile{prev}
	for pipelines[cur.X][cur.Y] != "S" {
		visited = append(visited, Tile{cur, pipelines[cur.X][cur.Y]})
		next := findNextPipe(cur, prev, pipelines)
		prev = cur
		cur = next
	}

	area := 0

	for x, row := range pipelines {
		for y := range row {
			if isEnclosed(x, y, visited, pipelines) {
				area += 1
			}
		}
	}
	return area - len(visited)
}

func isEnclosed(x, y int, visited []Tile, pipelines [][]string) bool {
	if x == 0 || x == len(pipelines)-1 || y == 0 || y == len(pipelines)-1 {
		return false // on the edge
	}
	if util.InSlice(Tile{coord: util.Coord{X: x, Y: y}, value: pipelines[x][y]}, visited) {
		return false // is a pipe
	}
	return leftOk(x, y, visited) && rightOk(x, y, visited) && upOk(x, y, visited) && downOk(x, y, visited)
}

func leftOk(x, y int, visited []Tile) bool {
	var count int
	var visitedCoords []util.Coord
	for col := 0; col < y; col++ {
		cur := util.Coord{X: x, Y: col}
		if util.InSlice(cur, visitedCoords) {
		}
	}
	return count%2 == 1
}

func findStartCoord(pipelines [][]string) util.Coord {
	for i, row := range pipelines {
		for j, col := range row {
			if col == "S" {
				return util.Coord{X: i, Y: j}
			}
		}
	}
	panic("no start coordinate present in input!")
}

func findFirstPipe(pipelines [][]string, start util.Coord) util.Coord {
	for _, neigh := range util.GetDirectNeighbourCoords(start.X, start.Y, pipelines) {
		if neigh.X < start.X && util.InSlice(pipelines[neigh.X][neigh.Y], []string{"|", "F", "7"}) {
			return neigh
		}
		if neigh.Y < start.Y && util.InSlice(pipelines[neigh.X][neigh.Y], []string{"-", "F", "L"}) {
			return neigh
		}
		if neigh.X > start.X && util.InSlice(pipelines[neigh.X][neigh.Y], []string{"|", "J", "L"}) {
			return neigh
		}
		if neigh.Y > start.Y && util.InSlice(pipelines[neigh.X][neigh.Y], []string{"-", "J", "7"}) {
			return neigh
		}

	}
	panic("no second coordinate found")
}

func findNextPipe(cur util.Coord, prev util.Coord, pipelines [][]string) util.Coord {
	candidates := util.RemoveFromSlice(prev, util.GetDirectNeighbourCoords(cur.X, cur.Y, pipelines))
	curval := pipelines[cur.X][cur.Y]
	if curval == "|" {
		candidates = util.RemoveFromSlice(util.Coord{X: cur.X, Y: cur.Y - 1}, candidates)
		candidates = util.RemoveFromSlice(util.Coord{X: cur.X, Y: cur.Y + 1}, candidates)
	}
	if curval == "-" {
		candidates = util.RemoveFromSlice(util.Coord{X: cur.X + 1, Y: cur.Y}, candidates)
		candidates = util.RemoveFromSlice(util.Coord{X: cur.X - 1, Y: cur.Y}, candidates)
	}
	if curval == "L" {
		candidates = util.RemoveFromSlice(util.Coord{X: cur.X, Y: cur.Y - 1}, candidates)
		candidates = util.RemoveFromSlice(util.Coord{X: cur.X + 1, Y: cur.Y}, candidates)
	}
	if curval == "J" {
		candidates = util.RemoveFromSlice(util.Coord{X: cur.X + 1, Y: cur.Y}, candidates)
		candidates = util.RemoveFromSlice(util.Coord{X: cur.X, Y: cur.Y + 1}, candidates)
	}
	if curval == "7" {
		candidates = util.RemoveFromSlice(util.Coord{X: cur.X - 1, Y: cur.Y}, candidates)
		candidates = util.RemoveFromSlice(util.Coord{X: cur.X, Y: cur.Y + 1}, candidates)
	}
	if curval == "F" {
		candidates = util.RemoveFromSlice(util.Coord{X: cur.X, Y: cur.Y - 1}, candidates)
		candidates = util.RemoveFromSlice(util.Coord{X: cur.X - 1, Y: cur.Y}, candidates)
	}
	return candidates[0]
}
