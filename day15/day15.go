package main

import (
	"aoc_2023_go/util"
	"fmt"
	"strconv"
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

type Step struct {
	label       string
	instruction string
	focalLength int
}

func one(lines []string) {
	rawString := split(lines)
	res := calculateHashScores(rawString)
	fmt.Printf("answer: %d\n", res)
}

func two(lines []string) {
	steps := getSteps(lines)
	res := getFocusPower(steps)
	fmt.Printf("answer: %d\n", res)
}

func split(lines []string) []string {
	return strings.Split(lines[0], ",")
}

func getSteps(lines []string) []Step {
	var res []Step
	rawStrings := split(lines)
	for _, s := range rawStrings {
		i := strings.Split(s, "=")
		if len(i) > 1 {
			fl, _ := strconv.Atoi(i[1])
			res = append(res, Step{label: i[0], instruction: "=", focalLength: fl})
		} else {
			j := strings.Split(i[0], "-")
			res = append(res, Step{label: j[0], instruction: "-", focalLength: -1})
		}
	}
	return res
}

func getFocusPower(steps []Step) int64 {
	boxes := make([][]Step, 256)
	for _, step := range steps {
		h := calculateHash(step.label)
		if step.instruction == "=" {
			boxes[h] = addToBox(boxes[h], step)
		} else {
			boxes[h] = removeByLabel(boxes[h], step.label)
		}
	}
	return calculateFocusPower(boxes)
}

func calculateFocusPower(boxes [][]Step) int64 {
	var res int64
	for i, box := range boxes {
		for j, step := range box {
			res += int64(i+1) * int64(j+1) * int64(step.focalLength)
		}
	}
	return res
}

func addToBox(steps []Step, step Step) []Step {
	var res []Step
	found := false
	for _, s := range steps {
		if s.label == step.label {
			res = append(res, step)
			found = true
		} else {
			res = append(res, s)
		}
	}
	if !found {
		res = append(res, step)
	}
	return res
}

func removeByLabel(steps []Step, label string) []Step {
	var res []Step
	for _, s := range steps {
		if s.label != label {
			res = append(res, s)
		}
	}
	return res
}

func calculateHashScores(rawStrings []string) int {
	var res int
	for _, s := range rawStrings {
		res += calculateHash(s)
	}
	return res
}

func calculateHash(s string) int {
	var res int
	for _, char := range s {
		res += int(char)
		res *= 17
		res %= 256
	}
	return res
}
