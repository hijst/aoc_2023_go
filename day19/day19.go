package main

import (
	"aoc_2023_go/util"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Part struct {
	x int
	m int
	a int
	s int
}

type Condition struct {
	part      string
	operator  string
	threshold int
	target    string
}

type Range struct {
	mn int64
	mx int64
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
	parts, conditions := readInput(lines)
	res := processParts(parts, conditions)
	fmt.Printf("answer: %d\n", res)
}

func two(lines []string) {
	_, conditions := readInput(lines)
	res := findCombinations("in", conditions, map[string]Range{"x": {1, 4000}, "m": {1, 4000}, "a": {1, 4000}, "s": {1, 4000}})
	fmt.Printf("answer: %d\n", res)
}

func readInput(lines []string) ([]Part, map[string][]Condition) {
	var partsStartLine int
	conditions := make(map[string][]Condition)
	var parts []Part

	for i, line := range lines {
		if line == "" {
			partsStartLine = i + 1
			break
		}
		s := strings.Split(line[:len(line)-1], "{")
		cs := strings.Split(s[1], ",")
		var conds []Condition
		for _, c := range cs {
			if !strings.Contains(c, ":") {
				conds = append(conds, Condition{part: "no", operator: "no", threshold: 0, target: c})
			} else {
				sp := strings.Split(c, ":")
				conds = append(conds, Condition{part: sp[0][0:1], operator: sp[0][1:2], threshold: atoi(sp[0][2:]), target: sp[1]})
			}
		}
		conditions[s[0]] = conds
	}

	for _, line := range lines[partsStartLine:] {
		s := line[1 : len(line)-1]
		p := strings.Split(s, ",")
		parts = append(parts, Part{x: atoi(p[0][2:]), m: atoi(p[1][2:]), a: atoi(p[2][2:]), s: atoi(p[3][2:])})
	}
	return parts, conditions
}

func atoi(s string) int {
	v, _ := strconv.Atoi(s)
	return v
}

func processParts(parts []Part, conditions map[string][]Condition) int {
	var res int
	for _, part := range parts {
		res += applyCondition(part, "in", conditions)
	}
	return res
}

func applyCondition(part Part, condition string, conditions map[string][]Condition) int {
	cond := conditions[condition]
	for _, c := range cond {
		if c.target == "A" && c.operator == "no" {
			return part.x + part.m + part.a + part.s
		}
		if c.target == "R" && c.operator == "no" {
			return 0
		}
		if c.operator == "no" {
			return applyCondition(part, c.target, conditions)
		}
		var v int
		switch c.part {
		case "x":
			v = part.x
		case "m":
			v = part.m
		case "a":
			v = part.a
		default:
			v = part.s
		}
		if c.operator == ">" {
			if v > c.threshold {
				if c.target == "A" {
					return part.x + part.m + part.a + part.s
				}
				if c.target == "R" {
					return 0
				}
				return applyCondition(part, c.target, conditions)
			}
		}
		if c.operator == "<" {
			if v < c.threshold {
				if c.target == "A" {
					return part.x + part.m + part.a + part.s
				}
				if c.target == "R" {
					return 0
				}
				return applyCondition(part, c.target, conditions)
			}
		}

	}
	panic("FAILED APPLYING CONDITION")
}

func findCombinations(condition string, conditions map[string][]Condition, pr map[string]Range) int64 {
	var res int64
	p := copyMap(pr)
	cond := conditions[condition]
	for _, c := range cond {
		np := copyMap(p)
		switch c.operator {
		case "<":
			np[c.part] = Range{mn: np[c.part].mn, mx: int64(c.threshold) - 1}
			p[c.part] = Range{mn: int64(c.threshold), mx: p[c.part].mx}
		case ">":
			np[c.part] = Range{mn: int64(c.threshold) + 1, mx: np[c.part].mx}
			p[c.part] = Range{mn: p[c.part].mn, mx: int64(c.threshold)}
		}
		switch c.target {
		case "A":
			res += resolvePartRange(np)
		case "R":
		default:
			res += findCombinations(c.target, conditions, np)
		}
	}
	return res
}

func resolvePartRange(pr map[string]Range) int64 {
	return (1 + pr["x"].mx - pr["x"].mn) * (1 + pr["m"].mx - pr["m"].mn) * (1 + pr["a"].mx - pr["a"].mn) * (1 + pr["s"].mx - pr["s"].mn)
}

func copyMap(m map[string]Range) map[string]Range {
	nm := make(map[string]Range)
	for k, v := range m {
		nm[k] = v
	}
	return nm
}
