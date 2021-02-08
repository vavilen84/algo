package merge_sort

func MergeSort(input []int) []int {
	l := len(input)
	if l == 1 {
		// length validation
		return input
	}
	// divide slice into 2 pieces
	middleIdx := l / 2
	left := input[:middleIdx]
	right := input[middleIdx:]

	// recursive merge
	return Merge(MergeSort(left), MergeSort(right))
}

func Merge(left, right []int) []int {
	result := make([]int, len(left)+len(right))
	i := 0
	// create new slice using key-by-key comparison left & right slices
	for (len(left) > 0) && (len(right) > 0) {
		if left[0] < right[0] {
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
			right = right[1:]
		}
		i++
	}
	for j := 0; j < len(left); j++ {
		result[i] = left[j]
	}
	for j := 0; j < len(right); j++ {
		result[i] = right[j]
	}
	return result
}
