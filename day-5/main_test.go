package main

import (
	"AoC2021/utils/log"
	"AoC2021/utils/timing"
	"fmt"
	"testing"
	"time"
)

const TRIALS = 1000

func TestDay5Part1(t *testing.T) {
	log.LOG_LEVEL = log.NONE

	duration := timing.TimeFunction(part1, TRIALS)
	avg := time.Duration(int64(duration) / int64(TRIALS))

	fmt.Printf("Part1: [Took %s across %d trials for an average of %s]\n", duration.String(), TRIALS, avg)
}

func TestDay5Part2(t *testing.T) {
	log.LOG_LEVEL = log.NONE

	duration := timing.TimeFunction(part2, TRIALS)
	avg := time.Duration(int64(duration) / int64(TRIALS))

	fmt.Printf("Part2: [Took %s across %d trials for an average of %s]\n", duration.String(), TRIALS, avg)
}
