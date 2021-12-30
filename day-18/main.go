/* Day 18: Snailfish */
package main

import (
	"AoC2021/utils/args"
	"AoC2021/utils/log"
	"AoC2021/utils/timer"
	"bufio"
	"fmt"
	"math"
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
		log.Printf(log.DEBUG, "Line #%3d: ", lines+1)
		if lines == 0 {
			log.Printf(log.DEBUG, "%s\n", line)
			s = ParseSnumStr(line)
		} else {
			log.Printf(log.DEBUG, "+ %s\n", line)
			s = AddSnum(s, ParseSnumStr(line))
			log.Printf(log.DEBUG, "%10s = %s\n", "", s.ToString())
		}
		lines++
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

// Snailfish number
type Snum []int

const (
	LBracket int = -1
	RBracket int = -2
	Comma    int = -3
)

func ParseSnumStr(line string) Snum {
	var s Snum
	log.Printf(log.DIAGNOSTIC, "Starting to parse snum line \"%s\"\n", line)

	for idx := 0; idx < len(line); idx++ {
		switch line[idx] {
		case '[':
			s = append(s, LBracket)
		case ']':
			s = append(s, RBracket)
		case ',':
			s = append(s, Comma)
		default:
			// if it's not a bracket or a comma then it must be a number
			endIdx := idx + 1
			for ; line[endIdx] >= '0' && line[endIdx] <= '9'; endIdx++ {
			}

			strVal := line[idx:endIdx]
			iVal, err := strconv.Atoi(strVal)
			check(err)
			s = append(s, iVal)

			idx = endIdx - 1
		}
	}

	return s
}

func (s Snum) ToString() string {
	str := ""

	for _, i := range s {
		switch i {
		case LBracket:
			str += "["
		case RBracket:
			str += "]"
		case Comma:
			str += ","
		default: // must be a number
			str += fmt.Sprintf("%d", i)
		}
	}

	return str
}

func AddSnum(s1, s2 Snum) Snum {
	log.Printf(log.DIAGNOSTIC, "  %s\n+ %s\n", s1.ToString(), s2.ToString())

	s := append(Snum{LBracket}, s1...)
	s = append(s, Comma)
	s = append(s, s2...)
	s = append(s, RBracket)

	s = s.Simplify()
	log.Printf(log.DIAGNOSTIC, "= %s\n", s.ToString())

	return s
}

func (s1 Snum) Add(s2 Snum) Snum {
	log.Printf(log.DIAGNOSTIC, "  %s\n+ %s\n", s1.ToString(), s2.ToString())

	s := append(Snum{LBracket}, s1...)
	s = append(s, Comma)
	s = append(s, s2...)
	s = append(s, RBracket)

	s = s.Simplify()
	log.Printf(log.DIAGNOSTIC, "= %s\n", s.ToString())

	return s
}

func (s Snum) Simplify() Snum {
	depth := 0

	for idx := 0; idx < len(s); idx++ {
		switch s[idx] {
		case LBracket:
			depth++
			if s.ShouldExplodeAt(idx) {
				s.ExplodeAt(idx)
				idx = 0
			}
		case RBracket:
			depth--
		case Comma:
			continue
		default: // a number
			if s.ShouldSplitAt(idx) {
				s.SplitAt(idx)
				idx = 0
			}
		}
	}
	return s
}

func (s Snum) DepthAt(pos int) int {
	depth := 0
	for _, i := range s[:pos+1] {
		if i == LBracket {
			depth++
		} else if i == RBracket {
			depth--
		}
	}

	return depth
}

func (s Snum) IsNormalAt(pos int) bool {
	// the next few symbols should be "[#,#]"
	return pos+4 < len(s) && s[pos] == LBracket && s[pos+1] >= 0 && s[pos+3] >= 0
}

// the current read position and depth (for ease of calculation)
func (s Snum) ShouldExplodeAt(pos int) bool {
	return s.IsNormalAt(pos) && s.DepthAt(pos) > 4
}

func (s *Snum) ExplodeAt(pos int) {
	log.Printf(log.DIAGNOSTIC, "Exploding the snum at %d: %s\n", pos, s.ToString())

	lVal, rVal := (*s)[pos+1], (*s)[pos+3]
	// explode left!
	for idx := pos; idx > 0; idx-- {
		if (*s)[idx] >= 0 {
			(*s)[idx] += lVal
			break
		}
	}
	// explode right!
	for idx := pos + 4; idx < len(*s); idx++ {
		if (*s)[idx] >= 0 {
			(*s)[idx] += rVal
			break
		}
	}
	// replace the current snum with 0!
	*s = append((*s)[0:pos], append(Snum{0}, (*s)[pos+5:]...)...)
}

func (s Snum) ShouldSplitAt(pos int) bool {
	return s[pos] > 9
}

func (s *Snum) SplitAt(pos int) {
	lVal, rVal := int(math.Floor(float64((*s)[pos])/2)), int(math.Ceil(float64((*s)[pos])/2))
	log.Printf(log.DIAGNOSTIC, "Splitting the snum at %d: %d, adding values %d and %d\n", pos, (*s)[pos], lVal, rVal)

	// Add a few extra values into the slice and then set the values
	*s = append((*s)[:pos+4], (*s)[pos:]...)
	(*s)[pos] = LBracket
	(*s)[pos+1] = lVal
	(*s)[pos+2] = Comma
	(*s)[pos+3] = rVal
	(*s)[pos+4] = RBracket
}

func (s Snum) CalcMagnitude() int {
	for idx := 0; idx < len(s) && len(s) > 1; idx++ {
		if s.IsNormalAt(idx) {
			mVal := 3*s[idx+1] + 2*s[idx+3]
			log.Printf(log.DIAGNOSTIC, "Replacing %s with %d: ", s[idx:idx+5].ToString(), mVal)
			s = append(s[:idx+1], s[idx+5:]...)
			s[idx] = mVal
			idx = 0
			log.Printf(log.DIAGNOSTIC, "%s\n", s.ToString())
		}
	}
	// when there's only one value left, that's the answer
	return 3*s[1] + 2*s[3]
}
