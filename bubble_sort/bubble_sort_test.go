package bubble_sort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSort(t *testing.T) {
	input := []int{2, 1, 3}
	preLastIndex := len(input) - 2
	input, iCount := SortRecursive(input, 0, false, preLastIndex, 0)
	assert.Equal(t, 4, iCount)
	assert.Equal(t, 1, input[0])
	assert.Equal(t, 2, input[1])
	assert.Equal(t, 3, input[2])

	input = []int{2, 1, 3, 7, 4, 6, 5}
	preLastIndex = len(input) - 2
	input, iCount = SortRecursive(input, 0, false, preLastIndex, 0)
	assert.Equal(t, 18, iCount)
	assert.Equal(t, 1, input[0])
	assert.Equal(t, 2, input[1])
	assert.Equal(t, 3, input[2])
	assert.Equal(t, 4, input[3])
	assert.Equal(t, 5, input[4])
	assert.Equal(t, 6, input[5])
	assert.Equal(t, 7, input[6])
}
