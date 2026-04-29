package ratelimiter

import (
	"sync"
	"time"
)

type userQuota struct {
	count int
	expiresAt time.Time
}

var (
	mu sync.Mutex
	quotas = make(map[string]*userQuota)
)

const (
	limit = 20
	windowSize = time.Hour
)

func Allow(fingerprint string) bool {
	mu.Lock()
	defer mu.Unlock()

	now := time.Now()
	q, exists := quotas[fingerprint]

	if !exists || now.After(q.expiresAt) {
		quotas[fingerprint] = &userQuota{
			count: 1,
			expiresAt: now.Add(windowSize),
		}

		return true
	}

	if q.count >= limit {
		return false
	}

	q.count++
	return true
}

