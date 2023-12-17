package slidingcounter

import (
	"sync"
	"time"
)

type SlidingWindowCounter struct {
	mu              sync.Mutex
	currentWindow   int
	previousWindow  int
	lastWindowStart int64
	windowSize      time.Duration
}

func NewSlidingWindowCounter(windowSize time.Duration) *SlidingWindowCounter {
	return &SlidingWindowCounter{
		lastWindowStart: time.Now().UnixNano(),
		windowSize:      windowSize,
	}
}

func (swc *SlidingWindowCounter) HandleRequest(clientIP string) bool {
	swc.mu.Lock()
	defer swc.mu.Unlock()

	now := time.Now()
	elapsed := now.Sub(time.Unix(0, swc.lastWindowStart))

	if elapsed >= swc.windowSize {
		swc.previousWindow = swc.currentWindow
		swc.currentWindow = 1
		swc.lastWindowStart = now.UnixNano()
	} else {
		swc.currentWindow++
	}

	threshold := 5
	return swc.currentWindow <= threshold
}
