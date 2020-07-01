package merge_sort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMergeSort(t *testing.T) {
	input := []int{2, 1, 3}
	output := MergeSort(input)
	assert.Equal(t, 1, output[0])
	assert.Equal(t, 2, output[1])
	assert.Equal(t, 3, output[2])
}
