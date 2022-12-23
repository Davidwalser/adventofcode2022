package main

import (
	"errors"
	"strings"

	"github.com/Davidwalser/adventofcode2022/utils"
)

func isStringDuplicatedInString(stringToFind string, stringText string) bool {
	return strings.Count(stringText, stringToFind) > 1
}

func containsValue[M ~map[K]V, K, V comparable](mapToCheck M, value V) bool {
	for _, x := range mapToCheck {
		if x == value {
			return true
		}
	}
	return false
}

func findFirstSeriesOfDifferentCharacters(value string, series int) (int, error) {
	if series > len(value) {
		return -1, errors.New("series cannot be greater than length of string value")
	}

	for i := 0; i < len(value)-series; i++ {
		foundDuplicates := make(map[int]bool)
		for j := i; j < i+series; j++ {
			foundDuplicates[j] = isStringDuplicatedInString(string(value[j]), value[i:i+series])
		}
		if !containsValue(foundDuplicates, true) {
			return i + series, nil
		}
	}
	return 0, nil
}

func main() {
	input := utils.ReadFile("input.txt")
	part1FoundDuplicates, err := findFirstSeriesOfDifferentCharacters(input, 4)
	utils.Check(err)
	println("Part1:", part1FoundDuplicates)
	part2FoundDuplicates, err := findFirstSeriesOfDifferentCharacters(input, 14)
	utils.Check(err)
	println("Part2:", part2FoundDuplicates)
}
