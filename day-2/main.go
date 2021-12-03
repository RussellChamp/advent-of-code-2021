/* Day 2: Dive! */
package main

import (
	"AoC2021/utils/timer"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		fmt.Println("Oh snap!")
		panic(e)
	}
}

func main() {
	fmt.Println("--- Day 2: Dive! ---")
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
	fmt.Println(" Goal: Calculate the horizontal position and depth you would have after following the planned course")
	fmt.Println(" Answer: What do you get if you multiply your final horizontal position by your final depth?")

	input, err := os.Open("./input.txt")
	defer input.Close()
	check(err)

	scanner := bufio.NewScanner(input)
	lines := 0
	posX, posY := 0, 0

	for scanner.Scan() {
		str := scanner.Text()
		values := strings.Fields(str)
		if len(values) < 2 {
			panic("read invalid line: " + str)
		}
		direction := values[0]
		units, err := strconv.Atoi(values[1])
		check(err)
		switch direction {
		case "forward":
			posX += units
		case "down":
			posY += units
		case "up":
			posY -= units
		default:
			panic("Invalid direction: " + direction)
		}

		lines++
	}

	fmt.Printf("Read %d lines from input\n", lines)
	fmt.Printf("The final position is (%d, %d) for an answer of %d\n", posX, posY, posX*posY)
}

func part2() {
	fmt.Println("* Part 2 *")
	fmt.Println(" Goal: Using this new interpretation of the commands, calculate the horizontal position and depth you would have after following the planned course")
	fmt.Println(" Answer: What do you get if you multiply your final horizontal position by your final depth?")

	input, err := os.Open("./input.txt")
	defer input.Close()
	check(err)

	scanner := bufio.NewScanner(input)
	lines := 0
	posX, posY, aim := 0, 0, 0

	for scanner.Scan() {
		str := scanner.Text()
		values := strings.Fields(str)
		if len(values) < 2 {
			panic("read invalid line: " + str)
		}
		direction := values[0]
		units, err := strconv.Atoi(values[1])
		check(err)
		switch direction {
		case "forward":
			posX += units
			posY += units * aim
		case "down":
			aim += units
		case "up":
			aim -= units
		default:
			panic("Invalid direction: " + direction)
		}

		lines++
	}

	fmt.Printf("Read %d lines from input\n", lines)
	fmt.Printf("The final position is (%d, %d) for an answer of %d\n", posX, posY, posX*posY)
}
