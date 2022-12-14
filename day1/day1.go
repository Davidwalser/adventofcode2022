package main

import (
	"fmt"
	"sort"

	"github.com/Davidwalser/adventofcode2022/utils"
)

func sortDescending(list []int) []int {
	sort.Slice(list, func(i, j int) bool {
		return list[i] > list[j]
	})
	return list
}

func sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func calculateWeights(textLines []string) []int {
	var weights []int
	count := 0
	for _, eachline := range textLines {
		if eachline == "" {
			weights = append(weights, count)
			count = 0
		} else {
			var num int
			_, err := fmt.Sscan(eachline, &num)
			utils.Check(err)
			count += num
		}
	}
	return weights
}

func main() {
	textLines := utils.ReadLines("day1_input.txt")
	weights := calculateWeights(textLines)
	var sortedWeights = sortDescending(weights)
	fmt.Println("PART1:", sortedWeights[0])
	fmt.Println("PART2:", sum(sortedWeights[0:3]))
}
