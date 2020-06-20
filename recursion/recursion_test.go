package recursion

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSum(t *testing.T) {
	input := []int{}
	r := Sum(input)
	assert.Equal(t, 0, r)

	input = []int{2}
	r = Sum(input)
	assert.Equal(t, 2, r)

	input = []int{0, 1, 2}
	r = Sum(input)
	assert.Equal(t, 3, r)
}
