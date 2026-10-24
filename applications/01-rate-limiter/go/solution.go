// Package solution 提供限流器的 Go 实现。
package solution

import (
	"sync"
	"time"
)

// Clock 返回当前时间，测试时可注入假时钟。
type Clock func() time.Time

const (
	// tokensPerRequest 表示单次请求消耗的令牌数。
	tokensPerRequest = 1.0
	// defaultCapacity 表示容量参数非法时使用的默认值。
	defaultCapacity = 1
)

// TokenBucket 实现令牌桶限流：桶内令牌按固定速率补充，请求消耗令牌。
type TokenBucket struct {
	mu              sync.Mutex
	capacity        float64
	tokens          float64
	refillPerSecond float64
	lastRefill      time.Time
	now             Clock
}

// NewTokenBucket 创建令牌桶；capacity 为桶容量，refillPerSecond 为每秒补充令牌数。
func NewTokenBucket(capacity int, refillPerSecond float64, now Clock) *TokenBucket {
	if now == nil {
		now = time.Now
	}
	if capacity < defaultCapacity {
		capacity = defaultCapacity
	}
	if refillPerSecond < 0 {
		refillPerSecond = 0
	}
	current := now()
	return &TokenBucket{
		capacity:        float64(capacity),
		tokens:          float64(capacity),
		refillPerSecond: refillPerSecond,
		lastRefill:      current,
		now:             now,
	}
}

// Allow 尝试获取一个令牌，成功返回 true。
func (b *TokenBucket) Allow() bool {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.refillLocked(b.now())
	if b.tokens < tokensPerRequest {
		return false
	}
	b.tokens -= tokensPerRequest
	return true
}

// Tokens 返回当前可用令牌数（用于观测）。
func (b *TokenBucket) Tokens() float64 {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.refillLocked(b.now())
	return b.tokens
}

// refillLocked 按经过时间补充令牌，调用方必须持有锁。
func (b *TokenBucket) refillLocked(current time.Time) {
	if elapsed := current.Sub(b.lastRefill).Seconds(); elapsed > 0 {
		b.tokens = min(b.capacity, b.tokens+elapsed*b.refillPerSecond)
		b.lastRefill = current
	}
}

// SlidingWindow 实现滑动窗口限流：窗口内请求数不超过 limit。
type SlidingWindow struct {
	mu     sync.Mutex
	limit  int
	window time.Duration
	events []time.Time
	now    Clock
}

// NewSlidingWindow 创建滑动窗口限流器。
func NewSlidingWindow(limit int, window time.Duration, now Clock) *SlidingWindow {
	if now == nil {
		now = time.Now
	}
	if limit < 0 {
		limit = 0
	}
	if window < 0 {
		window = 0
	}
	return &SlidingWindow{limit: limit, window: window, now: now}
}

// Allow 记录一次请求；窗口内请求数达到上限时返回 false。
func (w *SlidingWindow) Allow() bool {
	w.mu.Lock()
	defer w.mu.Unlock()

	current := w.now()
	cutoff := current.Add(-w.window)
	kept := w.events[:0]
	for _, event := range w.events {
		if event.After(cutoff) {
			kept = append(kept, event)
		}
	}
	w.events = kept

	if len(w.events) >= w.limit {
		return false
	}
	w.events = append(w.events, current)
	return true
}

// Count 返回当前窗口内的请求数。
func (w *SlidingWindow) Count() int {
	w.mu.Lock()
	defer w.mu.Unlock()

	current := w.now()
	cutoff := current.Add(-w.window)
	count := 0
	for _, event := range w.events {
		if event.After(cutoff) {
			count++
		}
	}
	return count
}
