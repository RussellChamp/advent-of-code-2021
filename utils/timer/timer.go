package timer

import (
	"AoC2021/utils/color"
	"AoC2021/utils/log"
	"fmt"
	"time"
)

var lastTime time.Time

func Start() {
	lastTime = time.Now()
}

func Tick() {
	now := time.Now()
	fmt.Printf(color.Yellow+"[Time taken: %s]\n"+color.Reset, now.Sub(lastTime).String())
	lastTime = now
}

func TickAtLevel(logLevel int) {
	if log.LOG_LEVEL >= logLevel {
		now := time.Now()
		log.Printf(logLevel, color.Yellow+"[Time taken: %s]\n"+color.Reset, now.Sub(lastTime).String())
		lastTime = now
	}
}
