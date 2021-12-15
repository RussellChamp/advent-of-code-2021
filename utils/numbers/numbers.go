package numbers

func Min(i1 int, a ...int) int {
	min := i1
	for _, i := range a {
		if i < min {
			min = i
		}
	}
	return min
}
