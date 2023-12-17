package main

import (
	"go-limits/bucket"
	"go-limits/sliding"
	"go-limits/slidingcounter"
	"go-limits/window"
	"time"
)

type Limiter interface {
	HandleRequest(ip string) bool
}

func getLimiter(limiterType string) Limiter {
	switch limiterType {
	case "bucket":
		return bucket.NewTable()
	case "fixedWindow":
		return window.NewFixedWindowLimiter(10, 1000)
	case "slidingWindow":
		return sliding.NewWindow(1000, 10)
	case "slidingCounter":
		return slidingcounter.NewSlidingWindowCounter(4 * time.Second)
	default:
		return nil
	}
}