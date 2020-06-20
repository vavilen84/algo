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
