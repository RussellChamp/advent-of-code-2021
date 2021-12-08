package main

import (
	"AoC2021/utils/assert"
	"testing"
)

func TestPrintPip(t *testing.T) {
	DEBUG = false
	str := GetPip(1, 1, 1, 100)
	assert.EqualsStr(t, str, "", "Should return an empty string")

	DEBUG = true
	cases := []struct {
		first, previous, current, total int
		expected, msg                   string
	}{
		{-1, -1, -1, -1, ".", ""},
		{-1, 0, 0, 42, ".", ""},
		{100, 100, 200, 42, "+", ""},
		{100, 200, 100, 42, "-", ""},
		{100, 100, 100, 100, ".\n", "should include newline"},
	}

	for _, c := range cases {
		str := GetPip(c.first, c.previous, c.current, c.total)
		assert.EqualsStr(t, str, c.expected, c.msg)
	}

}
