package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitializeMap(t *testing.T) {
	bla := initializeMap()
	assert.Equal(t, 1, bla['a'], "a should be 1")
	assert.Equal(t, 3, bla['c'], "c should be 3")
	assert.Equal(t, 25, bla['y'], "y should be 25")
	assert.Equal(t, 26, bla['z'], "z should be 26")
	assert.Equal(t, 27, bla['A'], "A should be 27")
	assert.Equal(t, 28, bla['B'], "B should be 28")
	assert.Equal(t, 52, bla['Z'], "Z should be 52")
}

func TestSplitInHalf(t *testing.T) {
	testString := "12345678"
	firstHalf, secondHalf := splitInHalf(testString)
	assert.Equal(t, "1234", firstHalf, "firsthalf should be 1234")
	assert.Equal(t, "5678", secondHalf, "secondhalf should be 5678")
}

func TestPart1(t *testing.T) {
	testString := []string{"vJrwpWtwJgWrhcsFMMfFFhFp", "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL", "PmmdzqPrVvPwwTWBwg", "wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn", "	ttgJtRGJQctTZtZT", "CrZsJsPPZsGzwwsLwLmpwMDw"}
	points := part1(testString)
	assert.Equal(t, 157, points, "points should be 157")
}

func TestPart2(t *testing.T) {
	testString := []string{"vJrwpWtwJgWrhcsFMMfFFhFp", "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL", "PmmdzqPrVvPwwTWBwg", "wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn", "	ttgJtRGJQctTZtZT", "CrZsJsPPZsGzwwsLwLmpwMDw"}
	points := part2(testString)
	assert.Equal(t, 70, points, "points should be 70")
}
