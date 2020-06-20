package recursion

func Sum(i []int) int {
	l := len(i)
	if l < 2 {
		if l == 0 {
			return 0
		}
		return i[0]
	}
	return i[0] + Sum(i[1:])
}

func FindBiggest(i []int, biggest int) int {
	l := len(i)
	switch l {
	case 0:
		return 0
	case 1:
		if i[0] > biggest {
			return i[0]
		}
		return biggest
	default:
		biggest = FindBiggest(i[1:], biggest)
		if i[0] > biggest {
			return i[0]
		}
		return biggest
	}
}
