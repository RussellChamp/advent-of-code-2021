/* Day 18: Snailfish */
package main

import (
	"AoC2021/utils/args"
	"AoC2021/utils/log"
	"AoC2021/utils/timer"
	"bufio"
	"fmt"
	"os"
	"strconv"
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

	log.Println(log.NORMAL, "--- Day 18: Snailfish ---")
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
	log.Println(log.NORMAL, " Goal: Add up all of the snailfish numbers from the homework assignment in the order they appear.")
	log.Println(log.NORMAL, " Answer: What is the magnitude of the final sum?")

	input, err := os.Open("./input.txt")
	check(err)
	defer input.Close()

	scanner := bufio.NewScanner(input)
	lines := 0
	var s Snum

	for scanner.Scan() {
		line := scanner.Text()
		if lines == 0 {
			s = ParseSnumStr(line)
		} else {
			s = AddSnum(s, ParseSnumStr(line))
		}

		// fmt.Printf("SNum #%d: %s\n", lines+1, s.ToString())

		lines++
		if lines > 1 {
			break
		}
	}

	log.Printf(log.NORMAL, "Read %d lines from input\n", lines)
	log.Printf(log.NORMAL, "The solution is: %s\n", s.ToString())
	log.Printf(log.NORMAL, " with a magnitude of %d\n", s.CalcMagnitude())
}

func part2() {
	log.Println(log.NORMAL, "* Part 2 *")
	log.Println(log.NORMAL, " Goal: ")
	log.Println(log.NORMAL, " Answer: ")
}

type Direction int

const (
	LEFT  Direction = 0
	RIGHT Direction = 1
	UP    Direction = 2
)

// Snailfish number
// consists of a pair of numbers
// the "left" and "right" side can consist of EITHER a normal value or another pair
type Snum struct {
	lVal      int
	rVal      int
	lPtr      *Snum
	rPtr      *Snum
	parentPtr *Snum
	dir       Direction
}

func ParseSnumStr(line string) Snum {
	var s Snum
	var sPtr = &s
	log.Printf(log.DIAGNOSTIC, "Starting to parse snum line \"%s\"\n", line)

	// start on idx 1 and skip the initial open bracket
	for idx := 1; idx < len(line); idx++ {
		switch line[idx] {
		case '[':
			if sPtr.dir == LEFT {
				// create a new snum on the left side
				sPtr.lPtr = &Snum{parentPtr: sPtr}
				sPtr.dir = RIGHT
				sPtr = sPtr.lPtr
			} else {
				sPtr.rPtr = &Snum{parentPtr: sPtr}
				sPtr = sPtr.rPtr
			}
		case ']':
			// go up a level
			sPtr.dir = LEFT
			sPtr = sPtr.parentPtr
		case ',':
			continue
		default:
			// if it's not a bracket or a comma then it must be a number
			endIdx := idx + 1
			for ; line[endIdx] >= '0' && line[endIdx] <= '9'; endIdx++ {
			}

			strVal := line[idx:endIdx]
			iVal, err := strconv.Atoi(strVal)
			check(err)

			if sPtr.dir == LEFT {
				sPtr.lVal = iVal
				sPtr.dir = RIGHT
			} else {
				sPtr.rVal = iVal
			}
			idx = endIdx - 1

		}
	}

	return s
}

func (s Snum) ToString() string {
	str := "["
	if s.lPtr != nil {
		str += s.lPtr.ToString()
	} else {
		str += fmt.Sprintf("%d", s.lVal)
	}
	str += ","
	if s.rPtr != nil {
		str += s.rPtr.ToString()
	} else {
		str += fmt.Sprintf("%d", s.rVal)
	}
	str += "]"

	return str
}

func AddSnum(s1, s2 Snum) Snum {
	log.Printf(log.DEBUG, "  %s\n+ %s\n", s1.ToString(), s2.ToString())
	s := Snum{lPtr: &s1, rPtr: &s2}
	s = s.Simplify()
	log.Printf(log.DEBUG, "= %s\n", s.ToString())

	return s
}

func (s1 Snum) Add(s2 Snum) Snum {
	log.Printf(log.DEBUG, "  %s\n+ %s\n", s1.ToString(), s2.ToString())
	s := Snum{lPtr: &s1, rPtr: &s2}
	s = s.Simplify()
	log.Printf(log.DEBUG, "= %s\n", s.ToString())

	return s
}

func (s Snum) Simplify() Snum {
	madeAChange := true
	for depth := 0; madeAChange; {
		madeAChange = false
		for sPtr := &s; sPtr != nil; {
			switch {
			case sPtr.ShouldExplode(depth):
				ExplodeValue(sPtr)
				madeAChange = true
				sPtr = nil

			// case sPtr.shouldSplit()

			// travel the left side
			case sPtr.dir == LEFT:
				sPtr.dir = RIGHT
				if sPtr.lPtr != nil {
					sPtr = sPtr.lPtr
					depth++
				}
			// travel the right side
			case sPtr.dir == RIGHT:
				sPtr.dir = UP
				if sPtr.rPtr != nil {
					sPtr = sPtr.rPtr
					depth++
				}
			case sPtr.dir == UP:
				sPtr = sPtr.parentPtr
				depth--
			}
		}
	}

	// check to explode
	// check to split
	return s
}

func (s Snum) IsNormal() bool {
	return s.lPtr == nil && s.rPtr == nil
}

func (s Snum) ShouldExplode(depth int) bool {
	return s.IsNormal() && depth > 3
}

func ExplodeValue(at *Snum) {
	log.Printf(log.DIAGNOSTIC, "Exploding the snum at %s\n", at.ToString())
	if !at.IsNormal() {
		panic("tried to explode a non-normal snum")
	}
	lVal, rVal := at.lVal, at.rVal
	var sPtr *Snum
	// explode left
	log.Printf(log.NEVER, "Starting at %s\n", at.parentPtr.ToString())
	for sPtr = at.parentPtr; sPtr.parentPtr != nil && sPtr.lPtr != nil; sPtr = sPtr.parentPtr {
		log.Printf(log.NEVER, " traveled left to %s\n", sPtr.ToString())
	}
	if sPtr != nil {
		log.Printf(log.NEVER, " added left %d to %d\n", lVal, sPtr.lVal)
		sPtr.lVal += lVal
	}

	// explode right
	for sPtr = at.parentPtr; sPtr.parentPtr != nil && sPtr.rPtr != nil; {
		log.Printf(log.NEVER, " traveled right to %s\n", sPtr.ToString())
		sPtr = sPtr.parentPtr
	}
	if sPtr != nil {
		log.Printf(log.NEVER, " added right %d to %d\n", rVal, sPtr.rVal)
		sPtr.rVal += rVal
	}

	// delete yourself
	pPtr := at.parentPtr
	if pPtr.lPtr == at {
		pPtr.lPtr = nil
	} else if pPtr.rPtr == at {
		pPtr.rPtr = nil
	}
}

func (s Snum) ShouldSplit() bool {
	return s.lVal > 9 || s.rVal > 9
}

func (s Snum) CalcMagnitude() int {
	lVal, rVal := s.lVal, s.rVal
	if s.lPtr != nil {
		lVal = s.lPtr.CalcMagnitude()
	}
	if s.rPtr != nil {
		rVal = s.rPtr.CalcMagnitude()
	}

	return 3*lVal + 2*rVal
}
