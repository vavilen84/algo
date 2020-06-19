package selection_sort

// panics if input has zero length
func FindSmallestValuePosition(input []int) int {
	smallestValuePosition := 0
	smallestValue := input[smallestValuePosition]
	for k, v := range input {
		if v < smallestValue {
			smallestValue = v
			smallestValuePosition = k
		}
	}
	return smallestValuePosition
}

func Sort(input []int) []int {
	l := len(input)
	if l == 0 {
		return input
	}
	sorted := make([]int, l)
	for i := 0; i < l; i++ {
		p := FindSmallestValuePosition(input)
		sorted[i] = input[p]
		input = append(input[:p], input[p+1:]...)
	}
	return sorted
}
