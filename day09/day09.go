package main

import (
	"aoc_2023_go/util"
	"fmt"
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
	histories := util.GetInt64Fields(lines)
	var res int64
	for _, history := range histories {
		res += extrapolate(history)
	}
	fmt.Printf("answer: %d\n", res)
}

func two(lines []string) {
	histories := util.GetInt64Fields(lines)
	var res int64
	for _, history := range histories {
		reversedHistory := util.ReverseSlice(history)
		res += extrapolate(reversedHistory)
	}
	fmt.Printf("answer: %d\n", res)
}

func extrapolate(history []int64) int64 {
	expansion := expand(history)
	resolved := resolve(expansion)
	return resolved[0][len(resolved[0])-1]
}

func expand(history []int64) [][]int64 {
	var expansion [][]int64
	expansion = append(expansion, history)

	for !hasOnlyZeros(expansion[len(expansion)-1]) {
		expansion = append(expansion, expandLine(expansion[len(expansion)-1]))
	}
	expansion[len(expansion)-1] = append(expansion[len(expansion)-1], int64(0))
	return expansion
}

func expandLine(line []int64) []int64 {
	var res []int64
	for i := 0; i < len(line)-1; i++ {
		res = append(res, line[i+1]-line[i])
	}
	return res
}

func resolve(expansion [][]int64) [][]int64 {
	for i := len(expansion) - 2; i >= 0; i-- {
		expansion[i] = append(expansion[i], util.LastElement(expansion[i])+util.LastElement(expansion[i+1]))
	}
	return expansion
}

func hasOnlyZeros(s []int64) bool {
	for _, value := range s {
		if value != int64(0) {
			return false
		}
	}
	return true
}
