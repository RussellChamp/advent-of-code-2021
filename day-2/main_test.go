package main

import (
	assert "AoC2021/utils/testing"
	"testing"
)

func TestDay2(t *testing.T) {
	assert.IsTrue(t, true, "cool")
	assert.StrEquals(t, "foo", "bar", "baz")
}