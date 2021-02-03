package binary_search

import "sort"

func SearchInIntSlice(haystack []int, needle int) (result bool, iterationsCount int) {
	sort.Ints(haystack) // this algorithm will not work with not sorted list

	lowKey := 0                  // first index
	highKey := len(haystack) - 1 // last index

	if (haystack[lowKey] > needle) || (haystack[highKey] < needle) {
		return // target value is not in the range
	}

	for lowKey <= highKey { // iteratively reduce list

		iterationsCount++

		mid := (lowKey + highKey) / 2 // middle index

		if haystack[mid] == needle {
			result = true // we found our value
			return
		}
		if haystack[mid] < needle { // if needle is bigger than middle - we take only part with highest values by increasing lowKey
			lowKey = mid + 1
			continue
		}
		if haystack[mid] > needle { // if needle is smaller than middle - we take only part with lowest values by decreasing highKey
			highKey = mid - 1
		}
	}
	return
}
