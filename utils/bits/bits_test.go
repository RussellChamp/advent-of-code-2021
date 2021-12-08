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
