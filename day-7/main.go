/* Day 7: The Treachery of Whales */
package main

import (
	"AoC2021/utils/log"
	"AoC2021/utils/timer"
	"bufio"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		log.Println(log.DEBUG, "Oh snap!")
		panic(e)
	}
}

func main() {
	log.Println(log.NORMAL, "--- Day 7: The Treachery of Whales ---")
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
	log.Println(log.NORMAL, " Goal: Determine the horizontal position that the crabs can align to using the least fuel possible")
	log.Println(log.NORMAL, " Answer: How much fuel must they spend to align to that position?")

	doTheNeedful(calcLinearFuelCost)
}

func part2() {
	log.Println(log.NORMAL, "* Part 2 *")
	log.Println(log.NORMAL, " Goal: Determine the horizontal position that the crabs can align to using the least fuel possible so they can make you an escape route!")
	log.Println(log.NORMAL, " Answer: How much fuel must they spend to align to that position?")

	doTheNeedful(calcSpecialFuelCost)
}

func doTheNeedful(calcFuelCost func([]int, int) int) {
	input, err := os.Open("./input.txt")
	check(err)
	defer input.Close()

	scanner := bufio.NewScanner(input)
	lines := 0

	var crabsPos []int

	// Parse the crab positions
	for scanner.Scan() {
		crabStrs := strings.Split(scanner.Text(), ",")
		for _, crabStr := range crabStrs {
			value, err := strconv.Atoi(crabStr)
			check(err)
			crabsPos = append(crabsPos, value)

		}
		lines++
	}

	// Find the optimal horizontal position
	// We're gonna brute force it for now and try to come up with a better method later
	// Basically do an integral manually. Starting from the lowest hPos calculate the cost
	// and move up until we are no longer reducing the cost
	minPos := minVal(crabsPos)
	maxPos := maxVal(crabsPos)

	bestPos := minPos
	bestCost := calcFuelCost(crabsPos, minPos)

	for pos := bestPos + 1; pos < maxPos; pos++ {
		curCost := calcFuelCost(crabsPos, pos)
		// we should have started from an unoptimized position. as we move closer to the center, the cost should reduce
		// once the cost starts increasing again we know that the *last* value was the optimial one
		if curCost > bestCost {
			bestPos = pos - 1
			break
		} else {
			// we found a better cost
			bestCost = curCost
		}
	}

	log.Printf(log.NORMAL, "Read %d crabs on %d lines from input between positions %d and %d\n", len(crabsPos), lines, minPos, maxPos)
	log.Printf(log.NORMAL, "The best position is %d with a cost of %d\n", bestPos, bestCost)
}

func calcLinearFuelCost(positions []int, hPos int) int {
	cost := 0
	for _, p := range positions {
		cost += abs(p - hPos)
	}
	return cost
}

func calcSpecialFuelCost(positions []int, hPos int) int {
	cost := 0
	for _, p := range positions {
		shift := abs(p - hPos)
		cost += (shift * (shift + 1)) / 2
	}

	return cost
}

func abs(x int) int {
	if x >= 0 {
		return x
	} else {
		return -x
	}
}

func minVal(a []int) int {
	if len(a) == 0 {
		panic("minVal Passed empty array")
	}

	min := 0
	for idx, val := range a {
		if idx == 0 || val < min {
			min = val
		}
	}
	return min
}

func maxVal(a []int) int {
	if len(a) == 0 {
		panic("minVal Passed empty array")
	}

	min := 0
	for idx, val := range a {
		if idx == 0 || val > min {
			min = val
		}
	}
	return min
}
