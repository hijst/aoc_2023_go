package util

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func ReadFileLines(subdirectory string) ([]string, error) {
	var fileName string

	if len(os.Args) >= 2 {
		fileName = os.Args[1]
	} else {
		fileName = "input.txt"
	}
	path := filepath.Join(subdirectory, fileName)
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("error opening file '%s': %v", path, err)
	}

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file '%s': %v", path, err)
	}
	file.Close()
	return lines, nil
}

func IntPow(n, m int) int {
	if m == 0 {
		return 1
	}
	result := n

	for i := 2; i <= m; i++ {
		result *= n
	}
	return result
}

func DeduplicateSplice[T comparable](splice []T) []T {
	uniqueElements := make(map[T]bool)
	for _, c := range splice {
		uniqueElements[c] = true
	}

	var res []T
	for element := range uniqueElements {
		res = append(res, element)
	}

	return res
}

func MapValues[K, V comparable](m map[K]V) []V {
	var res []V
	for _, v := range m {
		res = append(res, v)
	}
	return res
}

func MapKeys[K, V comparable](m map[K]V) []K {
	var res []K
	for k := range m {
		res = append(res, k)
	}
	return res
}

func GCD(a, b int64) int64 {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int64, integers ...int64) int64 {
	result := a * b / GCD(a, b)
	for i := range integers {
		result = LCM(result, integers[i])
	}
	return result
}

func GetInt64Fields(ss []string) [][]int64 {
	var res [][]int64
	for _, s := range ss {
		res = append(res, StrSliceToInt64Slice(strings.Fields(s)))
	}

	return res
}

func StrSliceToInt64Slice(ss []string) []int64 {
	var res []int64
	for _, s := range ss {
		i, _ := strconv.ParseInt(s, 10, 64)
		res = append(res, i)
	}
	return res
}

func PrintSlice[T any](s []T) {
	for _, item := range s {
		fmt.Println(item)
	}
}

func LastElement[T any](s []T) T {
	return s[len(s)-1]
}

func ReverseSlice[T any](s []T) []T {
	var res []T
	for i := len(s) - 1; i >= 0; i-- {
		res = append(res, s[i])
	}
	return res
}

// GRID UTILS

type Coord struct {
	X int
	Y int
}

func GetDirectNeighbourCoords[T any](x, y int, grid [][]T) []Coord {
	var res []Coord

	if x > 0 {
		res = append(res, Coord{x - 1, y})
	}
	if y > 0 {
		res = append(res, Coord{x, y - 1})
	}
	if x < len(grid[0])-1 {
		res = append(res, Coord{x + 1, y})
	}

	if y < len(grid)-1 {
		res = append(res, Coord{x, y + 1})
	}

	return res
}

func InSlice[T comparable](item T, slice []T) bool {
	for _, candidate := range slice {
		if candidate == item {
			return true
		}
	}
	return false
}

func RemoveFromSlice[T comparable](item T, slice []T) []T {
	var res []T
	for _, candidate := range slice {
		if candidate != item {
			res = append(res, candidate)
		}
	}
	return res
}
