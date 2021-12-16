package bits

import "fmt"

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

func ToInt(bits []bool) int {
	intVal := 0
	for idx := 0; idx < len(bits); idx++ {
		intVal = intVal << 1
		if bits[idx] {
			intVal += 1
		}
	}
	return intVal
}

func ToStr(bits []bool) string {
	str := ""
	for idx := 0; idx < len(bits); idx++ {
		if bits[idx] {
			str += "1"
		} else {
			str += "0"
		}
	}
	return str
}

func HexByteToBits(b byte) ([]bool, error) {
	// just do this manually because i can't be arsed
	switch b {
	case byte('0'):
		return []bool{false, false, false, false}, nil
	case byte('1'):
		return []bool{false, false, false, true}, nil
	case byte('2'):
		return []bool{false, false, true, false}, nil
	case byte('3'):
		return []bool{false, false, true, true}, nil
	case byte('4'):
		return []bool{false, true, false, false}, nil
	case byte('5'):
		return []bool{false, true, false, true}, nil
	case byte('6'):
		return []bool{false, true, true, false}, nil
	case byte('7'):
		return []bool{false, true, true, true}, nil
	case byte('8'):
		return []bool{true, false, false, false}, nil
	case byte('9'):
		return []bool{true, false, false, true}, nil
	case byte('A'):
		return []bool{true, false, true, false}, nil
	case byte('B'):
		return []bool{true, false, true, true}, nil
	case byte('C'):
		return []bool{true, true, false, false}, nil
	case byte('D'):
		return []bool{true, true, false, true}, nil
	case byte('E'):
		return []bool{true, true, true, false}, nil
	case byte('F'):
		return []bool{true, true, true, true}, nil
	default:
		return []bool{}, fmt.Errorf("read invalid byte")
	}
}
