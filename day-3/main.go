/* Day 3: Binary Diagnostic */
package main

import (
	"AoC2021/utils/timer"
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

const BITLENGTH = 12

func check(e error) {
	if e != nil {
		fmt.Println("Oh snap!")
		panic(e)
	}
}

func main() {
	fmt.Println("--- Day 3: Binary Diagnostic ---")
	timer.Start()
	part1()
	timer.Tick()
	fmt.Println()

	part2()
	timer.Tick()
	fmt.Println()
}

func bodgeData(bits [BITLENGTH]int, fn func(int) bool) int {
	value := 0
	for idx := 0; idx < BITLENGTH; idx++ {
		//fmt.Print("bodgeData: comparing ", bits[idx])
		if fn(bits[idx]) {
			value += 1 << (BITLENGTH - idx - 1)
			//fmt.Print(": value is now ", value)
		}
		//fmt.Println()
	}

	return value
}

func part1() {
	fmt.Println("* Part 1 *")
	fmt.Println(" Goal: Use the binary numbers in your diagnostic report to calculate the gamma rate and epsilon rate, then multiply them together")
	fmt.Println(" Answer: What is the power consumption of the submarine?")

	input, err := os.Open("./input.txt")
	defer input.Close()
	check(err)

	reader := bufio.NewReader(input)
	lines := 0
	var bits [BITLENGTH]int

	activeBit := 0
	for {
		if c, _, err := reader.ReadRune(); err != nil {
			if err == io.EOF {
				lines++
				break
			} else {
				panic(err)
			}
		} else {
			if string(c) == "\n" {
				lines++
				activeBit = 0
				continue
			} else {
				val, err := strconv.Atoi(string(c))
				check(err)
				bits[activeBit] += val
				activeBit += 1
			}
		}
	}

	gamma := bodgeData(bits, func(val int) bool { return val > lines/2 })
	epsilon := bodgeData(bits, func(val int) bool { return val < lines/2 })

	fmt.Printf("Read %d lines from input\n", lines)
	fmt.Println("bits:", bits)
	fmt.Printf("Gamma: %d, Epsilon: %d, Solution: %d\n", gamma, epsilon, gamma*epsilon)
}

// compare a byte slice with a filter
func byteMatch(value []byte, filter [BITLENGTH]byte) bool {
	if len(value) != len(filter) {
		return false
	}
	for idx, b := range filter {
		if b == byte(0) {
			// once we  hit a part of the filter with empty values, we are good
			return true
		}
		if b != value[idx] {
			// if a part of our value every doesn't match the filter, we are bad
			return false
		}
	}
	return true
}

func getBitCount(input *os.File, position int, filter [BITLENGTH]byte) (int, int) {
	scanner := bufio.NewScanner(input)
	count := 0
	lines := 0
	for scanner.Scan() {
		bytes := scanner.Bytes()
		if !byteMatch(bytes, filter) {
			continue
		}
		lines += 1
		if len(bytes) < position {
			panic("not enough bytes")
		}
		if bytes[position] == byte('1') {
			count += 1
		}
	}

	return count, lines
}

func gitBitValue(input *os.File, filter [BITLENGTH]byte) int {
	scanner := bufio.NewScanner(input)
	value := 0

	for scanner.Scan() {
		bytes := scanner.Bytes()
		if !byteMatch(bytes, filter) {
			continue
		}
		// we got a match!
		for idx, b := range bytes {
			if b == byte('1') {
				value += 1 << (BITLENGTH - idx - 1)
			}
		}
		break
	}
	return value
}

func doTheThing(input *os.File, reverse bool) int {
	var filter [BITLENGTH]byte
	yesByte, noByte := byte('1'), byte('0')
	if reverse {
		yesByte, noByte = byte('0'), byte('1')
	}

	for idx := 0; idx < BITLENGTH; idx++ {
		_, err := input.Seek(0, io.SeekStart)
		check(err)

		ones, lines := getBitCount(input, idx, filter)
		//fmt.Printf("BC@%d: %d (%d lines) using filter [%s]\n", idx+1, ones, lines, filter)

		if lines == 1 {
			break
		}
		if float64(ones) >= float64(lines)/2 {
			filter[idx] = yesByte
		} else {
			filter[idx] = noByte
		}
	}

	_, err := input.Seek(0, io.SeekStart)
	check(err)

	// fmt.Printf("=> finding value that matches %s\n", filter)
	value := gitBitValue(input, filter)

	return value
}

func getGeneratorValue(input *os.File) int {
	return doTheThing(input, false)
}

func getScrubberValue(input *os.File) int {
	return doTheThing(input, true)
}

func part2() {
	fmt.Println("* Part 2 *")
	fmt.Println(" Goal: Use the binary numbers in your diagnostic report to calculate the oxygen generator rating and CO2 scrubber rating, then multiply them together")
	fmt.Println(" Answer: What is the life support rating of the submarine?")

	input, err := os.Open("./input.txt")
	defer input.Close()
	check(err)

	generator := getGeneratorValue(input)
	scrubber := getScrubberValue(input)

	fmt.Printf("Generator: %d, Scrubber: %d, Solution: %d\n", generator, scrubber, generator*scrubber)
}
