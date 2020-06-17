package binary_search

import "sort"

func SearchInIntSlice(haystack []int, needle int) (result bool, iterationsCount int) {
	sort.Ints(haystack)
	lowKey := 0
	highKey := len(haystack) - 1
	if (haystack[lowKey] > needle) || (haystack[highKey] < needle) {
		return
	}
	for lowKey <= highKey {
		iterationsCount++
		mid := (lowKey + highKey) / 2
		if haystack[mid] == needle {
			result = true
			return
		}
		if haystack[mid] < needle {
			lowKey = mid + 1
			continue
		}
		if haystack[mid] > needle {
			highKey = mid - 1
		}
	}
	return
}
