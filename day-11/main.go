/* Day 11: Dumbo Octopus */
package main

import (
	"AoC2021/utils/arrays"
	"AoC2021/utils/log"
	"AoC2021/utils/timer"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const WIDTH = 10
const HEIGHT = 10
const ALREADY_FLASHED = -1
const STEPS = 100

type Grid = [][]int

func check(e error) {
	if e != nil {
		log.Println(log.DEBUG, "Oh snap!")
		panic(e)
	}
}

func main() {
	log.Println(log.NORMAL, "--- Day 11: Dumbo Octopus ---")
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
	log.Println(log.NORMAL, " Goal: Given the starting energy levels of the dumbo octopuses in your cavern, simulate 100 steps.")
	log.Println(log.NORMAL, " Answer: How many total flashes are there after 100 steps?")

	grid := loadGrid()
	flashCount := 0

	for step := 0; step < STEPS; step++ {
		newFlashes := stepSimulation(&grid)
		log.Printf(log.DEBUG, "Step %d: Saw %d new flashes\n", step, newFlashes)
		flashCount += newFlashes
	}

	log.Printf(log.NORMAL, "After %d steps, the total number of flashes was %d\n", STEPS, flashCount)
	log.Printf(log.DEBUG, " and the grid looked like: \n")
	printGrid(grid)
}

func part2() {
	log.Println(log.NORMAL, "* Part 2 *")
	log.Println(log.NORMAL, " Goal: If you can calculate the exact moments when the octopuses will all flash simultaneously, you should be able to navigate through the cavern.")
	log.Println(log.NORMAL, " Answer: What is the first step during which all octopuses flash?")

	grid := loadGrid()
	step := 0
	newFlashes := 0

	for ; newFlashes != HEIGHT*WIDTH; step++ {
		newFlashes = stepSimulation(&grid)
		log.Printf(log.DEBUG, "Step %d: Saw %d new flashes\n", step, newFlashes)
	}

	log.Printf(log.NORMAL, "After %d steps, we saw all %d nodes flash\n", step, newFlashes)
	log.Printf(log.DEBUG, " and the grid looked like: \n")
	printGrid(grid)
}

func loadGrid() Grid {
	input, err := os.Open("./input.txt")
	check(err)
	defer input.Close()

	scanner := bufio.NewScanner(input)
	lines := 0

	var grid Grid

	for scanner.Scan() {
		if len(grid) > HEIGHT {
			panic(fmt.Sprintf("Error: Tried to read more than %d lines from input", HEIGHT))
		}
		line := arrays.MapStrToInt(strings.Split(scanner.Text(), ""), myAtoi)

		if len(line) != WIDTH {
			panic(fmt.Sprintf("Error: Read line with invalid width: %d", len(line)))
		}
		grid = append(grid, line)
		lines++
	}
	if len(grid) != HEIGHT {
		panic(fmt.Sprintf("Error: Input was invalid length. Expected %d row but got %d", HEIGHT, len(grid)))
	}

	return grid
}

func myAtoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(fmt.Sprintf("Error: Attempted to parse invalid string to Atoi: '%s'", s))
	}
	return i
}

func printGrid(grid Grid) {
	for _, row := range grid {
		log.Println(log.DEBUG, row)
	}
}

// update the grid after a step and return the number of flashes that happened
func stepSimulation(grid *Grid) int {
	flashCount := 0

	// increase ALL energy levels by 1
	incrementEnergy(grid)

	for {
		// while nodes are still flashing, keep going
		flashes := flashFullNodesSinglePass(grid)
		if flashes == 0 {
			break
		}
		flashCount += flashes
	}

	// reset all nodes marked as -1 to 0
	resetTiredEnergy(grid)

	return flashCount
}

func incrementEnergy(grid *Grid) {
	for rIdx := 0; rIdx < HEIGHT; rIdx++ {
		for cIdx := 0; cIdx < WIDTH; cIdx++ {
			(*grid)[rIdx][cIdx] += 1
		}
	}
}

// flashes nodes that are at full energy, updates their neighbors, and returns the total number of nodes flashed
// does a SINGLE pass through the grid
func flashFullNodesSinglePass(grid *Grid) int {
	flashCount := 0

	for rIdx := 0; rIdx < HEIGHT; rIdx++ {
		for cIdx := 0; cIdx < WIDTH; cIdx++ {
			// an update can spiral from a single node
			// check how many flashed form this single node
			flashCount += checkNode(grid, rIdx, cIdx, 0)
		}
	}
	return flashCount
}

// check if a node should flash and then update it's neighbors
// if we're checking a node as a result of it's neighbor flashing, we also increment the energy
func checkNode(grid *Grid, row, col, bonusEnergy int) int {
	if !isNodeValid(grid, row, col) {
		return 0
	}
	if bonusEnergy > 0 {
		(*grid)[row][col] += bonusEnergy
	}
	if (*grid)[row][col] < 10 {
		return 0
	}

	flashes := 1

	// set the node to -1 to represent that it has flashed this step
	(*grid)[row][col] = ALREADY_FLASHED

	// increment all eight valid neighbors by 1 and check if they flash
	/* UP    */
	flashes += checkNode(grid, row-1, col, 1)
	/* DOWN  */ flashes += checkNode(grid, row+1, col, 1)
	/* LEFT  */ flashes += checkNode(grid, row, col-1, 1)
	/* RIGHT */ flashes += checkNode(grid, row, col+1, 1)
	/* UP-LEFT    */ flashes += checkNode(grid, row-1, col-1, 1)
	/* UP-RIGHT   */ flashes += checkNode(grid, row-1, col+1, 1)
	/* DOWN-LEFT  */ flashes += checkNode(grid, row+1, col-1, 1)
	/* DOWN-RIGHT */ flashes += checkNode(grid, row+1, col+1, 1)

	return flashes
}

// check if a node is valid
func isNodeValid(grid *Grid, row, col int) bool {
	if row < 0 || col < 0 || row >= HEIGHT || col >= WIDTH {
		return false
	}
	if (*grid)[row][col] == ALREADY_FLASHED {
		return false
	}

	return true
}

// nodes that have flashed this step are set to -1 so that they aren't accidentally incremented anymore
// reset that value back to 0 at the end of the step
func resetTiredEnergy(grid *Grid) {
	for rIdx := 0; rIdx < HEIGHT; rIdx++ {
		for cIdx := 0; cIdx < WIDTH; cIdx++ {
			if (*grid)[rIdx][cIdx] == ALREADY_FLASHED {
				(*grid)[rIdx][cIdx] = 0
			}
		}
	}
}
