package bits

import (
	"AoC2021/utils/assert"
	"testing"
)

func TestGetBitCount(t *testing.T) {
	assert.EqualsInt(t, GetBitCount(0), 0, "should count zero bits")
	assert.EqualsInt(t, GetBitCount(15), 4, "should count bits")
	assert.EqualsInt(t, GetBitCount(16), 1, "should not count empty bits")
}
func TestMatchingOnBits(t *testing.T) {
	assert.EqualsInt(t, MatchingOnBits(5 /*101*/, 13 /*1101*/), 2, "should match 2 bits")
}

func TestBitwiseAnd(t *testing.T) {
	assert.EqualsInt(t, BitwiseAnd(7 /*111*/, 18 /*10010*/), 23 /*10111*/, "should bitwise add")
}

func TestBitwiseAndA(t *testing.T) {
	assert.EqualsInt(t, BitwiseAndA(1, 4, 8, 12), 13 /*1101*/, "should bitwise add lists")
}

func TestByBitCount(t *testing.T) {
	twoBitFilter := ByBitCount(2)
	assert.IsTrue(t, twoBitFilter(17 /*10001*/), "should count bits")
}

func TestHexByteToBits(t *testing.T) {
	cases := []struct {
		input  rune
		output string
	}{
		{'0', "0000"},
		{'1', "0001"},
		{'2', "0010"},
		{'3', "0011"},
		{'4', "0100"},
		{'5', "0101"},
		{'6', "0110"},
		{'7', "0111"},
		{'8', "1000"},
		{'9', "1001"},
		{'A', "1010"},
		{'B', "1011"},
		{'C', "1100"},
		{'D', "1101"},
		{'E', "1110"},
		{'F', "1111"},
	}

	for _, c := range cases {
		bits, err := HexByteToBits(byte(c.input))
		if err != nil {
			t.Logf("Got error from HexByteToBits: %s", err.Error())
			t.Fail()
		}
		bitStr := ToStr(bits)
		assert.EqualsStr(t, bitStr, c.output, "incorrect value")
	}
}
