package timer

import (
	"fmt"
	"time"
)

var lastTime time.Time

func Start() {
	lastTime = time.Now()
}

func Tick() {
	now := time.Now()
	fmt.Printf("[Time taken: %s]\n", now.Sub(lastTime).String())
	lastTime = now
}
