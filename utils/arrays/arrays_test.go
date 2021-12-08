package arrays

import (
	"AoC2021/utils/assert"
	"testing"
)

func TestFindFirst(t *testing.T) {
	list := []string{"foo", "bar", "baz"}
	filter := func(s string) bool { return s == "baz" }
	item := FindFirst(list, filter)

	assert.EqualsStr(t, item, "baz", "should find item baz")
}

func TestFindFirstInt(t *testing.T) {
	list := []int{1, 3, 4, 5}
	filter := func(i int) bool { return i%2 == 0 }
	item := FindFirstInt(list, filter)

	assert.EqualsInt(t, item, 4, "should find even number")
}

func TestFindAll(t *testing.T) {
	list := []string{"hello", "world", "foo"}
	filter := func(s string) bool { return len(s) > 3 }
	items := FindAll(list, filter)

	assert.EqualsInt(t, len(items), 2, "should filter lists of strings")
}

func TestFindAllInt(t *testing.T) {
	list := []int{1, 3, 99, 200, 9000}
	filter := func(i int) bool { return i > 50 }
	items := FindAllInt(list, filter)

	assert.EqualsInt(t, len(items), 3, "should filter lists of ints")
}

func TestMapStrToInt(t *testing.T) {
	list := []string{"alice", "bob", "carol"}
	mapFn := func(s string) int { return int(s[0]) }

	items := MapStrToInt(list, mapFn)

	assert.EqualsInt(t, items[0], int('a'), "should run map function across all strings")
}
