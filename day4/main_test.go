package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRanges(t *testing.T) {
	testString := []string{"2-4,6-8", "2-3,4-5", "5-7,7-9", "2-8,3-7", "6-6,4-6", "2-6,4-8"}
	ranges := getRanges(testString)
	assert.Equal(t, Pair{Range{2, 4}, Range{6, 8}}, ranges[0])
}

func TestFullyOverlap(t *testing.T) {
	assert.Equal(t, false, fullyOverlap(Pair{Range{2, 4}, Range{6, 8}}))
	assert.Equal(t, true, fullyOverlap(Pair{Range{2, 10}, Range{6, 8}}))
	assert.Equal(t, false, fullyOverlap(Pair{Range{2, 10}, Range{1, 5}}))
}

func TestCountFullyOverlap(t *testing.T) {
	// testString := []string{"2-4,6-8", "2-3,4-5", "5-7,7-9", "2-8,3-7", "6-6,4-6", "2-6,4-8"}
	ranges := []Pair{{Range{2, 4}, Range{6, 8}}, {Range{2, 3}, Range{4, 5}}, {Range{5, 7}, Range{7, 9}}, {Range{2, 8}, Range{3, 7}}, {Range{6, 6}, Range{4, 6}}, {Range{2, 6}, Range{4, 8}}}
	count := countFullyOverlap(ranges)
	assert.Equal(t, 2, count)
}

func TestOverlap(t *testing.T) {
	assert.Equal(t, false, overlap(Pair{Range{2, 4}, Range{6, 8}}))
	assert.Equal(t, true, overlap(Pair{Range{2, 10}, Range{6, 8}}))
	assert.Equal(t, true, overlap(Pair{Range{2, 10}, Range{1, 5}}))
	assert.Equal(t, true, overlap(Pair{Range{5, 10}, Range{1, 5}}))
	assert.Equal(t, true, overlap(Pair{Range{5, 10}, Range{2, 6}}))
}
