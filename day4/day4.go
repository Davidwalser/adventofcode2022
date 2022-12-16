package main

import (
	"fmt"
	"strings"

	"github.com/Davidwalser/adventofcode2022/utils"
)

type Pair struct {
	a Range
	b Range
}

type Range struct {
	start int
	end   int
}

func getRange(text string) Range {
	numbers := strings.Split(text, "-")
	start := utils.ParseToInt(numbers[0])
	end := utils.ParseToInt(numbers[1])
	return Range{start, end}
}

func getRanges(textLines []string) []Pair {
	var returnRanges []Pair
	for _, line := range textLines {
		ranges := strings.Split(line, ",")
		firstRange := getRange(ranges[0])
		secondRange := getRange(ranges[1])
		returnRanges = append(returnRanges, Pair{firstRange, secondRange})
	}
	return returnRanges
}

func fullyOverlap(pair Pair) bool {
	if (pair.a.start >= pair.b.start && pair.a.end <= pair.b.end) || (pair.a.start <= pair.b.start && pair.a.end >= pair.b.end) {
		return true
	}
	return false
}

func countFullyOverlap(pairs []Pair) int {
	countFullyOverlap := 0
	for _, pair := range pairs {
		if fullyOverlap(pair) {
			countFullyOverlap++
		}
	}
	return countFullyOverlap
}

func overlap(pair Pair) bool {
	if fullyOverlap(pair) || (pair.a.start <= pair.b.end && pair.a.end >= pair.b.start) || (pair.b.start <= pair.a.end && pair.b.end >= pair.a.start) {
		return true
	}
	return false
}

func countOverlap(pairs []Pair) int {
	countOverlap := 0
	for _, pair := range pairs {
		if overlap(pair) {
			countOverlap++
		}
	}
	return countOverlap
}

func main() {
	lines := utils.ReadLines("day4_input.txt")
	ranges := getRanges(lines)
	fmt.Println("PART1 Fully overlap: ", countFullyOverlap(ranges))
	fmt.Println("PART2 overlap: ", countOverlap(ranges))
}
