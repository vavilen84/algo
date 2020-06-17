package binary_search

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

type IntSliceAssertItem struct {
	Result          bool
	IterationsCount int
	ItemsCount      int
	Needle          int
}

var intSliceAssertData = []IntSliceAssertItem{
	{
		Result:          true,
		IterationsCount: 8,
		ItemsCount:      128,
		Needle:          128,
	},
	{
		Result:          true,
		IterationsCount: 7,
		ItemsCount:      100,
		Needle:          100,
	},
	{
		Result:          false,
		IterationsCount: 0,
		ItemsCount:      10,
		Needle:          11,
	},
	{
		Result:          false,
		IterationsCount: 0,
		ItemsCount:      10,
		Needle:          0,
	},
}

func getIntSlice(count int) (result []int) {
	for i := 1; i <= count; i++ {
		result = append(result, i)
	}
	rand.Shuffle(len(result), func(i, j int) {
		result[i], result[j] = result[j], result[i]
	})
	return
}

func TestSearchInIntSlice(t *testing.T) {
	for _, v := range intSliceAssertData {
		haystack := getIntSlice(v.ItemsCount)
		result, iterationsCount := SearchInIntSlice(haystack, v.Needle)
		assert.Equal(t, v.Result, result)
		assert.Equal(t, v.IterationsCount, iterationsCount)
	}
}
