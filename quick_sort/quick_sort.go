package quick_sort

func Sort(input []int) []int {
	l := len(input)
	if l < 2 {
		return input
	}
	less := make([]int, 0)
	bigger := make([]int, 0)
	pivot := input[0]
	for _, v := range input[1:] {
		if v > pivot {
			bigger = append(bigger, v)
		} else {
			less = append(less, v)
		}
	}
	input = append(Sort(less), pivot)
	input = append(input, Sort(bigger)...)
	return input
}
