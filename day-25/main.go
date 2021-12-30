/* Day 25: Sea Cucumber */
package main

import (
	"AoC2021/utils/args"
	"AoC2021/utils/log"
	"AoC2021/utils/timer"
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		log.Println(log.DEBUG, "Oh snap!")
		panic(e)
	}
}

func main() {
	argMap := args.Parse(os.Args)
	log.SetLogLevelFromArgs(argMap)
	selectedPart, partSpecified := argMap["part"]

	log.Println(log.NORMAL, "--- Day 25: Sea Cucumber ---")
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
	log.Println(log.NORMAL, " Goal: Find somewhere safe to land your submarine.")
	log.Println(log.NORMAL, " Answer: What is the first step on which no sea cucumbers move?")

	grid := readData()

	round, hCukes, vCukes := 0, 0, 0
	for cukesMoved := 0; round == 0 || cukesMoved > 0; round++ {
		hCukes, vCukes = grid.Move()
		cukesMoved = hCukes + vCukes
		if log.LOG_LEVEL >= log.DEBUG && round%1000 == 0 {
			log.Printf(log.DEBUG, "After %d rounds, %d hCukes and %d vCukes moved...\n", round, hCukes, vCukes)
		}
	}
	log.Printf(log.NORMAL, "Everything settled on round %d\n", round)
}

func part2() {
	log.Println(log.NORMAL, "* Part 2 *")
	log.Println(log.NORMAL, " Goal: ")
	log.Println(log.NORMAL, " Answer: ")
}

func readData() Grid {
	input, err := os.Open("./input.txt")
	check(err)
	defer input.Close()

	scanner := bufio.NewScanner(input)
	lines := 0
	hCukes, vCukes, blanks := 0, 0, 0

	var grid Grid

	for scanner.Scan() {
		line := scanner.Bytes()
		for _, b := range line {
			switch b {
			case '>':
				hCukes++
			case 'v':
				vCukes++
			case '.':
				blanks++
			default:
				panic(fmt.Sprintf("got invalid byte! %b", b))
			}
		}
		grid = append(grid, line)
		lines++
	}

	log.Printf(log.DEBUG, "Made a %dx%d grid from %d lines from input\n", len(grid[0]), len(grid), lines)
	log.Printf(log.DEBUG, "Grid contains %d hCukes and %d vCukes and %d blanks\n", hCukes, vCukes, blanks)

	return grid
}

type Grid [][]byte
type Point struct{ rIdx, cIdx int }

func (g *Grid) Move() (int, int) {
	// NOPE! not moving right!!

	// find horizontal movers
	var hMovers []Point
	for rIdx, row := range *g {
		for cIdx, value := range row {
			newCIdx := (cIdx + 1) % len(row)
			if value == '>' && row[newCIdx] == '.' {
				hMovers = append(hMovers, Point{rIdx: rIdx, cIdx: cIdx})
			}
		}
	}
	// move them
	for _, p := range hMovers {
		newCIdx := (p.cIdx + 1) % len((*g)[p.rIdx])
		log.Printf(log.DIAGNOSTIC, "Moving hCuke from (%d,%d) to (%d,%d)...\n", p.rIdx, p.cIdx, p.rIdx, newCIdx)
		(*g)[p.rIdx][p.cIdx] = '.'
		(*g)[p.rIdx][newCIdx] = '>'
	}

	// find vertical movers
	var vMovers []Point
	for rIdx, row := range *g {
		for cIdx, value := range row {
			newRIdx := (rIdx + 1) % len(*g)
			if value == 'v' && (*g)[newRIdx][cIdx] == '.' {
				vMovers = append(vMovers, Point{rIdx: rIdx, cIdx: cIdx})
			}
		}
	}
	// move them
	for _, p := range vMovers {
		newRIdx := (p.rIdx + 1) % len(*g)
		log.Printf(log.DIAGNOSTIC, "Moving vCuke from (%d,%d) to (%d,%d)...\n", p.rIdx, p.cIdx, newRIdx, p.cIdx)
		(*g)[p.rIdx][p.cIdx] = '.'
		(*g)[newRIdx][p.cIdx] = 'v'
	}

	return len(hMovers), len(vMovers)
}

func (g Grid) Print() {
	for _, row := range g {
		log.Println(log.DEBUG, string(row))
	}
}
