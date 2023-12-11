package main

import (
	"aoc_2023_go/util"
	"fmt"
	"regexp"
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

type Mapping struct {
	m []Map
}

type Map struct {
	target int64
	source int64
	size   int64
}

type Range struct {
	start int64
	end   int64
}

func one(lines []string) {
	seeds, mappings := processSeedNumbers(lines[0]), processMappings(lines)
	res := findLowestLocationNumber(seeds, mappings)

	fmt.Printf("answer: %d\n", res)
}

func two(lines []string) {
	seeds, mappings := processSeedRanges(lines[0]), processMappings(lines)
	res := findLowestLocationNumberFromRange(seeds, mappings)

	fmt.Printf("answer: %d\n", res)
}

func processSeedNumbers(line string) []int64 {
	rawSeeds, _ := strings.CutPrefix(line, "seeds: ")
	seeds := strings.Split(rawSeeds, " ")
	var seedNumbers []int64

	for _, seed := range seeds {
		seedNumbers = append(seedNumbers, strToInt64(seed))
	}

	return seedNumbers
}

func processSeedRanges(line string) []Range {
	seedNumbers := processSeedNumbers(line)
	var res []Range
	for i := 0; i < len(seedNumbers); {
		res = append(res, Range{seedNumbers[i], seedNumbers[i] + seedNumbers[i+1] - 1})
		i += 2
	}
	return res
}

func processMappings(lines []string) []Mapping {
	var res []Mapping
	var temp []Map
	for _, line := range lines {
		if len(line) > 3 && isMap(line) {
			temp = append(temp, toMap(line))
		} else {
			if temp != nil {
				res = append(res, Mapping{temp})
				temp = nil
			}
		}
	}
	res = append(res, Mapping{temp})
	return res
}

func isMap(line string) bool {
	re := regexp.MustCompile("^[0-9 ]*$")
	return re.MatchString(line)
}

func toMap(line string) Map {
	split := strings.Split(line, " ")
	return Map{
		strToInt64(split[0]),
		strToInt64(split[1]),
		strToInt64(split[2]),
	}
}

func strToInt64(s string) int64 {
	res, _ := strconv.ParseInt(s, 10, 64)
	return res
}

func findLowestLocationNumber(seeds []int64, mappings []Mapping) int64 {
	ans := int64(999999999999)

	for _, seed := range seeds {
		locationNumber := findLocationNumber(seed, mappings)
		if locationNumber < ans {
			ans = locationNumber
		}
	}
	return ans
}

func findLocationNumber(seed int64, mappings []Mapping) int64 {
	original := seed
	ans := seed
	for _, mapping := range mappings {
		for _, m := range mapping.m {
			if original >= m.source && original <= m.source+m.size {
				ans = original + m.target - m.source
			}
		}
		original = ans
	}
	return ans
}

func findLowestLocationNumberFromRange(seeds []Range, mappings []Mapping) int64 {
	min := int64(9223372036854775807)
	var locationNumbers []int64
	for _, seed := range seeds {
		locationNumbers = append(locationNumbers, findLocationNumberFromRange(seed, mappings))
	}

	for _, l := range locationNumbers {
		if l < min {
			min = l
		}
	}

	return min
}

func findLocationNumberFromRange(seed Range, mappings []Mapping) int64 {
	ranges := []Range{seed}
	for _, mapping := range mappings {
		inputRanges := ranges
		ranges = nil
		for _, r := range inputRanges {
			newRanges := findRanges(r, mapping.m, 0)
			ranges = append(ranges, newRanges...)
		}
	}
	return minLocationNumber(ranges)
}

func minLocationNumber(ranges []Range) int64 {
	min := int64(9223372036854775807)
	for _, r := range ranges {
		if r.start < min {
			min = r.start
		}
	}
	return min
}

func dontOverlap(r1, r2 Range) bool {
	return r1.start > r2.end || r2.start > r1.end
}

func findRanges(r Range, mappings []Map, i int) []Range {
	var res []Range
	if i >= len(mappings) {
		res = append(res, r)
		return res
	}
	m := mappings[i]
	r2 := Range{m.source, m.source + m.size - 1}
	if dontOverlap(r, r2) {
		return append(res, findRanges(r, mappings, i+1)...)
	}

	if r.start < r2.start {
		res = append(res, findRanges(Range{r.start, r2.start - 1}, mappings, i+1)...)
	}

	if r.end > r2.end {
		res = append(res, findRanges(Range{r2.end + 1, r.end}, mappings, i+1)...)
	}

	overlapStart := max(r.start, r2.start) + m.target - m.source
	overlapEnd := min(r.end, r2.end) + m.target - m.source

	res = append(res, Range{overlapStart, overlapEnd})
	return res
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
