/* Day 10: Syntax Scoring */
package main

import (
	"AoC2021/utils/arrays"
	"AoC2021/utils/log"
	"AoC2021/utils/timer"
	"bufio"
	"os"
	"sort"
	"strings"
)

func check(e error) {
	if e != nil {
		log.Println(log.DEBUG, "Oh snap!")
		panic(e)
	}
}

func main() {
	log.Println(log.NORMAL, "--- Day 10: Syntax Scoring ---")
	timer.Start()
	part1()
	timer.Tick()
	log.Println(log.NORMAL)

	part2()
	timer.Tick()
	log.Println(log.NORMAL)
}

func part1() {
	log.LOG_LEVEL = log.DEBUG
	log.Println(log.NORMAL, "* Part 1 *")
	log.Println(log.NORMAL, " Goal: Find the first illegal character in each corrupted line of the navigation subsystem.")
	log.Println(log.NORMAL, " Answer: What is the total syntax error score for those errors?")

	input, err := os.Open("./input.txt")
	check(err)
	defer input.Close()

	scanner := bufio.NewScanner(input)
	lines := 0
	tokenCount := 0
	totalScore := 0

	for scanner.Scan() {
		line := scanner.Text()
		_, errScore := validateNavLine(line)
		totalScore += errScore
		lines++
		tokenCount += len(line)
	}

	log.Printf(log.NORMAL, "Read %d lines with %d tokens from input\n", lines, tokenCount)
	log.Printf(log.NORMAL, "The total syntax error score is %d\n", totalScore)
}

func part2() {
	log.Println(log.NORMAL, "* Part 2 *")
	log.Println(log.NORMAL, " Goal: Find the completion string for each incomplete line, score the completion strings, and sort the scores.")
	log.Println(log.NORMAL, " Answer: What is the middle score?")

	input, err := os.Open("./input.txt")
	check(err)
	defer input.Close()

	scanner := bufio.NewScanner(input)
	lines := 0
	tokenCount := 0

	var autoScores []int

	for scanner.Scan() {
		line := scanner.Text()

		lines++
		tokenCount += len(line)

		remainingStr, errScore := validateNavLine(line)
		if errScore > 0 || len(remainingStr) == 0 {
			continue
		}

		autoScores = append(autoScores, getAutoScore(remainingStr))
	}

	// stort the scores
	sort.Slice(autoScores, func(i, j int) bool { return autoScores[i] < autoScores[j] })
	// find the middle-ist index
	middleIdx := len(autoScores) / 2

	log.Printf(log.NORMAL, "Read %d lines with %d tokens from input\n", lines, tokenCount)
	log.Printf(log.NORMAL, "Autocompleted %d lines\n", len(autoScores))
	log.Printf(log.NORMAL, "The middle-ist auto-correct score at pos %d is %d\n", middleIdx, autoScores[middleIdx])
}

// validates the nav line and returns a non-zero value as a syntax error score
func BAD_validateNavLine(line string) int {
	// split the line into tokens
	tokens := strings.Split(line, "")
	invalidStr := "An invalid token '%s' was found at position %3d while validating string: %s\n"

	// I thought of doing this generically but there are only 4 hard-coded options so just do it manually
	parens, brackets, braces, gators := 0, 0, 0, 0

	for idx, token := range tokens {
		switch token {
		case "(":
			parens++
		case "[":
			brackets++
		case "{":
			braces++
		case "<":
			gators++
		case ")":
			if parens > 0 {
				parens--
			} else {
				log.Printf(log.DEBUG, invalidStr, token, idx, line)
				return 3
			}
		case "]":
			if brackets > 0 {
				brackets--
			} else {
				log.Printf(log.DEBUG, invalidStr, token, idx, line[:idx+1])
				return 57
			}
		case "}":
			if braces > 0 {
				braces--
			} else {
				log.Printf(log.DEBUG, invalidStr, token, idx, line)
				return 1197
			}
		case ">":
			if gators > 0 {
				gators--
			} else {
				log.Printf(log.DEBUG, invalidStr, token, idx, line)
				return 25137
			}
		}
	}

	return 0
}

func validateNavLine(line string) (string, int) {
	reducedLine := reduceNavLine(line)
	if len(reducedLine) == 0 {
		return reducedLine, 0
	}
	// find the first closing token
	firstCloseTokenIdx := strings.IndexAny(reducedLine, ")]}>")
	// if no closing tokens are found, this is just an incomplete line (not corrupt)
	if firstCloseTokenIdx == -1 {
		return reducedLine, 0
	}
	log.Printf(log.DIAGNOSTIC, "First invalid token is at %d on %s\n", firstCloseTokenIdx, reducedLine)
	switch reducedLine[firstCloseTokenIdx] {
	case ')':
		return reducedLine, 3
	case ']':
		return reducedLine, 57
	case '}':
		return reducedLine, 1197
	case '>':
		return reducedLine, 25137
	}

	return reducedLine, 0
}

func reduceNavLine(line string) string {
	//recursively remove empty chunks from the line
	oldLen := len(line)
	for {
		line = strings.ReplaceAll(line, "()", "")
		line = strings.ReplaceAll(line, "[]", "")
		line = strings.ReplaceAll(line, "{}", "")
		line = strings.ReplaceAll(line, "<>", "")

		if len(line) == oldLen {
			break
		}
		oldLen = len(line)
	}

	return line
}

// get an auto-complete score from a line that has already been reduced
func getAutoScore(reducedLine string) int {
	// we'll have to auto-complete tokens in REVERSE order of the opening tokens in the reducedLine
	score := 0
	for _, token := range arrays.Reverse(strings.Split(reducedLine, "")) {
		score *= 5
		switch token {
		case "(":
			score += 1
		case "[":
			score += 2
		case "{":
			score += 3
		case "<":
			score += 4
		}
	}

	return score
}
