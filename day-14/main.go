/* Day 14: Extended Polymerization */
package main

import (
	"AoC2021/utils/args"
	"AoC2021/utils/log"
	"AoC2021/utils/timer"
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	argMap := args.Parse(os.Args)
	setLogLevel(argMap)
	selectedPart, partSpecified := argMap["part"]

	log.Println(log.NORMAL, "--- Day 14: Extended Polymerization ---")
	timer.Start()
	if !partSpecified || selectedPart == "1" {
		part1()
		timer.TickAtLevel(log.NORMAL)
		log.Println(log.NORMAL)
	}

	if !partSpecified || selectedPart == "2" {
		part2()
		timer.TickAtLevel(log.NORMAL)
		log.Println(log.NORMAL)
	}
}

func part1() {
	log.Println(log.NORMAL, "* Part 1 *")
	log.Println(log.NORMAL, " Goal: Apply 10 steps of pair insertion to the polymer template and find the most and least common elements in the result.")
	log.Println(log.NORMAL, " Answer: What do you get if you take the quantity of the most common element and subtract the quantity of the least common element?")

	doTheThing(10)
}

func part2() {
	log.Println(log.NORMAL, "* Part 2 *")
	log.Println(log.NORMAL, " Goal: Apply 40 steps of pair insertion to the polymer template and find the most and least common elements in the result.")
	log.Println(log.NORMAL, " Answer: What do you get if you take the quantity of the most common element and subtract the quantity of the least common element?")

	doTheThing(40)
}

func doTheThing(totalSteps int) {

	template, pairs := readInput()
	log.Println(log.DIAGNOSTIC, template, pairs)

	// perform the loop 'totalSteps' times
	for step := 0; step < totalSteps; step++ {
		// starting from the front of the template, check two-letter pairs
		for tIdx := 0; tIdx < len(template)-1; tIdx++ {
			// the lookup key will be the next two letters in the template
			key := template[tIdx : tIdx+2]
			// if we find a match in the pairs map
			value, found := pairs[key]
			if found {
				// insert the new character between the original two letters
				template = template[:tIdx+1] + value + template[tIdx+1:]
				// increment the idx counter once more so that we skip over the character that we just inserted
				tIdx++
			}
		}
		log.Printf(log.DEBUG, "After step #%d template is length %d\n", step+1, len(template))
		timer.TickAtLevel(log.DEBUG)
	}

	// COUNT all character occurances
	countMap := make(map[byte]int)

	for idx := 0; idx < len(template); idx++ {
		countMap[template[idx]]++
	}
	log.Println(log.DIAGNOSTIC, "CountMap:", countMap)

	mostCommonLetter, mostCommonCount := getMostCommon(countMap)
	leastCommonLetter, leastCommonCount := getLeastCommon(countMap)

	log.Printf(log.NORMAL, "After %d steps, the most common (%s=%d) minus the least common (%s=%d) is %d\n", totalSteps, mostCommonLetter, mostCommonCount, leastCommonLetter, leastCommonCount, mostCommonCount-leastCommonCount)
}

func check(e error) {
	if e != nil {
		log.Println(log.DEBUG, "Oh snap!")
		panic(e)
	}
}

func setLogLevel(argMap map[string]string) {
	logLevelStr, found := argMap["log"]
	if found {
		logLevel, err := strconv.Atoi(logLevelStr)
		check(err)
		log.SetLogLevel(logLevel)
	}
}

func readInput() (string, map[string]string) {
	input, err := os.Open("./input.txt")
	check(err)
	defer input.Close()

	scanner := bufio.NewScanner(input)
	lines := 0

	// TEMPLATE LINE
	if !scanner.Scan() {
		panic("Error parsing input template")
	}
	template := scanner.Text()
	lines++

	// NEW LINE
	if !scanner.Scan() {
		panic("Error parsing input newline")
	}
	_ = scanner.Text()
	lines++

	// PAIRS
	pairs := make(map[string]string)
	for scanner.Scan() {
		pair := strings.Split(scanner.Text(), " -> ")
		if len(pair) != 2 {
			panic("Error parsing input pairs")
		}
		pairs[pair[0]] = pair[1]
		lines++
	}

	log.Printf(log.NORMAL, "Read %d pairs from %d lines in input\n", len(pairs), lines)

	return template, pairs
}

func getMostCommon(countMap map[byte]int) (string, int) {
	mostCommonStr := ""
	mostCommonCount := -math.MaxInt
	for key, value := range countMap {
		if value > mostCommonCount {
			mostCommonStr, mostCommonCount = string(key), value
		}
	}

	return mostCommonStr, mostCommonCount
}

func getLeastCommon(countMap map[byte]int) (string, int) {
	leastCommonStr := ""
	leastCommonCount := math.MaxInt
	for key, value := range countMap {
		if value < leastCommonCount {
			leastCommonStr, leastCommonCount = string(key), value
		}
	}

	return leastCommonStr, leastCommonCount
}
