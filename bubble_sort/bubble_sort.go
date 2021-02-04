package bubble_sort

func SortRecursive(input []int, i int, shifted bool, preLastIndex, iterationsCount int) ([]int, int) {
	iterationsCount++ // sort iterations count
	if i == 0 {
		shifted = false // reset shifted flag for each index reset
	}
	if input[i] > input[i+1] { // compare two values
		input[i], input[i+1] = input[i+1], input[i] // swap values if left is bigger
		// handle pre-last index
		if preLastIndex == i {
			// if there was no shift except last 2 digits - so we change last two values and return result
			if shifted == false {
				return input, iterationsCount
			}
		}
		shifted = true
	} else {
		// handle pre-last index
		if preLastIndex == i {
			// if there wes no shift and last two digits are sorted - so we return result
			if shifted == false {
				return input, iterationsCount
			}
		}
	}
	// handle pre-last index for both cases (shift/no shift)
	if preLastIndex == i {
		i = 0 // if pre-last index - reset index counter
	} else {
		i++
	}
	return SortRecursive(input, i, shifted, preLastIndex, iterationsCount)
}
