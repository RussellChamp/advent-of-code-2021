/* Day 5: Hydrothermal Venture */
package main

import (
	"AoC2021/utils/timer"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const MAX_SIZE = 1000

func check(e error) {
	if e != nil {
		fmt.Println("Oh snap!")
		panic(e)
	}
}

func checkCoord(x int, y int) {
	badVal := false
	if x < 0 || x >= MAX_SIZE {
		badVal = true
		fmt.Println("Invalid x value:", x)
	}
	if y < 0 || y >= MAX_SIZE {
		badVal = true
		fmt.Println("Invalid y value:", y)
	}

	if badVal {
		panic(fmt.Sprintf("Failed at (%d,%d)", x, y))
	}
}

func main() {
	fmt.Println("--- Day 5: Hydrothermal Venture ---")
	timer.Start()
	// part1()
	// timer.Tick()
	// fmt.Println()

	part2()
	timer.Tick()
	fmt.Println()
}

func parseLine(line string) (int, int, int, int) {
	coords := strings.Split(line, " -> ")
	if len(coords) != 2 {
		panic("Did not find coordinate pair!")
	}
	coord1 := strings.Split(coords[0], ",")
	x1, err := strconv.Atoi(coord1[0])
	check(err)
	y1, err := strconv.Atoi(coord1[1])
	check(err)

	coord2 := strings.Split(coords[1], ",")
	x2, err := strconv.Atoi(coord2[0])
	check(err)
	y2, err := strconv.Atoi(coord2[1])
	check(err)

	return x1, y1, x2, y2
}

// Goodness this is stupid that there are no generic functions for maths things
func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func Sign(a int) int {
	switch {
	case a > 0:
		return 1
	case a < 0:
		return -1
	default:
		return 0
	}
}

func part1() {
	fmt.Println("* Part 1 *")
	fmt.Println(" Goal: Determine the number of points where at least two lines overlap")
	fmt.Println(" Answer: At how many points do at least two lines overlap?")

	input, err := os.Open("./input.txt")
	check(err)
	defer input.Close()

	scanner := bufio.NewScanner(input)
	lines := 0
	goodLines := 0

	var grid [MAX_SIZE][MAX_SIZE]uint
	// fmt.Printf("Initialized grid of size %d\n", MAX_SIZE)

	// Draw the lines
	for scanner.Scan() {
		x1, y1, x2, y2 := parseLine(scanner.Text())
		lines++

		// only consider horizontal or vertical lines
		if x1 != x2 && y1 != y2 {
			continue
		}
		goodLines++
		xMin := Min(x1, x2)
		xMax := Max(x1, x2)
		yMin := Min(y1, y2)
		yMax := Max(y1, y2)
		for x := xMin; x <= xMax; x++ {
			for y := yMin; y <= yMax; y++ {
				// it doesn't matter if x or y is first as long as I'm consistent
				grid[x][y] += 1
			}
		}
	}
	// Count the intersections
	intersections := 0
	for x := 0; x < MAX_SIZE; x++ {
		for y := 0; y < MAX_SIZE; y++ {
			if grid[x][y] > 1 {
				intersections += 1
			}
		}
	}

	fmt.Printf("Read %d lines from input. Found %d straight segments\n", lines, goodLines)
	fmt.Printf("Found %d line intersections\n", intersections)
}

func part2() {
	fmt.Println("* Part 2 *")
	fmt.Println(" Goal: Including diagonal line segments, determine the number of points where at least two lines overlap")
	fmt.Println(" Answer: At how many points do at least two lines overlap?")

	input, err := os.Open("./input.txt")
	check(err)
	defer input.Close()

	scanner := bufio.NewScanner(input)
	lines := 0

	var grid [MAX_SIZE][MAX_SIZE]uint

	// Draw the lines
	for scanner.Scan() {
		x1, y1, x2, y2 := parseLine(scanner.Text())
		lines++

		// Start at x1, y1 and determine which "direction" you will need to go to reach x2, y2
		// eg
		// if x1 is 3 and x2 is 3 then the x direction will be '0', not moving at all in the x direction
		// if x1 is 5 and x2 is 10 then the x direction will be '1', going in the "positive" direction
		// if y1 is 100 and y2 is 20 then the y direction will be '-1', going in the "negative" direction
		xVel := Sign(x2 - x1)
		yVel := Sign(y2 - y1)

		//fmt.Printf("Drawing line from (%d,%d) to (%d,%d) using velocity [%d,%d]", x1, y1, x2, y2, xVel, yVel)

		dots := 0
		for x, y := x1, y1; x != x2 || y != y2; {
			checkCoord(x, y)

			grid[x][y] += 1

			x += xVel
			y += yVel
			dots += 1
		}
		// Mark the last point
		grid[x2][y2] += 1
		//fmt.Printf(" -- Drew %d dots\n", dots+1)
	}
	// Count the intersections
	intersections := 0
	for x := 0; x < MAX_SIZE; x++ {
		for y := 0; y < MAX_SIZE; y++ {
			if grid[x][y] > 1 {
				intersections += 1
			}
		}
	}

	fmt.Printf("Read %d lines from input.\n", lines)
	fmt.Printf("Found %d line intersections\n", intersections)
}
