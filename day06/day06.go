package main

import (
	"aoc_2023_go/util"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

type Race struct {
	t float64
	d float64
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
	races := parseInput(lines)
	product := int64(1)

	for _, race := range races {
		ways := solveRace(race)
		product *= ways
	}

	fmt.Printf("answer: %d\n", product)
}

func two(lines []string) {
	race := parseRace(lines)
	ways := solveRace(race)
	fmt.Printf("answer: %d\n", ways)
}

func parseInput(lines []string) []Race {
	var res []Race
	var ts []float64
	var ds []float64
	firstCut, _ := strings.CutPrefix(lines[0], "Time: ")
	firstSplit := strings.Fields(firstCut)
	for _, split := range firstSplit {
		time, _ := strconv.ParseFloat(split, 64)
		ts = append(ts, time)
	}

	secondCut, _ := strings.CutPrefix(lines[1], "Distance: ")
	secondSplit := strings.Fields(secondCut)
	for _, split := range secondSplit {
		distance, _ := strconv.ParseFloat(split, 64)
		ds = append(ds, distance)
	}

	for i := range ts {
		res = append(res, Race{ts[i], ds[i]})
	}
	return res
}

func parseRace(lines []string) Race {
	firstCut, _ := strings.CutPrefix(lines[0], "Time: ")
	firstSplitJoined := strings.Join(strings.Fields(firstCut), "")
	t, _ := strconv.ParseFloat(firstSplitJoined, 64)

	secondCut, _ := strings.CutPrefix(lines[1], "Distance: ")
	secondSplitJoined := strings.Join(strings.Fields(secondCut), "")
	d, _ := strconv.ParseFloat(secondSplitJoined, 64)
	return Race{t, d}
}

func solveRace(race Race) int64 {
	var sol1 float64
	var sol2 float64
	sol1 = float64(0.5) * (race.t - math.Sqrt(math.Pow(race.t, 2)-4*race.d))
	sol2 = float64(0.5) * (math.Sqrt(math.Pow(race.t, 2)-4*race.d) + race.t)

	minSol := minFloat(sol1, sol2) + 0.00001
	maxSol := maxFloat(sol1, sol2) - 0.00001

	minTime := max(0, int64(math.Ceil(minSol)))
	maxTime := min(int64(race.t), int64(math.Floor(maxSol)))
	return 1 + maxTime - minTime
}

func minFloat(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

func maxFloat(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
