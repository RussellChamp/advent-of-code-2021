/* Day 6: Lanternfish */
package main

import (
	"AoC2021/utils/log"
	"AoC2021/utils/timer"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const REPRODUCE_DAYS = 7
const JUVINILE_DAYS = 2

func check(e error) {
	if e != nil {
		log.Println(log.DEBUG, "Oh snap!")
		panic(e)
	}
}

func main() {
	log.LOG_LEVEL = log.NORMAL

	log.Println(log.NORMAL, "--- Day 6: Lanternfish ---")
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
	log.Println(log.NORMAL, " Goal: Find a way to simulate lanternfish.")
	log.Println(log.NORMAL, " Answer: How many lanternfish would there be after 80 days?")

	justDoIt(80)
}

func part2() {
	log.Println(log.NORMAL, "* Part 2 *")
	log.Println(log.NORMAL, " Goal: Would they take over the entire ocean?")
	log.Println(log.NORMAL, " Answer: How many lanternfish would there be after 256 days?")

	justDoIt(256)
}

func justDoIt(total_days int) {
	input, err := os.Open("./input.txt")
	check(err)
	defer input.Close()

	scanner := bufio.NewScanner(input)

	// Keeping track of each individual fish is stupid
	// Instead track how *many* fish are at each gestation period
	groupCount := make([]uint, REPRODUCE_DAYS+JUVINILE_DAYS)

	startingCount := 0
	lines := 0

	// Read in all the starting fish
	for scanner.Scan() {
		startingFish := strings.Split(scanner.Text(), ",")
		startingCount += len(startingFish)

		lines++
		// A fish read in as day '3' will be counted at index '2' since day '0' is a valid day
		for _, fish := range startingFish {
			day, err := strconv.Atoi(fish)
			check(err)
			if day > REPRODUCE_DAYS {
				panic(fmt.Sprintf("Read an invalid value for fish day: %d", day))
			}
			groupCount[day] += 1
		}
	}
	log.Printf(log.NORMAL, "Parsed %d fish from %d lines from input\n", startingCount, lines)
	log.Println(log.DEBUG, "Fishes group count", groupCount)
	log.Println(log.DEBUG)

	// Begin the simulation
	for day := 0; day < total_days; day++ {
		// keep track of this for later
		readyToReproduce := groupCount[0]
		log.Printf(log.DEBUG, "Day %d starts with %d fish ", day+1, sumFish(groupCount))
		log.Print(log.DEBUG, groupCount)
		// shift down ALL groups
		for fishDay := 1; fishDay < REPRODUCE_DAYS+JUVINILE_DAYS; fishDay++ {
			groupCount[fishDay-1] = groupCount[fishDay]
		}
		// Spawn new fish
		groupCount[REPRODUCE_DAYS+JUVINILE_DAYS-1] = readyToReproduce
		// Add the fish that reproduced back into the pool
		groupCount[REPRODUCE_DAYS-1] += readyToReproduce
		log.Printf(log.DEBUG, " and ends with %d fish ", sumFish((groupCount)))
		log.Print(log.DEBUG, groupCount)
		log.Println(log.DEBUG)
	}

	// Add them all up and be sad that there is no map reduce in golang
	totalCount := sumFish(groupCount)

	log.Printf(log.NORMAL, "At the end of %d days there were %d fish\n", total_days, totalCount)
}

func sumFish(counts []uint) uint {
	total := uint(0)
	for _, count := range counts {
		total += count
	}

	return total
}
