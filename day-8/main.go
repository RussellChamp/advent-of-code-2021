/* Day 8: Seven Segment Search */
package main

import (
	"AoC2021/utils/arrays"
	"AoC2021/utils/bits"
	"AoC2021/utils/log"
	"AoC2021/utils/timer"
	"bufio"
	"math"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		log.Println(log.DEBUG, "Oh snap!")
		panic(e)
	}
}

func main() {
	log.Println(log.NORMAL, "--- Day 8: Seven Segment Search ---")
	timer.Start()
	part1()
	timer.Tick()
	log.Println(log.NORMAL)

	part2()
	timer.Tick()
	log.Println(log.NORMAL)
}

func part1() {
	log.Println(log.NORMAL, "* Part 1 *")
	log.Println(log.NORMAL, " Goal: For now, focus on the easy digits.")
	log.Println(log.NORMAL, " Answer: In the output values, how many times do digits 1, 4, 7, or 8 appear?")

	input, err := os.Open("./input.txt")
	check(err)
	defer input.Close()

	scanner := bufio.NewScanner(input)
	lines := 0
	ones := 0
	fours := 0
	sevens := 0
	eights := 0

	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), " | ")
		if len(fields) != 2 {
			panic("Read invalid input line")
		}
		_ = strings.Fields(fields[0])
		// count all the ones, fours, sevens, and eights in the output strings based on how many segments are lit up
		outputDigits := strings.Fields(fields[1])
		for _, d := range outputDigits {
			switch len(d) {
			case 2:
				ones += 1
			case 3:
				sevens += 1
			case 4:
				fours += 1
			case 7:
				eights += 1
			}
		}
		lines++
	}

	total := ones + fours + sevens + eights
	log.Printf(log.NORMAL, "Read %d lines from input\n", lines)
	log.Printf(log.NORMAL, "We counted %d ones, %d fours, %d sevens, and %d eights for a total of %d\n", ones, fours, sevens, eights, total)
}

func part2() {
	log.Println(log.NORMAL, "* Part 2 *")
	log.Println(log.NORMAL, " Goal: For each entry, determine all of the wire/segment connections and decode the four-digit output values")
	log.Println(log.NORMAL, " Answer: What do you get if you add up all of the output values?")

	input, err := os.Open("./input.txt")
	check(err)
	defer input.Close()

	scanner := bufio.NewScanner(input)
	lines := 0
	total := 0

	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), " | ")
		if len(fields) != 2 {
			panic("Read invalid input line")
		}
		inputFields := strings.Fields(fields[0])
		segMap := createPatternMap(arrays.MapStrToInt(inputFields, ParseField))
		log.Println(log.DEBUG, "Created pattern", segMap)

		outputFields := strings.Fields(fields[1])

		fieldValue := 0
		fieldsCount := len(outputFields)
		for idx, field := range outputFields {
			fValue := segMap[ParseField(field)]
			log.Printf(log.DIAGNOSTIC, "field %s => %d; ", field, fValue)
			fieldValue += fValue * int(math.Pow10(fieldsCount-idx-1))
		}
		total += fieldValue

		log.Println(log.DIAGNOSTIC)
		log.Printf(log.DIAGNOSTIC, "Unscrambled %s => %d\n", outputFields, fieldValue)
		lines++
	}

	log.Printf(log.NORMAL, "Read %d lines from input\n", lines)
	log.Printf(log.NORMAL, "Calculated total is %d\n", total)
}

// Represent a segment as a 7-bit value
/* Segment values are aligned as below
*   aaaa
*  b    c
*  b    c
*   dddd
*  e    f
*  e    f
*   gggg

* segment 'a' is found in [0, 2, 3, 5, 6, 7, 8, 9]    but not [1, 4]
* segment 'b' is found in [0, 4, 5, 6, 8, 9]          but not [1, 2, 3, 7]
* segment 'c' is found in [0, 1, 2, 3, 4, 7, 8, 9]    but not [5, 6]
* segment 'd' is found in [2, 3, 4, 5, 6, 8, 9]       but not [0, 1, 7]
* segment 'e' is found in [0, 2, 6, 8]                but not [1, 3, 4, 5, 7, 9]
* segment 'f' is found in [0, 1, 3, 4, 5, 6, 7, 8, 9] but not [2]
* segment 'g' is found in [0, 2, 3, 5, 6, 8, 9]       but not [1, 4, 7]
 */

// parse a field from a string value to a bit-map int value (eg 'cda' to b1101 => 13)
func ParseField(field string) int {
	value := 0
	for _, c := range field {
		value += 1 << int(rune(c)-rune('a'))
	}
	return value
}

// func getDefaultMap() map[int]int {
// 	return map[int]int{
// 		ParseField("abcefg"):  0,
// 		ParseField("cf"):      1,
// 		ParseField("acdeg"):   2,
// 		ParseField("acdfg"):   3,
// 		ParseField("bcdf"):    4,
// 		ParseField("abdfg"):   5,
// 		ParseField("abdefg"):  6,
// 		ParseField("acf"):     7,
// 		ParseField("abcdefg"): 8,
// 		ParseField("abcdfg"):  9}
// }

// Given a list of a all 10 possible 7-segment displays create a map that can be used to unscramble values
func createPatternMap(input []int) map[int]int {
	// initialize our translation map that will convert from character to character (eg wire 'a' actually goes to segment 'c')
	segMap := make(map[int]int)

	// first find and map all the values that can be known by counting segments
	// find the only pattern that matches '1'
	one := arrays.FindFirstInt(input, bits.ByBitCount(2))

	// next find the one that matches '7'
	seven := arrays.FindFirstInt(input, bits.ByBitCount(3))

	// next find the one that matches '4'
	four := arrays.FindFirstInt(input, bits.ByBitCount(4))

	// the one that matches '8' will be the only one with 7 segments but tells us nothing about what those segments map to
	eight := arrays.FindFirstInt(input, bits.ByBitCount(7))

	// the two segments in '1' will match 'c' or 'f' but we don't know yet
	// the additional segment in '7' will map to segment 'a'
	segMap['a'] = seven - one
	// the two additional segments in '4' will match 'b' or 'd' but we're not sure which ones yet

	// all of '0', '6' and '9' have 6 segments
	// the six-segment field that shares only 1 side with '1' is '6'
	six := arrays.FindFirstInt(input, func(v int) bool { return bits.ByBitCount(6)(v) && bits.MatchingOnBits(one, v) == 1 })

	// and we know that the segment not in six is 'c'
	segMap['c'] = eight - six
	// and the other value 'f' will be the other segment in '1'
	segMap['f'] = one - segMap['c']
	// but we can't yet place '0' or '9' with what we know

	// all of '2', '3', and '5' have 5 segments
	// three is the only 5 segment field that matches 3 bits in '7'
	three := arrays.FindFirstInt(input, func(v int) bool { return bits.ByBitCount(5)(v) && bits.MatchingOnBits(seven, v) == 3 })
	// the additional two segments are 'd' ang 'g' but we can't place them yet
	// five is the other 5 segment field that matches 3 bits in '4'
	five := arrays.FindFirstInt(input, func(v int) bool { return bits.ByBitCount(5)(v) && v != three && bits.MatchingOnBits(four, v) == 3 })
	// segments 'b', 'd', and 'g' can't yet be placed
	// but segment 'e' is what we get by subtracting five from six
	segMap['e'] = six - five

	// the last five segment number must be 2
	two := arrays.FindFirstInt(input, func(v int) bool { return bits.ByBitCount(5)(v) && v != three && v != five })

	// segment 'g' is the only one that doesn't appear in '1', '4', or '7' and is not 'e'
	segMap['g'] = eight - bits.BitwiseAndA(one, four, seven, segMap['e'])
	// segment 'd' is '3' minus 'a', 'c', 'f', and 'g'
	segMap['d'] = three - segMap['a'] - segMap['c'] - segMap['f'] - segMap['g']

	// then we know that zero is eight minus 'd'
	zero := eight - segMap['d']
	// and nine will be eight minus 'e'
	nine := eight - segMap['e']

	log.Println(log.DIAGNOSTIC, "segMap", segMap)

	return map[int]int{
		zero:  0,
		one:   1,
		two:   2,
		three: 3,
		four:  4,
		five:  5,
		six:   6,
		seven: 7,
		eight: 8,
		nine:  9,
	}
}
