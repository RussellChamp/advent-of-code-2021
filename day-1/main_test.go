package main

import (
	assert "AoC2021/utils/testing"
	"testing"
)

func TestPrintPip(t *testing.T) {
	DEBUG = false
	str := GetPip(1, 1, 1, 100)
	assert.StrEquals(t, str, "", "Should return an empty string")

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
		assert.StrEquals(t, str, c.expected, c.msg)
	}

}
