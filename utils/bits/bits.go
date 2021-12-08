package bits

func GetBitCount(val int) int {
	count := 0
	for val != 0 {
		count += val & 1
		val >>= 1
	}

	return count
}

// returns the number of matching bits in the two value's binary expression
func MatchingOnBits(v1 int, v2 int) int {
	count := 0
	for v1 != 0 && v2 != 0 {
		if v1&1 == 1 && v2&1 == 1 {
			count += 1
		}
		v1 >>= 1
		v2 >>= 1
	}

	return count
}

func BitwiseAnd(v1 int, v2 int) int {
	val := 0
	for block := 1; v1 != 0 || v2 != 0; block <<= 1 {
		if v1&1 == 1 || v2&1 == 1 {
			val += block
		}
		if v1 > 0 {
			v1 >>= 1
		}
		if v2 > 0 {
			v2 >>= 1
		}
	}
	return val
}

func BitwiseAndA(a ...int) int {
	total := 0
	for _, i := range a {
		total = BitwiseAnd(total, i)
	}

	return total
}

func ByBitCount(size int) func(int) bool {
	return func(val int) bool { return GetBitCount(val) == size }
}
