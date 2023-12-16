package main

import (
	"aoc_2023_go/util"
	"fmt"
	"strings"
	"time"
)

type Tile struct {
	coord util.Coord
	value string
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
	pipelines := readPipe(lines)
	res := findFurthestLocation(pipelines)
	fmt.Printf("answer: %d\n", res)
}

func two(lines []string) {
	pipelines := readPipe(lines)
	res := findArea(pipelines)
	fmt.Printf("answer: %d\n", res)
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

func findArea(pipelines [][]string) int {
	prev := findStartCoord(pipelines)
	cur := findFirstPipe(pipelines, prev)
	visited := []Tile{{coord: prev, value: pipelines[prev.X][prev.Y]}}
	for pipelines[cur.X][cur.Y] != "S" {
		visited = append(visited, Tile{cur, pipelines[cur.X][cur.Y]})
		next := findNextPipe(cur, prev, pipelines)
		prev = cur
		cur = next
	}

	area := 0

	var visitedCoords []util.Coord
	for _, v := range visited {
		visitedCoords = append(visitedCoords, v.coord)
	}
	// replace start coordinate with the actual pipe
	pipelines = replaceStartCoord(pipelines, visitedCoords)
	// replace all junk with ground
	for x := range pipelines {
		for y := range pipelines[x] {
			if !util.InSlice(util.Coord{X: x, Y: y}, visitedCoords) {
				pipelines[x][y] = "."
			}
		}
	}

	for _, row := range pipelines {
		isIn := false
		for _, col := range row {
			if isIn && col == "." {
				area += 1
			}
			if col == "|" {
				isIn = !isIn
			}
			if col == "J" {
				isIn = !isIn
			}
			if col == "L" {
				isIn = !isIn
			}
		}
	}
	return area
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

func replaceStartCoord(pipelines [][]string, vc []util.Coord) [][]string {
	sc := findStartCoord(pipelines)
	if (vc[1].X < sc.X && vc[len(vc)-1].Y > sc.Y) || (vc[len(vc)-1].X < sc.X && vc[1].Y > sc.Y) {
		pipelines[sc.X][sc.Y] = "L"
	}
	if (vc[1].X > sc.X && vc[len(vc)-1].Y > sc.Y) || (vc[len(vc)-1].X > sc.X && vc[1].Y > sc.Y) {
		pipelines[sc.X][sc.Y] = "F"
	}
	if (vc[1].X < sc.X && vc[len(vc)-1].Y < sc.Y) || (vc[len(vc)-1].X < sc.X && vc[1].Y < sc.Y) {
		pipelines[sc.X][sc.Y] = "J"
	}
	if (vc[1].X > sc.X && vc[len(vc)-1].Y < sc.Y) || (vc[len(vc)-1].X > sc.X && vc[1].Y < sc.Y) {
		pipelines[sc.X][sc.Y] = "7"
	}
	return pipelines
}
