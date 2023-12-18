package main

import (
	"aoc_2023_go/util"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"gonum.org/v1/gonum/stat/combin"
)

type Group struct {
	springs   []string
	condition []int
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
	groups := parseGroups(lines)
	var sum int64
	for _, group := range groups {
		sum += solveGroup(group)
	}
	fmt.Printf("answer: %d\n", sum)
}

func two(lines []string) {
	fmt.Printf("part two not yet implemented, number of lines: %d\n", len(lines))
}

func parseGroups(lines []string) []Group {
	var groups []Group
	for _, line := range lines {
		parts := strings.Fields(line)
		springs := strings.Split(parts[0], "")
		conditionBase := strings.Split(parts[1], ",")
		condition := util.Map(conditionBase, atoi)
		groups = append(groups, Group{springs: springs, condition: condition})
	}
	return groups
}

func solveGroup(group Group) int64 {
	var res int64
	unknownCount := countUnknown(group)
	requiredCount := countRequired(group)
	combinations := combin.Combinations(unknownCount, requiredCount)
	for _, combination := range combinations {
		springConfig := configureSprings(combination, group.springs)
		if satisfies(springConfig, group.condition) {
			res += 1
		}
	}
	return res
}

func configureSprings(combination []int, springs []string) []string {
	var res []string
	ci := 0
	cl := len(combination)
	i := 0
	for _, spring := range springs {
		if spring == "?" {
			if ci >= cl || combination[ci] != i {
				res = append(res, ".")
			} else {
				ci += 1
				res = append(res, "#")
			}
			i += 1
		} else {
			res = append(res, spring)
		}
	}
	return res
}

func countUnknown(group Group) int {
	var count int
	for _, char := range group.springs {
		if char == "?" {
			count += 1
		}
	}
	return count
}

func countRequired(group Group) int {
	var count int
	for _, num := range group.condition {
		count += num
	}
	for _, char := range group.springs {
		if char == "#" {
			count -= 1
		}
	}
	return count
}

func satisfies(springs []string, condition []int) bool {
	springsCondition := toSpringsCondition(springs)
	return reflect.DeepEqual(springsCondition, condition)
}

func toSpringsCondition(springs []string) []int {
	var res []int
	prevSpring := false
	count := 0
	for _, spring := range springs {
		if spring == "#" {
			prevSpring = true
			count += 1
		} else {
			if prevSpring {
				res = append(res, count)
				count = 0
			}
			prevSpring = false
		}
	}
	if count != 0 {
		res = append(res, count)
	}
	return res
}

func atoi(s string) int {
	r, _ := strconv.Atoi(s)
	return r
}
