package bubble_sort

func Sort(input []int, i int, ok bool, iCount int) ([]int, int) {
	l := len(input)
	if l < 2 {
		return input, iCount
	}
	iCount++
	if i == 0 {
		ok = true
	}
	if input[i] > input[i+1] {
		input[i], input[i+1] = input[i+1], input[i]
		ok = false
	}
	if (l - 3) == (i - 1) {
		if ok {
			return input, iCount
		}
		i = 0
	} else {
		i++
	}
	return Sort(input, i, ok, iCount)
}
