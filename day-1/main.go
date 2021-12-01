/* Day 1: Sonar Sweep */
package main

import (
	"AoC2021/utils/timer"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const DEBUG = true
const WIDTH = 100

func check(e error) {
	if e != nil {
		fmt.Println("Oh snap!")
		panic(e)
	}
}

func main() {
	fmt.Println("--- Day 1: Sonar Sweep ---")
	timer.Start()
	part1()
	timer.Tick()
	fmt.Println()

	part2()
	timer.Tick()
	fmt.Println()
}

func part1() {
	fmt.Println("* Part 1 *")
	fmt.Println(" Goal: Count the number of times a depth measurement increases from the previous measurement.")
	fmt.Println(" Answer: How many measurements are larger than the previous measurement?")

	input, err := os.Open("./input.txt")
	check(err)
	defer input.Close()

	scanner := bufio.NewScanner(input)

	lastValue := -1
	measurements := 0
	increases := 0

	for scanner.Scan() {
		value, err := strconv.Atoi(scanner.Text())
		check(err)
		measurements++
		if lastValue != -1 && lastValue < value {
			increases++
		}
		printPip(lastValue, lastValue, value, measurements)

		lastValue = value
	}

	fmt.Printf("Read %d measurements resulting in %d increases\n", measurements, increases)
}

func part2() {
	fmt.Println("* Part 2 *")
	fmt.Println(" Goal: Count the number of times the sum of measurements in this sliding window increases")
	fmt.Println(" Answer: How many sums are larger than the previous sum?")

	input, err := os.Open("./input.txt")
	check(err)
	defer input.Close()

	scanner := bufio.NewScanner(input)

	w1, w2, w3 := -1, -1, -1
	measurements := 0
	increases := 0

	for scanner.Scan() {
		value, err := strconv.Atoi(scanner.Text())
		check(err)
		measurements++

		previous := w1 + w2 + w3
		current := w2 + w3 + value

		if w1 != -1 && current > previous {
			increases++
		}

		printPip(w1, previous, current, measurements)

		// shift over the window
		w1, w2, w3 = w2, w3, value
	}

	fmt.Printf("Read %d measurements resulting in %d increases\n", measurements, increases)
}

func printPip(first int, previous int, current int, total int) {
	if !DEBUG {
		return
	}

	if first != -1 && current > previous {
		fmt.Print("+")
	} else if first != -1 && current < previous {
		fmt.Print("-")
	} else {
		fmt.Print(".")
	}
	if total%WIDTH == 0 {
		fmt.Println() // linebreak
	}
}
