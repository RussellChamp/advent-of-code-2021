package main

import (
	"AoC2021/utils/assert"
	"testing"
)

func TestDay8ParseField(t *testing.T) {
	assert.EqualsInt(t, ParseField("abc"), 7, "should parse fields correctly")
}
