package main

import (
	"strings"
	"unicode"

	"github.com/Davidwalser/adventofcode2022/utils"
)

func getContainers(containersAndRows string) Containers {
	containersAndRowLines := utils.SplitStringInLines(containersAndRows)
	rowLines := containersAndRowLines[len(containersAndRowLines)-1]
	containerLines := containersAndRowLines[:len(containersAndRowLines)-1]
	rowMapper := make(map[int]int)
	for index, row := range rowLines {
		if !unicode.IsSpace(row) {
			rowMapper[index] = int(row - '0')
		}
	}
	containers := make(Containers)
	for i := len(containerLines) - 1; i >= 0; i-- {
		for index, char := range containerLines[i] {
			if !unicode.IsSpace(char) && string(char) != "[" && string(char) != "]" {
				row := rowMapper[index]
				containers[row] = append(containers[row], char)
			}
		}
	}
	return containers
}

type Move struct {
	count   int
	fromRow int
	toRow   int
}

type Containers map[int][]rune

func getMoves(moveAsText string) []Move {
	lines := utils.SplitStringInLines(strings.TrimSuffix(moveAsText, "\n"))
	var moves []Move
	for _, moveText := range lines {
		splittedText := strings.Split(moveText, " ")
		moves = append(moves, Move{utils.ParseToInt(splittedText[1]), utils.ParseToInt(splittedText[3]), utils.ParseToInt(splittedText[5])})
	}
	return moves
}

func push(containers []rune, container rune) []rune {
	return append(containers, container)
}

func pop(containers []rune) []rune {
	return containers[:len(containers)-1]
}

func copyContainerMap(value Containers) Containers {
	newMap := make(Containers)
	for k, v := range value {
		newMap[k] = v
	}
	return newMap
}

func moveContainersPart1(containers Containers, moves []Move) Containers {
	orderedContainers := copyContainerMap(containers)
	for _, move := range moves {
		for i := 1; i <= move.count; i++ {
			containerToMove := orderedContainers[move.fromRow][len(orderedContainers[move.fromRow])-1]
			pop(orderedContainers[move.fromRow])
			push(orderedContainers[move.toRow], containerToMove)
		}
	}
	return orderedContainers
}

func moveContainersPart2(containers Containers, moves []Move) Containers {
	orderedContainers := copyContainerMap(containers)
	for _, move := range moves {
		containersToMove := orderedContainers[move.fromRow][len(orderedContainers[move.fromRow])-move.count:]
		orderedContainers[move.fromRow] = orderedContainers[move.fromRow][:len(orderedContainers[move.fromRow])-move.count]
		orderedContainers[move.toRow] = append(orderedContainers[move.toRow], containersToMove...)
	}
	return orderedContainers
}

func getLastContainers(containers Containers) string {
	var result []rune
	for i := 1; i <= len(containers); i++ {
		result = append(result, containers[i][len(containers[i])-1])
	}
	return string(result)
}

func main() {
	file := utils.ReadFile("input.txt")
	splittedFile := strings.Split(file, "\n\n")
	containersAndRows := splittedFile[0]
	movesAsText := splittedFile[1]
	moves := getMoves(movesAsText)
	containers := getContainers(containersAndRows)

	part1Containers := moveContainersPart1(containers, moves)
	part1 := getLastContainers(part1Containers)
	println("PART1:", part1)

	part2Containers := moveContainersPart2(containers, moves)
	part2 := getLastContainers(part2Containers)
	println("PART2:", part2)
}
