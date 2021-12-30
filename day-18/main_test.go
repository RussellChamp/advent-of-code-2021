package main

import (
	"AoC2021/utils/assert"
	"AoC2021/utils/log"
	"testing"
)

func TestDay18_ParseSnumStr(t *testing.T) {
	s1 := ParseSnumStr("[1,2]")
	s2 := ParseSnumStr("[11,22]")

	assert.EqualsInt(t, s1.lVal, 1, "should parse left value")
	assert.EqualsInt(t, s1.rVal, 2, "should parse right value")

	assert.EqualsInt(t, s2.lVal, 11, "should parse larger left value")
	assert.EqualsInt(t, s2.rVal, 22, "should parse larger right value")
}

func TestDay18_ToString(t *testing.T) {
	s := Snum{lVal: 1, rVal: 2}
	assert.EqualsStr(t, s.ToString(), "[1,2]", "should convert to string")
}

func TestDay18_AddSnum(t *testing.T) {
	s1 := Snum{lVal: 1, rVal: 2}
	s2 := Snum{lVal: 2, rVal: 3}
	s3 := AddSnum(s1, s2)

	assert.EqualsStr(t, s3.ToString(), "[[1,2],[2,3]]", "should add snums")
}

func TestDay18_Add(t *testing.T) {
	log.SetLogLevel(log.DIAGNOSTIC)
	s := Snum{lVal: 1, rVal: 1}
	s = s.Add(Snum{lVal: 2, rVal: 2})
	s = s.Add(Snum{lVal: 3, rVal: 3})
	s = s.Add(Snum{lVal: 4, rVal: 4})

	assert.EqualsStr(t, s.ToString(), "[[[[1,1],[2,2]],[3,3]],[4,4]]", "should add snums")

	s = s.Add(Snum{lVal: 5, rVal: 5})
	assert.EqualsStr(t, s.ToString(), "[[[[3,0],[5,3]],[4,4]],[5,5]]", "should simplify after adding snums")

	s = s.Add(Snum{lVal: 6, rVal: 6})
	assert.EqualsStr(t, s.ToString(), "[[[[5,0],[7,4]],[5,5]],[6,6]]", "should simplify after adding snums")
}

func TestDay18_Simplify(t *testing.T) {
	s := ParseSnumStr("[0,[[[1,[2,3]],4],5]]")

	s = s.Simplify()

	assert.EqualsStr(t, s.ToString(), "[0,[[[3,0],7],5]]", "should convert to string")
}

func TestDay18_IsNormal(t *testing.T) {
	s1 := ParseSnumStr("[0,1]")
	s2 := ParseSnumStr("[0,[[[1,[2,3]],4],5]]")

	assert.IsTrue(t, s1.IsNormal(), "should be a normal snum")
	assert.IsFalse(t, s2.IsNormal(), "should NOT be a normal snum")
}

func TestDay18_ShouldExplode(t *testing.T) {
	s1 := ParseSnumStr("[0,1]")
	s2 := ParseSnumStr("[0,[[[1,[2,3]],4],5]]")

	assert.IsFalse(t, s1.ShouldExplode(3), "should NOT explode a shallow normal snum")
	assert.IsTrue(t, s1.ShouldExplode(6), "should explode a deep normal snum")
	assert.IsFalse(t, s2.ShouldExplode(3), "should NOT explode a shallow nested snum")
	assert.IsFalse(t, s2.ShouldExplode(9), "should NOT explode a deep nested snum")
}

// func TestDay18_ExplodeValue(t *testing.T) {
// 	// skip
// }

func TestDay18_ShouldSplit(t *testing.T) {
	s1 := ParseSnumStr("[0,1]")
	s2 := ParseSnumStr("[11,1]")
	s3 := ParseSnumStr("[[1,2],13]")

	assert.IsFalse(t, s1.ShouldSplit(), "should NOT split a small normal snum")
	assert.IsTrue(t, s2.ShouldSplit(), "should split a large normal snum")
	assert.IsTrue(t, s3.ShouldSplit(), "should split a large nested snum")
}

func TestDay18_CalcMagnitude(t *testing.T) {
	s1 := ParseSnumStr("[3,1]")
	s2 := ParseSnumStr("[1,[2,3]]")

	assert.EqualsInt(t, s1.CalcMagnitude(), 11, "should calc normal snum")
	assert.EqualsInt(t, s2.CalcMagnitude(), 27, "should calc nested snum")
}
