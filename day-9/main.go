/* Day 9: Smoke Basin */
package main

import (
	"AoC2021/utils/arrays"
	"AoC2021/utils/log"
	"AoC2021/utils/timer"
	"bufio"
	"os"
	"sort"
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
	log.Println(log.NORMAL, "--- Day 9: Smoke Basin ---")
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
	log.Println(log.NORMAL, " Goal: Find all of the low points on your heightmap.")
	log.Println(log.NORMAL, " Answer: What is the sum of the risk levels of all low points on your heightmap?")

	points, lines := loadPoints()
	minPoints := findMinPoints(points)

	riskLevelSum := 0
	for _, p := range minPoints {
		riskLevelSum += getRiskLevel(points, p)
	}

	log.Printf(log.NORMAL, "Read %d lines from input\n", lines)
	log.Printf(log.DIAGNOSTIC, "Min points are:")
	for _, minPoint := range minPoints {
		log.Printf(log.DIAGNOSTIC, " (%d, %d)", minPoint.col, minPoint.row)
	}
	log.Println(log.DIAGNOSTIC)
	log.Printf(log.NORMAL, "The risk point total of all min points is %d\n", riskLevelSum)
}

func part2() {
	log.Println(log.NORMAL, "* Part 2 *")
	log.Println(log.NORMAL, " Goal: Find the three largest basins.")
	log.Println(log.NORMAL, " Answer: What do you get if you multiply together the sizes of the three largest basins?")

	points, lines := loadPoints()
	minPoints := findMinPoints(points)
	basins := getBasins(points, minPoints)
	sort.Slice(basins, func(i, j int) bool { return basins[j].size < basins[i].size })

	log.Printf(log.NORMAL, "Read %d lines from input\n", lines)
	if len(basins) < 3 {
		panic("expected to find at least 3 basins!")
	}
	log.Printf(log.NORMAL, "The three largest basins are")
	for _, b := range basins[:3] {
		log.Printf(log.NORMAL, " (%d, %d): %d", b.minPoint.col, b.minPoint.row, b.size)
	}
	log.Println(log.NORMAL)
	totalValue := basins[0].size * basins[1].size * basins[2].size
	log.Printf(log.NORMAL, "The multiplied value is %d\n", totalValue)
}

func loadPoints() ([][]int, int) {
	input, err := os.Open("./input.txt")
	check(err)
	defer input.Close()

	scanner := bufio.NewScanner(input)
	lines := 0
	var points [][]int

	for scanner.Scan() {
		// Read in a line and convert each character into an int in the array
		values := arrays.MapStrToInt(strings.Split(scanner.Text(), ""), func(s string) int { i, _ := strconv.Atoi(s); return i })
		points = append(points, values)
		log.Printf(log.DIAGNOSTIC, "Line %d: Read in %d values\n", lines+1, len(values))
		lines++
	}

	return points, lines
}

type Point struct {
	row, col int
}

func findMinPoints(points [][]int) []Point {
	var minPoints []Point
	for rowIdx, row := range points {
		for colIdx := range row {
			if isMinPoint(points, rowIdx, colIdx) {
				minPoints = append(minPoints, Point{rowIdx, colIdx})
			}
		}
	}

	return minPoints
}

// check if the passed point actually exists within the list
func isPointValid(points [][]int, row, col int) bool {
	if row < 0 || row >= len(points) {
		return false
	}
	if col < 0 || col >= len(points[row]) {
		return false
	}

	return true
}

// check if the current point is a minimum
func isMinPoint(points [][]int, row, col int) bool {
	// it is a minimum point if all valid points around it have a higher value
	value := points[row][col]
	// LEFT
	if isPointValid(points, row, col-1) && points[row][col-1] <= value {
		return false
	}
	// RIGHT
	if isPointValid(points, row, col+1) && points[row][col+1] <= value {
		return false
	}
	// UP
	if isPointValid(points, row-1, col) && points[row-1][col] <= value {
		return false
	}
	// DOWN
	if isPointValid(points, row+1, col) && points[row+1][col] <= value {
		return false
	}

	return true
}

func getRiskLevel(points [][]int, point Point) int {
	return 1 + points[point.row][point.col]
}

type Basin struct {
	minPoint Point
	size     int
}

func getBasins(points [][]int, minPoints []Point) []Basin {
	var basins []Basin
	for _, minPoint := range minPoints {
		size := getBasinSize(points, minPoint)
		basins = append(basins, Basin{minPoint, size})
	}

	return basins
}

func getBasinSize(points [][]int, minPoint Point) int {
	// to calculate the size of a basin
	// * start at the min point
	basinPoints := []Point{minPoint}

	// * recursively check all adjacent points
	expandBasin(&basinPoints, minPoint, points)

	// * if the point is LESS THAN 9, NOT OFF THE EDGE, and NOT ALREADY IN THE BAISN
	// *   then add it to the basin and recursively keep going!
	return len(basinPoints)
}

// returns whether the point should be added to a basin
// note: this is a naive and doesn't check that the point should actually exist in the given basin but rather if it meets a few criteria
func addToBasin(allPoints [][]int, basinPoints []Point, row, col int) bool {
	// only add valid points
	if !isPointValid(allPoints, row, col) {
		return false
	}
	// that are not already in the list
	if contains(basinPoints, Point{row, col}) {
		return false
	}
	// and have a value less than 9
	return allPoints[row][col] < 9
}

func contains(points []Point, point Point) bool {
	for _, p := range points {
		if p.row == point.row && p.col == point.col {
			return true
		}
	}
	return false
}

func expandBasin(basinPoints *[]Point, fromPoint Point, allPoints [][]int) {
	// LEFT
	if addToBasin(allPoints, *basinPoints, fromPoint.row, fromPoint.col-1) {
		newPoint := Point{fromPoint.row, fromPoint.col - 1}
		*basinPoints = append(*basinPoints, newPoint)
		expandBasin(basinPoints, newPoint, allPoints)
	}
	// RIGHT
	if addToBasin(allPoints, *basinPoints, fromPoint.row, fromPoint.col+1) {
		newPoint := Point{fromPoint.row, fromPoint.col + 1}
		*basinPoints = append(*basinPoints, newPoint)
		expandBasin(basinPoints, newPoint, allPoints)
	}
	// UP
	if addToBasin(allPoints, *basinPoints, fromPoint.row-1, fromPoint.col) {
		newPoint := Point{fromPoint.row - 1, fromPoint.col}
		*basinPoints = append(*basinPoints, newPoint)
		expandBasin(basinPoints, newPoint, allPoints)
	}
	// DOWN
	if addToBasin(allPoints, *basinPoints, fromPoint.row+1, fromPoint.col) {
		newPoint := Point{fromPoint.row + 1, fromPoint.col}
		*basinPoints = append(*basinPoints, newPoint)
		expandBasin(basinPoints, newPoint, allPoints)
	}
}
