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

func Max(i1 int, a ...int) int {
	max := i1
	for _, i := range a {
		if i > max {
			max = i
		}
	}
	return max
}

func BoolToInt(b bool) int {
	if b {
		return 1
	} else {
		return 0
	}
}
