package util

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func ReadInputLines(subdirectory, fileName string) []string {
	path := filepath.Join(subdirectory, fileName)
	readFile, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	var lines []string

	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}

	readFile.Close()
	return lines
}

func ReadFileLinesFromSubdirectory(subdirectory, filename string) ([]string, error) {
	path := filepath.Join(subdirectory, filename)
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
