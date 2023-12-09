package util

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
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
