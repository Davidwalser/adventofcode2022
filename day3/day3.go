package main

import (
	"adventofcode/m/utils"
	"fmt"
)

func splitInHalf(str string) (string, string) {
	sizeOfStr := len(str)
	var firstHalf []rune
	var secondHalf []rune
	for i := 0; i < sizeOfStr/2; i++ {
		firstHalf = append(firstHalf, rune(str[i]))
	}
	for i := sizeOfStr / 2; i < sizeOfStr; i++ {
		secondHalf = append(secondHalf, rune(str[i]))
	}
	return string(firstHalf), string(secondHalf)
}

func findDuplicates(firstHalf string, secondHalf string) rune {
	for _, firstChar := range firstHalf {
		for _, secondChar := range secondHalf {
			if firstChar == secondChar {
				return firstChar
			}
		}
	}
	panic("No duplicates found")
}

func initializeMap() map[rune]int {
	letterToNumber := make(map[rune]int)
	index := 1
	for letter := 'a'; letter <= 'z'; letter++ {
		letterToNumber[letter] = index
		index++
	}
	for letter := 'A'; letter <= 'Z'; letter++ {
		letterToNumber[letter] = index
		index++
	}
	return letterToNumber
}

func part1(lines []string) int {
	var runeToIntMap map[rune]int = initializeMap()
	points := 0
	for _, eachline := range lines {
		firstHalf, secondHalf := splitInHalf(eachline)
		points += runeToIntMap[findDuplicates(firstHalf, secondHalf)]
	}
	return points
}

func findDuplicatesInLines(firstLine string, secondLine string, thirdLine string) rune {
	for _, firstChar := range firstLine {
		for _, secondChar := range secondLine {
			for _, thridChar := range thirdLine {
				if firstChar == secondChar && secondChar == thridChar {
					return firstChar
				}
			}
		}
	}
	panic("No duplicates found")
}

func part2(lines []string) int {
	var runeToIntMap map[rune]int = initializeMap()
	points := 0
	chunkSize := 3
	for i := 0; i < len(lines); i += chunkSize {
		end := i + chunkSize
		threeLines := lines[i:end]
		points += runeToIntMap[findDuplicatesInLines(threeLines[0], threeLines[1], threeLines[2])]
	}
	return points
}

func main() {
	lines := utils.ReadLines("day3_input.txt")
	part1Points := part1(lines)
	fmt.Println("TOTAL POINTS PART1: ", part1Points)
	part2Points := part2(lines)
	fmt.Println("TOTAL POINTS PART2: ", part2Points)
}
