package window

import "sync"

type FixedWindowLimiter struct {
	maxRequests int
	windowSize  int64
	store       map[string]*Window
	mu          sync.Mutex
}

func NewFixedWindowLimiter(maxRequests int, windowSize int64) *FixedWindowLimiter {
	return &FixedWindowLimiter{
		maxRequests: maxRequests,
		windowSize:  windowSize,
		store:       make(map[string]*Window),
	}
}

func (l *FixedWindowLimiter) HandleRequest(ip string) bool {
	l.mu.Lock()
	defer l.mu.Unlock()
	
	if window, ok := l.store[ip]; ok {
		if window.count < l.maxRequests {
			window.count++
			return true
		}
		return false
	}
	l.store[ip] = NewWindow(l.windowSize, 1)
	return true
}
