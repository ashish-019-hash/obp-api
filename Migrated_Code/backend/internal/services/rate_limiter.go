package services

import (
	"sync"
	"time"
)

type RateLimiter struct {
	windows map[string]*SlidingWindow
	mutex   sync.RWMutex
	
	requestsPerMinute int
	requestsPerHour   int
}

type SlidingWindow struct {
	requests []time.Time
	mutex    sync.Mutex
}

func NewRateLimiter() *RateLimiter {
	return &RateLimiter{
		windows:           make(map[string]*SlidingWindow),
		requestsPerMinute: 100,
		requestsPerHour:   1000,
	}
}

func (rl *RateLimiter) IsLimited(identifier string) bool {
	rl.mutex.Lock()
	defer rl.mutex.Unlock()
	
	window, exists := rl.windows[identifier]
	if !exists {
		window = &SlidingWindow{requests: make([]time.Time, 0)}
		rl.windows[identifier] = window
	}
	
	return window.checkLimit(rl.requestsPerMinute, rl.requestsPerHour)
}

func (rl *RateLimiter) SetLimits(perMinute, perHour int) {
	rl.mutex.Lock()
	defer rl.mutex.Unlock()
	
	rl.requestsPerMinute = perMinute
	rl.requestsPerHour = perHour
}

func (rl *RateLimiter) GetLimits() (int, int) {
	rl.mutex.RLock()
	defer rl.mutex.RUnlock()
	
	return rl.requestsPerMinute, rl.requestsPerHour
}

func (sw *SlidingWindow) checkLimit(perMinute, perHour int) bool {
	sw.mutex.Lock()
	defer sw.mutex.Unlock()
	
	now := time.Now()
	oneMinuteAgo := now.Add(-time.Minute)
	oneHourAgo := now.Add(-time.Hour)
	
	var validRequests []time.Time
	minuteCount := 0
	hourCount := 0
	
	for _, req := range sw.requests {
		if req.After(oneHourAgo) {
			validRequests = append(validRequests, req)
			hourCount++
			if req.After(oneMinuteAgo) {
				minuteCount++
			}
		}
	}
	
	sw.requests = validRequests
	
	if minuteCount >= perMinute || hourCount >= perHour {
		return true
	}
	
	sw.requests = append(sw.requests, now)
	return false
}

func (rl *RateLimiter) Reset(identifier string) {
	rl.mutex.Lock()
	defer rl.mutex.Unlock()
	
	delete(rl.windows, identifier)
}

func (rl *RateLimiter) GetStats(identifier string) (int, int) {
	rl.mutex.RLock()
	defer rl.mutex.RUnlock()
	
	window, exists := rl.windows[identifier]
	if !exists {
		return 0, 0
	}
	
	window.mutex.Lock()
	defer window.mutex.Unlock()
	
	now := time.Now()
	oneMinuteAgo := now.Add(-time.Minute)
	oneHourAgo := now.Add(-time.Hour)
	
	minuteCount := 0
	hourCount := 0
	
	for _, req := range window.requests {
		if req.After(oneHourAgo) {
			hourCount++
			if req.After(oneMinuteAgo) {
				minuteCount++
			}
		}
	}
	
	return minuteCount, hourCount
}
