package selection_sort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindSmallestValuePosition(t *testing.T) {
	input := []int{5, 3, 6, 2, 10}
	position := FindSmallestValuePosition(input)
	assert.Equal(t, 3, position)
}

func TestSort(t *testing.T) {
	input := []int{5, 3, 6, 2, 10}
	sorted := Sort(input)
	assert.Equal(t, 2, sorted[0])
	assert.Equal(t, 3, sorted[1])
	assert.Equal(t, 5, sorted[2])
	assert.Equal(t, 6, sorted[3])
	assert.Equal(t, 10, sorted[4])

	// test zero length input
	input = []int{}
	sorted = Sort(input)
	assert.Equal(t, len(input), len(sorted))
}
