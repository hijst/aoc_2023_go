package main

import (
	"aoc_2023_go/util"
	"fmt"
	"regexp"
	"time"
)

type Node struct {
	left  string
	right string
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
	instructions, nodes := parseInput(lines)
	steps := findSteps(instructions, nodes)
	fmt.Printf("answer: %d\n", steps)
}

func two(lines []string) {
	fmt.Printf("part two not yet implemented, number of lines: %d\n", len(lines))
}

func parseInput(lines []string) (string, map[string]Node) {
	instructions := lines[0]
	nodes := make(map[string]Node)
	for _, line := range lines[2:] {
		name, node := parseNodeMapping(line)
		nodes[name] = node
	}
	return instructions, nodes
}

func parseNodeMapping(line string) (string, Node) {
	re := regexp.MustCompile("([A-Z])+")
	parts := re.FindAllString(line, -1)
	name := parts[0]
	node := Node{
		parts[1],
		parts[2],
	}
	return name, node
}

func findSteps(instructions string, nodes map[string]Node) int64 {
	steps := int64(0)
	cur := "AAA"
	i := 0
	for cur != "ZZZ" {

		if instructions[i] == 'L' {
			cur = nodes[cur].left
		} else {
			cur = nodes[cur].right
		}

		if i == len(instructions)-1 {
			i = 0
		} else {
			i++
		}
		steps++
	}
	return steps
}
