package timing

import "time"

func TimeFunction(fn func(), trials int) time.Duration {
	startTime := time.Now()

	for i := 0; i < trials; i++ {
		fn()
	}
	endTime := time.Now()

	return endTime.Sub(startTime)
}
