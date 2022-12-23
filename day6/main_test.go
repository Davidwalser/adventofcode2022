package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindFirstXSeriesOfDifferentCharacters_TooBigX_ThrowsError(t *testing.T) {
	testString := "mjqjpqmgbljsphdztnvjfqwrcgsmlb"

	result, err := findFirstSeriesOfDifferentCharacters(testString, 31)

	assert.ErrorContains(t, err, "series cannot be greater than length of string value")
	assert.Equal(t, -1, result)
}

func TestFindFirstXSeriesOfDifferentCharacters_NotFindingSeries_Returns0(t *testing.T) {
	testString := "aaaaaaaaa"

	result, err := findFirstSeriesOfDifferentCharacters(testString, 2)

	assert.NoError(t, err)
	assert.Equal(t, 0, result)
}

func TestFindFirstXSeriesOfDifferentCharacters_SeriesOf1_Returns1(t *testing.T) {
	testString := "abc"

	result, err := findFirstSeriesOfDifferentCharacters(testString, 1)

	assert.NoError(t, err)
	assert.Equal(t, 1, result)
}

func TestFindFirstXSeriesOfDifferentCharacters_SeriesOf2_Returns3(t *testing.T) {
	testString := "abc"

	result, err := findFirstSeriesOfDifferentCharacters(testString, 2)

	assert.NoError(t, err)
	assert.Equal(t, 2, result)
}

func TestFindFirstXSeriesOfDifferentCharacters_SeriesOf2_Returns4(t *testing.T) {
	testString := "aabc"

	result, err := findFirstSeriesOfDifferentCharacters(testString, 2)

	assert.NoError(t, err)
	assert.Equal(t, 3, result)
}

func TestIsCharacterDuplicatedInString_NoDuplicates_ShouldReturnFalse(t *testing.T) {
	testString := "abcdef"

	result := isStringDuplicatedInString("a", testString)

	assert.False(t, result)
}

func TestIsCharacterDuplicatedInString_Duplicates_ShouldReturnTrue(t *testing.T) {
	testString := "abcdeaf"

	result := isStringDuplicatedInString("a", testString)

	assert.True(t, result)
}

func TestFindFirstXSeriesOfDifferentCharacters_SeriesOf3_Returns4(t *testing.T) {
	testString := "abcdef"

	result, err := findFirstSeriesOfDifferentCharacters(testString, 3)

	assert.NoError(t, err)
	assert.Equal(t, 3, result)
}

func TestMapContainsValue_DoesntContainValue_ReturnsFalse(t *testing.T) {
	testMap := map[string]bool{
		"Bla":  true,
		"Blub": true,
	}

	result := containsValue(testMap, false)

	assert.False(t, result)
}

func TestMapContainsValue_ContainsValue_ReturnsTrue(t *testing.T) {
	testMap := map[string]bool{
		"Bla":  true,
		"Blub": false,
	}

	result := containsValue(testMap, false)

	assert.True(t, result)
}

func TestFindFirstXSeriesOfDifferentCharacters_SeriesOf4_Returns7(t *testing.T) {
	// 123456789
	testString := "mjqjpqmgbljsphdztnvjfqwrcgsmlb"

	result, err := findFirstSeriesOfDifferentCharacters(testString, 4)

	assert.NoError(t, err)
	assert.Equal(t, 7, result)
}

func TestFindFirstXSeriesOfDifferentCharacters_SeriesOf14_Returns19(t *testing.T) {
	// 123456789
	testString := "mjqjpqmgbljsphdztnvjfqwrcgsmlb"

	result, err := findFirstSeriesOfDifferentCharacters(testString, 14)

	assert.NoError(t, err)
	assert.Equal(t, 19, result)
}
