/* Day 13: Transparent Origami */
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
	log.Println(log.NORMAL, "--- Day 13: Transparent Origami ---")
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
	log.Println(log.NORMAL, " Goal: The transparent paper is pretty big, so for now, focus on just completing the first fold.")
	log.Println(log.NORMAL, " Answer: How many dots are visible after completing just the first fold instruction on your transparent paper?")

	points, folds := readInput()
	grid := initGrid(points)
	dotCount := countDots(grid)

	log.Printf(log.DEBUG, "Initialized a grid of size %d by %d with %d dots\n", len(grid[0]), len(grid), dotCount)

	grid = applyFold(grid, folds[0])
	dotCount = countDots(grid)
	log.Printf(log.NORMAL, "After the first fold, the grid is now %d by %d with %d dots\n", len(grid[0]), len(grid), dotCount)
}

func part2() {
	log.Println(log.NORMAL, "* Part 2 *")
	log.Println(log.NORMAL, " Goal: Finish folding the transparent paper according to the instructions. The manual says the code is always eight capital letters.")
	log.Println(log.NORMAL, " Answer: What code do you use to activate the infrared thermal imaging camera system?")

	points, folds := readInput()
	grid := initGrid(points)
	dotCount := countDots(grid)

	log.Printf(log.DEBUG, "Initialized a grid of size %d by %d with %d dots\n", len(grid[0]), len(grid), dotCount)

	for fIdx, fold := range folds {
		grid = applyFold(grid, fold)
		dotCount = countDots(grid)
		log.Printf(log.DEBUG, "After the fold #%d, the grid is now %d by %d with %d dots\n", fIdx+1, len(grid[0]), len(grid), dotCount)
	}

	log.Printf(log.NORMAL, "After %d folds the grid is %d by %d and has %d dots\n", len(folds), len(grid[0]), len(grid), dotCount)
	log.Printf(log.NORMAL, "The secret code in ASCII format is:\n")
	printGrid(grid)
	log.Printf(log.NORMAL, "(You figure it out. Your brain will work better than janky OCR\n")
}

type Point struct {
	x, y int
}

type Fold struct {
	direction string
	value     int
}

func readInput() ([]Point, []Fold) {
	input, err := os.Open("./input.txt")
	check(err)
	defer input.Close()

	scanner := bufio.NewScanner(input)
	lines := 0

	// Read Coordinates
	var points []Point
	for scanner.Scan() {
		line := scanner.Text()
		lines++

		if line == "" {
			break
		}
		coords := strings.Split(line, ",")

		x, err := strconv.Atoi(coords[0])
		check(err)
		y, err := strconv.Atoi(coords[1])
		check(err)

		points = append(points, Point{x, y})
	}

	// Read Instructions
	const INSTRUCTION_PREFIX = "fold along "
	const INSTRUCTION_PREFIX_LEN = len(INSTRUCTION_PREFIX)
	var folds []Fold
	for scanner.Scan() {
		line := scanner.Text()
		lines++
		if strings.Contains(line, INSTRUCTION_PREFIX) {
			instruction := strings.Split(line[INSTRUCTION_PREFIX_LEN:], "=")
			direction := instruction[0]
			value, err := strconv.Atoi(instruction[1])
			check(err)

			folds = append(folds, Fold{direction, value})
		}
	}

	log.Printf(log.DEBUG, "Read %d coordinates and %d fold instructions from %d lines\n", len(points), len(folds), lines)

	return points, folds
}

func getMaxSize(points []Point) (int, int) {
	maxX, maxY := 0, 0
	for _, point := range points {
		if point.x > maxX {
			maxX = point.x
		}
		if point.y > maxY {
			maxY = point.y
		}
	}

	log.Printf(log.DIAGNOSTIC, "Max size is %d wide by %d tall\n", maxX, maxY)
	return maxX, maxY
}

type Grid = [][]bool

func initGrid(points []Point) Grid {
	maxX, maxY := getMaxSize(points)

	grid := make(Grid, maxY+1)
	for y := 0; y <= maxY; y++ {
		grid[y] = make([]bool, maxX+1)
	}

	for _, point := range points {
		grid[point.y][point.x] = true
	}
	log.Printf(log.DIAGNOSTIC, "Marked %d points on the grid\n", len(points))

	return grid
}

func applyFold(grid Grid, fold Fold) Grid {
	if len(grid) == 0 {
		panic("Cannot process empty grid")
	}

	log.Printf(log.DEBUG, "Applying fold {%s=%d} to grid\n", fold.direction, fold.value)

	startY, startX, dirY, dirX := 0, 0, 1, 1
	if fold.direction == "y" {
		startY = fold.value
		dirY = -1
	}
	if fold.direction == "x" {
		startX = fold.value
		dirX = -1
	}

	ovDots := 0
	newDots := 0

	for y := startY; y < len(grid); y++ {
		for x := startX; x < len(grid[y]); x++ {
			if grid[y][x] {
				// to flip a point, we basically move it as far 'left' or 'up' from the fold point as it was 'right' or 'down'
				newX, newY := 2*startX+dirX*x, 2*startY+dirY*y
				log.Printf(log.DIAGNOSTIC, "Drawing (%d, %d) to (%d, %d)", x, y, newX, newY)
				if grid[newY][newX] {
					ovDots++
					log.Println(log.DIAGNOSTIC, " Ov")
				} else {
					newDots++
					log.Println(log.DIAGNOSTIC, " (new)")
				}
				grid[newY][newX] = true
			}
		}
	}

	log.Printf(log.DIAGNOSTIC, "Overwrite %d dots and added %d dots\n", ovDots, newDots)

	// chop the grid either width- or length-wise as appropriate
	if fold.direction == "y" {
		grid = grid[:startY]
	} else if fold.direction == "x" {
		for y, row := range grid {
			grid[y] = row[:startX]
		}
	}

	return grid
}

// count the number of dots in the grid
func countDots(grid Grid) int {
	count := 0
	for _, row := range grid {
		for _, value := range row {
			if value {
				count++
			}
		}
	}
	return count
}

func printGrid(grid Grid) {
	log.Printf(log.NORMAL, "/%s\\\n", nTimesStr("=", len(grid[0])))
	for _, row := range grid {
		log.Print(log.NORMAL, "|")
		for _, value := range row {
			switch value {
			case false:
				log.Print(log.NORMAL, " ")
			case true:
				log.Print(log.NORMAL, "*")
			}
		}
		log.Println(log.NORMAL, "|")
	}
	log.Printf(log.NORMAL, "\\%s/\n", nTimesStr("=", len(grid[0])))
}

func nTimesStr(s string, t int) string {
	newStr := ""
	for c := 0; c < t; c++ {
		newStr += s
	}
	return newStr
}
