package main

import (
	"AoC2021/utils/assert"
	"testing"
)

func TestDay18_ParseSnumStr(t *testing.T) {
	s1 := ParseSnumStr("[1,2]")
	s2 := ParseSnumStr("[11,22]")

	assert.EqualsInt(t, s1[1], 1, "should parse left value")
	assert.EqualsInt(t, s1[3], 2, "should parse right value")

	assert.EqualsInt(t, s2[1], 11, "should parse larger left value")
	assert.EqualsInt(t, s2[3], 22, "should parse larger right value")
}

func TestDay18_ToString(t *testing.T) {
	s := Snum{LBracket, 1, Comma, 2, RBracket}
	assert.EqualsStr(t, s.ToString(), "[1,2]", "should convert to string")
}

func TestDay18_AddSnum(t *testing.T) {
	s1 := Snum{LBracket, 1, Comma, 2, RBracket}
	s2 := Snum{LBracket, 2, Comma, 3, RBracket}
	s3 := AddSnum(s1, s2)

	assert.EqualsStr(t, s3.ToString(), "[[1,2],[2,3]]", "should add snums")
}

func TestDay18_Add(t *testing.T) {
	s := Snum{LBracket, 1, Comma, 1, RBracket}
	s = s.Add(Snum{LBracket, 2, Comma, 2, RBracket})
	assert.EqualsStr(t, s.ToString(), "[[1,1],[2,2]]", "should add snums")

	s = s.Add(Snum{LBracket, 3, Comma, 3, RBracket})
	assert.EqualsStr(t, s.ToString(), "[[[1,1],[2,2]],[3,3]]", "should add snums")

	s = s.Add(Snum{LBracket, 4, Comma, 4, RBracket})
	assert.EqualsStr(t, s.ToString(), "[[[[1,1],[2,2]],[3,3]],[4,4]]", "should add snums")

	s = s.Add(Snum{LBracket, 5, Comma, 5, RBracket})
	assert.EqualsStr(t, s.ToString(), "[[[[3,0],[5,3]],[4,4]],[5,5]]", "should simplify after adding snums")

	s = s.Add(Snum{LBracket, 6, Comma, 6, RBracket})
	assert.EqualsStr(t, s.ToString(), "[[[[5,0],[7,4]],[5,5]],[6,6]]", "should simplify after adding snums")
}

func TestDay18_Simplify(t *testing.T) {
	s := ParseSnumStr("[0,[[[1,[2,3]],4],5]]")

	s = s.Simplify()

	assert.EqualsStr(t, s.ToString(), "[0,[[[3,0],7],5]]", "should convert to string")
}

func TestDay18_IsNormalAt(t *testing.T) {
	s1 := ParseSnumStr("[0,1]")
	s2 := ParseSnumStr("[0,[[[1,[2,3]],4],5]]")

	assert.IsTrue(t, s1.IsNormalAt(0), "should be a normal snum")
	assert.IsFalse(t, s2.IsNormalAt(0), "should NOT be a normal snum")
}

func TestDay18_DepthAt(t *testing.T) {
	s1 := ParseSnumStr("[1,[[2,[3,4]],5]]")

	assert.EqualsInt(t, s1.DepthAt(0), 1, "should calulate depth")
	assert.EqualsInt(t, s1.DepthAt(3), 2, "should calulate depth")
	assert.EqualsInt(t, s1.DepthAt(4), 3, "should calulate depth")
	assert.EqualsInt(t, s1.DepthAt(7), 4, "should calulate depth")

	s2 := ParseSnumStr("[[0,1],[[2,[3,4]],5]]")
	assert.EqualsInt(t, s2.DepthAt(7), 2, "should calulate depth")
}

func TestDay18_ShouldExplodeAt(t *testing.T) {
	s := ParseSnumStr("[0,[[[1,[[2,3],4]],4],5]]")

	assert.IsTrue(t, s.ShouldExplodeAt(9), "should explode a deep normal snum")
	assert.IsFalse(t, s.ShouldExplodeAt(0), "should NOT explode a shallow nested snum")
	assert.IsFalse(t, s.ShouldExplodeAt(8), "should NOT explode a deep nested snum")
}

// func TestDay18_ExplodeValue(t *testing.T) {
// 	// skip
// }

func TestDay18_ShouldSplit(t *testing.T) {
	s1 := ParseSnumStr("[0,1]")
	s2 := ParseSnumStr("[11,1]")
	s3 := ParseSnumStr("[[1,2],13]")

	assert.IsFalse(t, s1.ShouldSplitAt(0), "should NOT split a small normal snum")
	assert.IsTrue(t, s2.ShouldSplitAt(1), "should split a large normal snum")
	assert.IsTrue(t, s3.ShouldSplitAt(7), "should split a large nested snum")
}

func TestDay18_SplitAt(t *testing.T) {
	s := ParseSnumStr("[11,1]")
	s.SplitAt(1)

	assert.EqualsStr(t, s.ToString(), "[[5,6],1]", "should split to a smaller set")
}

func TestDay18_CalcMagnitude(t *testing.T) {
	assert.EqualsInt(t, ParseSnumStr("[[1,2],[[3,4],5]]").CalcMagnitude(), 143, "should calc magnitude")
	assert.EqualsInt(t, ParseSnumStr("[[[[0,7],4],[[7,8],[6,0]]],[8,1]]").CalcMagnitude(), 1384, "should calc magnitude")
	assert.EqualsInt(t, ParseSnumStr("[[[[1,1],[2,2]],[3,3]],[4,4]]").CalcMagnitude(), 445, "should calc magnitude")
	assert.EqualsInt(t, ParseSnumStr("[[[[3,0],[5,3]],[4,4]],[5,5]]").CalcMagnitude(), 791, "should calc magnitude")
	assert.EqualsInt(t, ParseSnumStr("[[[[5,0],[7,4]],[5,5]],[6,6]]").CalcMagnitude(), 1137, "should calc magnitude")
	assert.EqualsInt(t, ParseSnumStr("[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]").CalcMagnitude(), 3488, "should calc magnitude")
}

func Test18_AddABunchOfNumbers(t *testing.T) {
	s := ParseSnumStr("[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]")

	s = AddSnum(s, ParseSnumStr("[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]"))
	assert.EqualsStr(t, s.ToString(), "[[[[4,0],[5,4]],[[7,7],[6,0]]],[[8,[7,7]],[[7,9],[5,0]]]]", "should match AoC example")

	s = AddSnum(s, ParseSnumStr("[[2,[[0,8],[3,4]]],[[[6,7],1],[7,[1,6]]]]"))
	assert.EqualsStr(t, s.ToString(), "[[[[6,7],[6,7]],[[7,7],[0,7]]],[[[8,7],[7,7]],[[8,8],[8,0]]]]", "should match AoC example")

	s = AddSnum(s, ParseSnumStr("[[[[2,4],7],[6,[0,5]]],[[[6,8],[2,8]],[[2,1],[4,5]]]]"))
	assert.EqualsStr(t, s.ToString(), "[[[[7,0],[7,7]],[[7,7],[7,8]]],[[[7,7],[8,8]],[[7,7],[8,7]]]]", "should match AoC example")

	s = AddSnum(s, ParseSnumStr("[7,[5,[[3,8],[1,4]]]]"))
	assert.EqualsStr(t, s.ToString(), "[[[[7,7],[7,8]],[[9,5],[8,7]]],[[[6,8],[0,8]],[[9,9],[9,0]]]]", "should match AoC example")

	s = AddSnum(s, ParseSnumStr("[[2,[2,2]],[8,[8,1]]]"))
	assert.EqualsStr(t, s.ToString(), "[[[[6,6],[6,6]],[[6,0],[6,7]]],[[[7,7],[8,9]],[8,[8,1]]]]", "should match AoC example")

	s = AddSnum(s, ParseSnumStr("[2,9]"))
	assert.EqualsStr(t, s.ToString(), "[[[[6,6],[7,7]],[[0,7],[7,7]]],[[[5,5],[5,6]],9]]", "should match AoC example")

	s = AddSnum(s, ParseSnumStr("[1,[[[9,3],9],[[9,0],[0,7]]]]"))
	assert.EqualsStr(t, s.ToString(), "[[[[7,8],[6,7]],[[6,8],[0,8]]],[[[7,7],[5,0]],[[5,5],[5,6]]]]", "should match AoC example")

	s = AddSnum(s, ParseSnumStr("[[[5,[7,4]],7],1]"))
	assert.EqualsStr(t, s.ToString(), "[[[[7,7],[7,7]],[[8,7],[8,7]]],[[[7,0],[7,7]],9]]", "should match AoC example")

	s = AddSnum(s, ParseSnumStr("[[[[4,2],2],6],[8,7]]"))
	assert.EqualsStr(t, s.ToString(), "[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]", "should match AoC example")
}
