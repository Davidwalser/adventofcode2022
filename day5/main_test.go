package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetContainers(t *testing.T) {
	testString := "    [D]\n[N] [C]\n[Z] [M] [P]\n 1   2   3"

	containers := getContainers(testString)

	assert.Equal(t, "Z", string(containers[1][0]))
	assert.Equal(t, "N", string(containers[1][1]))
	assert.Equal(t, "M", string(containers[2][0]))
	assert.Equal(t, "C", string(containers[2][1]))
	assert.Equal(t, "D", string(containers[2][2]))
	assert.Equal(t, "P", string(containers[3][0]))
}

func TestGetMoves(t *testing.T) {
	testString := "move 1 from 2 to 1\nmove 3 from 1 to 3\nmove 2 from 2 to 1\nmove 1 from 1 to 2\n"

	moves := getMoves(testString)

	assert.Equal(t, Move{1, 2, 1}, moves[0])
	assert.Equal(t, Move{3, 1, 3}, moves[1])
	assert.Equal(t, Move{2, 2, 1}, moves[2])
	assert.Equal(t, Move{1, 1, 2}, moves[3])
}

func TestCopyContainers(t *testing.T) {
	containers := make(map[int][]rune)
	containers[1] = []rune{'A', 'B', 'C'}
	containers[2] = []rune{'D', 'E', 'F'}

	copiedContainers := copyContainerMap(containers)

	assert.Equal(t, containers[1], copiedContainers[1])
	assert.Equal(t, containers[2], copiedContainers[2])
}

func TestCopyContainersHasNoReferenceOnSource(t *testing.T) {
	containers := make(map[int][]rune)
	containers[1] = []rune{'A', 'B', 'C'}
	containers[2] = []rune{'D', 'E', 'F'}

	copiedContainers := copyContainerMap(containers)
	containers[2] = []rune{'G', 'H', 'J', 'K', 'L', 'M'}
	containers[3] = []rune{'O'}

	assert.Equal(t, []rune{'A', 'B', 'C'}, copiedContainers[1])
	assert.Equal(t, []rune{'D', 'E', 'F'}, copiedContainers[2])
}
