/* Day 4: Giant Squid */
// TODO: Redo this whole thing with bit arrays instead
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
	fmt.Println("--- Day 4: Giant Squid ---")
	timer.Start()
	part1()
	timer.Tick()
	fmt.Println()

	part2()
	timer.Tick()
	fmt.Println()
}

type Card = [][]uint8

func readData(input *os.File) ([]uint8, []Card) {
	scanner := bufio.NewScanner(input)

	var draws []uint8

	if scanner.Scan() {
		for _, val := range strings.Split(scanner.Text(), ",") {
			iVal, err := strconv.Atoi(val)
			check(err)

			draws = append(draws, uint8(iVal))
		}
	}

	var cards []Card
	cardIdx := -1 // the first new line will initialize card 0
	rowIdx := 0

	// fmt.Printf("Created cards %s", cards)
	for scanner.Scan() {
		line := scanner.Text()
		switch {
		case line == "": // when we hit an epty line, initialize a new card
			cards = append(cards, Card{})
			cardIdx += 1
			rowIdx = 0
			//fmt.Printf("Added card %d to cards {%s}\n", cardIdx, cards)
		default:
			rowVals := strings.Fields(line)
			var newRow []uint8
			for _, strVal := range rowVals {
				val, err := strconv.Atoi(strVal)
				check(err)
				newRow = append(newRow, uint8(val))
			}
			//fmt.Printf("Attempting to push [%s] to card %d row %d {%s}\n", line, cardIdx, rowIdx, cards[cardIdx])
			cards[cardIdx] = append(cards[cardIdx], newRow)
			rowIdx += 1
		}
	}

	return draws, cards
}

type MarkedCard = [][]bool

func initializeMarks(cards []Card) []MarkedCard {
	// create a number of marked cards equal to the number of bingo cards
	cardsLen := len(cards)
	marks := make([]MarkedCard, cardsLen)
	// find the number of rows on a card
	rowCount := len(cards[0])
	colCount := len(cards[0][0])

	for cIdx := 0; cIdx < cardsLen; cIdx++ {
		marks[cIdx] = make([][]bool, rowCount)
		for rIdx := 0; rIdx < rowCount; rIdx++ {
			//fmt.Printf("Attempting to add row %d to card %d", rIdx, cIdx)
			//fmt.Println(marks[cIdx])
			marks[cIdx][rIdx] = make([]bool, colCount)
		}
	}

	return marks
}

func markCards(cards []Card, marks []MarkedCard, draw uint8) []MarkedCard {
	for cIdx, card := range cards {
		for rowIdx, row := range card {
			for colIdx, value := range row {
				if value == draw {
					marks[cIdx][rowIdx][colIdx] = true
				}
			}
		}
	}

	return marks
}

// TODO: currently assumes only one card will win per round
func findWinningCardIdx(marks []MarkedCard) int {
	for cIdx, card := range marks {
		if matchesRow(card) || matchesColumn(card) {
			return cIdx
		}
	}

	return -1
}

func matchesRow(card MarkedCard) bool {
	for _, row := range card {
		marks := 0
		for _, value := range row {
			if value {
				marks += 1
			}
		}
		if marks == len(row) {
			return true
		}
	}
	return false
}

func matchesColumn(card MarkedCard) bool {
	rowCount := len(card)
	colCount := len(card[0])
	for colIdx := 0; colIdx < colCount; colIdx++ {
		marks := 0
		for rowIdx := 0; rowIdx < rowCount; rowIdx++ {
			if card[rowIdx][colIdx] {
				marks += 1
			}
		}

		if marks == colCount {
			return true
		}
	}
	return false
}

func getCardScore(card Card, mark MarkedCard) int {
	score := 0

	// sum up all the non-marked bingo spots
	for rIdx, row := range card {
		for cIdx, value := range row {
			if !mark[rIdx][cIdx] {

				score += int(value)
			}
		}
	}
	return score
}

func part1() {
	fmt.Println("* Part 1 *")
	fmt.Println(" Goal: Figure out which bingo board will win first.")
	fmt.Println(" Answer: What will your final score be if you choose that board?")

	input, err := os.Open("./input.txt")
	check(err)
	defer input.Close()

	draws, cards := readData(input)
	marks := initializeMarks(cards)

	//fmt.Printf("Read in %d draws and %d cards\n", len(draws), len(cards))
	//fmt.Printf("Initialized %d mark cards\n", len(marks))
	//fmt.Printf("First card %s\n", cards[0])
	//fmt.Println(marks[0])

	drawTimes := 0
	winningCardIdx := -1
	lastDraw := uint8(0)
	score := 0
	for _, draw := range draws {
		drawTimes += 1
		marks = markCards(cards, marks, draw)
		winningCardIdx = findWinningCardIdx(marks)
		if winningCardIdx >= 0 {
			winningCard := cards[winningCardIdx]
			winningMarks := marks[winningCardIdx]
			score = getCardScore(winningCard, winningMarks)

			// multiply by the last drawn tile

			lastDraw = draw
			score *= int(draw)

			break
		}
	}

	fmt.Printf("After %d draws, card #%d won on a %d with a score of %d\n", drawTimes, winningCardIdx, lastDraw, score)
}

func part2() {
	fmt.Println("* Part 2 *")
	fmt.Println(" Goal: Figure out which board will win last.")
	fmt.Println(" Answer: Once it wins, what would its final score be?")

	input, err := os.Open("./input.txt")
	check(err)
	defer input.Close()

	draws, cards := readData(input)
	marks := initializeMarks(cards)

	drawTimes := 0
	lastDraw := uint8(0)
	score := 0
	foundLoser := false

	for _, draw := range draws {
		if foundLoser {
			break
		}

		drawTimes += 1
		marks = markCards(cards, marks, draw)
		winningCardIdx := findWinningCardIdx(marks)

		eliminatedCards := 0
		for winningCardIdx > -1 {
			// if there's only one card left, it's the losingist one
			if len(marks) == 1 {
				score = getCardScore(cards[0], marks[0])

				// multiply by the last drawn tile

				lastDraw = draw
				score *= int(draw)
				foundLoser = true
				break
			}

			eliminatedCards += 1
			// if a card wins, just remove it from the list of cards and don't track it anymore
			cards[winningCardIdx] = cards[len(cards)-1]
			cards[len(cards)-1] = Card{}
			cards = cards[:len(cards)-1]

			marks[winningCardIdx] = marks[len(marks)-1]
			marks[len(marks)-1] = MarkedCard{}
			marks = marks[:len(marks)-1]

			winningCardIdx = findWinningCardIdx(marks)
		}
		if eliminatedCards > 0 {
			//fmt.Printf("Round %d: Eliminated %d cards with a %s (%d remaining)\n", drawTimes, eliminatedCards, draw, len(cards))
		}
	}

	//fmt.Println("Card: ", cards[0])
	//fmt.Println("Mark: ", marks[0])
	fmt.Printf("After %d draws, the losingist card won on a %d with a score of %d\n", drawTimes, lastDraw, score)
}
