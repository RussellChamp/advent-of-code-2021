/* Day 15: Chiton */
package main

import (
	"AoC2021/utils/arrays"
	"AoC2021/utils/color"
	"AoC2021/utils/log"
	"AoC2021/utils/numbers"
	"AoC2021/utils/timer"
	"bufio"
	"math"
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
	log.Println(log.NORMAL, "--- Day 15: Chiton ---")
	timer.Start()
	part1()
	timer.Tick()
	log.Println(log.NORMAL)

	// TODO: Write a faster algorithm
	// part2()
	// timer.Tick()
	// log.Println(log.NORMAL)
}

func part1() {
	log.Println(log.NORMAL, "* Part 1 *")
	log.Printf(log.NORMAL, color.Green)
	log.Println(log.NORMAL, " Goal: Your goal is to find a path with the lowest total risk.")
	log.Println(log.NORMAL, " Answer: What is the lowest total risk of any path from the top left to the bottom right?")
	log.Printf(log.NORMAL, color.Reset)

	risks := readRiskNodes()
	costs := initCosts(risks)

	calcCosts(risks, &costs)
	lastRow := len(costs) - 1
	lastCol := len(costs[lastRow]) - 1
	exitCost := costs[lastRow][lastCol] - costs[0][0]

	log.Printf(log.NORMAL, "The lowest risk cost from row=0 col=0 to row=%d col=%d is %d\n", lastRow, lastCol, exitCost)
	walkPath(risks, costs)
}

func part2() {
	log.SetLogLevel(log.DEBUG)

	log.Println(log.NORMAL, "* Part 2 *")
	log.Printf(log.NORMAL, color.Green)
	log.Println(log.NORMAL, " Goal: Calc using the full map")
	log.Println(log.NORMAL, " Answer: What is the lowest total risk of any path from the top left to the bottom right?")
	log.Printf(log.NORMAL, color.Reset)

	risks := readRiskNodes()
	expandRisks(&risks)
	costs := initCosts(risks)

	calcCosts(risks, &costs)
	lastRow := len(costs) - 1
	lastCol := len(costs[lastRow]) - 1
	exitCost := costs[lastRow][lastCol] - costs[0][0]

	log.Printf(log.NORMAL, "The lowest risk cost from row=0 col=0 to row=%d col=%d is %d\n", lastRow, lastCol, exitCost)
	walkPath(risks, costs)
}

type RiskNodes [][]int

func readRiskNodes() RiskNodes {
	input, err := os.Open("./input.txt")
	check(err)
	defer input.Close()

	scanner := bufio.NewScanner(input)
	lines := 0

	var risks RiskNodes

	for scanner.Scan() {
		line := scanner.Text()
		lines++
		// We'll just hope that our input is well formed
		risks = append(risks, arrays.MapStrToInt(strings.Split(line, ""), func(s string) int { i, _ := strconv.Atoi(s); return i }))
	}

	log.Printf(log.DEBUG, "Read %d lines from input\n", lines)
	return risks
}

func expandRisks(risks *RiskNodes) {
	multiplier := 5
	// expand out all the lines to the new width
	rowLen, colLen := len(*risks), len((*risks)[0])

	// extend each row N times
	for rIdx, row := range *risks {
		for m := 1; m < multiplier; m++ {
			(*risks)[rIdx] = append((*risks)[rIdx], row...)
		}
	}
	// copy each row N times
	for m := 1; m < multiplier; m++ {
		for rIdx := 0; rIdx < rowLen; rIdx++ {
			(*risks) = append((*risks), (*risks)[0])
		}
	}
	// increment the risk number in each cell
	for rIdx, row := range *risks {
		for cIdx, value := range row {
			if rIdx < rowLen && cIdx < colLen {
				continue
			}
			// add 1 for every section right or down that we've copied
			newVal := (value + rIdx/rowLen + cIdx/colLen) % 10
			(*risks)[rIdx][cIdx] = newVal
		}
	}
	log.Println(log.DIAGNOSTIC, "Expanded risk nodes to new size")
}

type CostNodes [][]int

func initCosts(risks RiskNodes) CostNodes {
	var costs CostNodes
	for _, row := range risks {
		costs = append(costs, make([]int, len(row)))
	}

	log.Printf(log.DIAGNOSTIC, "Initialized costs node of size row=%d by col=%d\n", len(risks), len(risks[0]))

	return costs
}

var filledNodes int
var recalcCount int
var totalNodes int

func calcCosts(risks RiskNodes, costs *CostNodes) {
	// the goal is to MINIMIZE the risk cost from traveling from START to END
	// starting from the current node check each adjacent valid node
	// if the total cost on that node is 0 or if the toal cost is > the current node total cost + next node travel cost
	//   then update that next node's total and continue iterating at it
	//   otherwise just stop
	(*costs)[0][0] = math.MaxInt // so that we don't go back to start

	filledNodes = 1
	totalNodes = len(*costs) * len((*costs)[0])
	log.Println(log.DIAGNOSTIC, "Starting calc from (0,0)")
	calcAllNeighbors(risks, costs, 0, 0, 0)
}

type RowCol struct {
	r int
	c int
}

func calcAllNeighbors(risks RiskNodes, costs *CostNodes, row, col, currentTotal int) {
	if !nodeIsValid(*costs, row, col) {
		return
	}
	log.Printf(log.DIAGNOSTIC, "Calculating neighbors at row=%d col=%d\n", row, col)
	if (*costs)[row][col] == 0 || (*costs)[row][col] > currentTotal+risks[row][col] {
		if (*costs)[row][col] == 0 {
			filledNodes++
			if filledNodes%1000 == 0 {
				log.Printf(log.DEBUG, "Calculated %d / %d nodes\n", filledNodes, totalNodes)
			}
		} else {
			recalcCount++
			if recalcCount%10000 == 0 {
				log.Printf(log.DEBUG, "Recalculated %d / ?? nodes\n", recalcCount)
			}
		}
		(*costs)[row][col] = currentTotal + risks[row][col]
		/* UP    */ calcAllNeighbors(risks, costs, row-1, col, currentTotal+risks[row][col])
		/* DOWN  */ calcAllNeighbors(risks, costs, row+1, col, currentTotal+risks[row][col])
		/* LEFT  */ calcAllNeighbors(risks, costs, row, col-1, currentTotal+risks[row][col])
		/* RIGHT */ calcAllNeighbors(risks, costs, row, col+1, currentTotal+risks[row][col])
	} else {
		log.Printf(log.DIAGNOSTIC, "Hit a calc dead end at row=%d, col=%d\n", row, col)
	}
}

func nodeIsValid(costs CostNodes, row, col int) bool {
	switch {
	case col < 0, col >= len(costs):
		return false
	case row < 0, row >= len(costs[col]):
		return false
	default:
		return true
	}
}

func walkPath(risks RiskNodes, costs CostNodes) {
	// start from the end and go toward the beginning then flip the script
	row := len(costs) - 1
	col := len(costs[row]) - 1

	log.Printf(log.DIAGNOSTIC, "Starting to walk path at row=%d, col=%d\n", row, col)

	nodes := []RowCol{{row, col}}

	sum := 0
	for ; row != 0 || col != 0; row, col = lowestNeighbor(costs, row, col) {
		nodes = append(nodes, RowCol{row, col})
		sum += risks[row][col]
	}
	nodes = append(nodes, RowCol{row, col})
	// don't count 0,0 in the sum

	// print all the nodes and if they are on the path then color them
	printColor := color.DarkGray

	for row := 0; row < len(risks); row++ {
		for col := 0; col < len(risks[row]); col++ {
			// for the first and last node, print in yellow
			// for nodes on the path, print in green
			// else print in normal
			if (row == 0 && col == 0) || (row == len(risks)-1 && col == len(risks[row])-1) {
				printColor = color.YellowBold
				log.Printf(log.NORMAL, printColor)
			} else if inPath(nodes, row, col) {
				if printColor != color.BlueBold {
					printColor = color.BlueBold
					log.Printf(log.NORMAL, printColor)
				}
			} else if printColor != color.DarkGray {
				printColor = color.DarkGray
				log.Printf(log.NORMAL, printColor)
			}

			log.Printf(log.NORMAL, "%d", risks[row][col])
		}
		log.Println(log.NORMAL)
	}
	log.Printf(log.NORMAL, color.Reset)

	// iterate through backward to go from start node to end
	row, col = nodes[len(nodes)-1].r, nodes[len(nodes)-1].c
	log.Printf(log.NORMAL, color.Blue+"%d (%03d)"+color.Reset, risks[row][col], costs[row][col])
	for nIdx := len(nodes) - 2; nIdx >= 0; nIdx-- {
		row, col = nodes[nIdx].r, nodes[nIdx].c
		log.Printf(log.NORMAL, " -> "+color.Blue+"%d (%03d)"+color.Reset, risks[row][col], costs[row][col])
		if (len(nodes)-nIdx)%16 == 0 {
			log.Println(log.NORMAL)
		}
	}
	log.Println(log.NORMAL)

	log.Printf(log.DEBUG, "Total risk cost is %d in %d steps\n", sum, len(nodes))
}

func lowestNeighbor(costs CostNodes, row, col int) (int, int) {
	// if we're at 0,1 or 1,0 then just go to 0,0
	if (row == 0 && col == 1) || (row == 1 && col == 0) {
		return 0, 0
	}

	// compare UP, DOWN, LEFT, RIGHT
	upCost, downCost, leftCost, rightCost := math.MaxInt, math.MaxInt, math.MaxInt, math.MaxInt
	if nodeIsValid(costs, row-1, col) {
		upCost = costs[row-1][col]
	}
	if nodeIsValid(costs, row+1, col) {
		downCost = costs[row+1][col]
	}
	if nodeIsValid(costs, row, col-1) {
		leftCost = costs[row][col-1]
	}
	if nodeIsValid(costs, row, col+1) {
		rightCost = costs[row][col+1]
	}

	switch numbers.Min(upCost, downCost, leftCost, rightCost) {
	case upCost:
		return row - 1, col
	case downCost:
		return row + 1, col
	case leftCost:
		return row, col - 1
	case rightCost:
		return row, col + 1
	default:
		panic("at the disco!")
	}
}

func inPath(nodes []RowCol, row, col int) bool {
	for _, n := range nodes {
		if n.r == row && n.c == col {
			return true
		}
	}
	return false
}
