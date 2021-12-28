/* Day 17: Trick Shot */
// This question is written very strange. Skipping for now
package main

import (
	"AoC2021/utils/args"
	"AoC2021/utils/log"
	"AoC2021/utils/timer"
	"bufio"
	"os"
	"regexp"
	"strconv"
)

func check(e error) {
	if e != nil {
		log.Println(log.DEBUG, "Oh snap!")
		panic(e)
	}
}

const XMIN = 20
const XMAX = 30
const YMIN = -10
const YMAX = -5

func main() {
	argMap := args.Parse(os.Args)
	log.SetLogLevelFromArgs(argMap)
	selectedPart, partSpecified := argMap["part"]

	log.Println(log.NORMAL, "--- Day 17: Trick Shot ---")
	timer.Start()

	if !partSpecified || selectedPart == "" || selectedPart == "1" {
		part1()
		timer.Tick()
		log.Println(log.NORMAL)
	}

	if !partSpecified || selectedPart == "" || selectedPart == "2" {
		part2()
		timer.Tick()
		log.Println(log.NORMAL)
	}
}

func part1() {
	log.Println(log.NORMAL, "* Part 1 *")
	log.Println(log.NORMAL, " Goal: How high can you make the probe go while still reaching the target area?")
	log.Println(log.NORMAL, " Answer: What is the highest y position it reaches on this trajectory?")

	x1, x2, y1, y2 := parseInputFile()

	drawGrid(x1, x2, y1, y2)

	log.Printf(log.NORMAL, "Read values from input. x: %d..%d, y: %d..%d\n", x1, x2, y1, y2)
	log.Printf(log.NORMAL, "The solution is %d\n", 42)
}

func part2() {
	log.Println(log.NORMAL, "* Part 2 *")
	log.Println(log.NORMAL, " Goal: ")
	log.Println(log.NORMAL, " Answer: ")
}

func parseInputFile() (int, int, int, int) {

	input, err := os.Open("./input.txt")
	check(err)
	defer input.Close()

	scanner := bufio.NewScanner(input)

	if !scanner.Scan() {
		panic("failed to load input data")
	}

	r, _ := regexp.Compile("target area: x=(-?[0-9]+)..(-?[0-9]+), y=(-?[0-9]+)..(-?[0-9]+)")
	matches := r.FindStringSubmatch(scanner.Text())
	if len(matches) != 5 {
		panic("failed to parse input data")
	}

	x1, x2, y1, y2 := 0, 0, 0, 0
	x1, err = strconv.Atoi(matches[1])
	check(err)
	x2, err = strconv.Atoi(matches[2])
	check(err)
	y1, err = strconv.Atoi(matches[3])
	check(err)
	y2, err = strconv.Atoi(matches[2])
	check(err)

	return x1, x2, y1, y2
}

func drawGrid(x1, x2, y1, y2 int) {
	for y := y1; y <= y2; y++ {
		for x := x1; x <= x2; x++ {
			switch {
			case x == 0 && y == 0:
				log.Printf(log.NORMAL, "S")
			case x >= XMIN && x <= XMAX && y >= YMIN && y <= YMAX:
				log.Printf(log.NORMAL, "T")
			default:
				log.Printf(log.NORMAL, ".")
			}
		}
		log.Println(log.NORMAL)
	}
	log.Println(log.NORMAL)
}
