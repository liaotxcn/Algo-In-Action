package solution

import (
	"testing"
	"time"
)

// fakeClock 是可手动推进的测试时钟。
type fakeClock struct {
	current time.Time
}

// Now 返回假时钟当前时间。
func (c *fakeClock) Now() time.Time {
	return c.current
}

// Advance 推进假时钟。
func (c *fakeClock) Advance(duration time.Duration) {
	c.current = c.current.Add(duration)
}

func TestTokenBucket(t *testing.T) {
	clock := &fakeClock{current: time.Unix(0, 0)}
	bucket := NewTokenBucket(2, 1, clock.Now)

	if !bucket.Allow() || !bucket.Allow() {
		t.Fatal("初始容量内的请求应被允许")
	}
	if bucket.Allow() {
		t.Fatal("令牌耗尽后请求应被拒绝")
	}

	clock.Advance(500 * time.Millisecond)
	if bucket.Allow() {
		t.Fatal("补充不足一个令牌时应拒绝")
	}

	clock.Advance(500 * time.Millisecond)
	if !bucket.Allow() {
		t.Fatal("补满一个令牌后应允许")
	}
}

func TestTokenBucketRefillCapped(t *testing.T) {
	clock := &fakeClock{current: time.Unix(0, 0)}
	bucket := NewTokenBucket(2, 10, clock.Now)
	bucket.Allow()
	bucket.Allow()

	clock.Advance(time.Hour)
	if tokens := bucket.Tokens(); tokens != 2 {
		t.Fatalf("补充后令牌数 = %v, want 2", tokens)
	}
}

func TestSlidingWindow(t *testing.T) {
	clock := &fakeClock{current: time.Unix(0, 0)}
	limiter := NewSlidingWindow(2, time.Second, clock.Now)

	if !limiter.Allow() || !limiter.Allow() {
		t.Fatal("窗口内前两次请求应被允许")
	}
	if limiter.Allow() {
		t.Fatal("超过窗口上限应被拒绝")
	}

	clock.Advance(400 * time.Millisecond)
	if limiter.Allow() {
		t.Fatal("窗口尚未滑出旧请求时应拒绝")
	}

	clock.Advance(700 * time.Millisecond)
	if !limiter.Allow() {
		t.Fatal("旧请求滑出窗口后应允许")
	}
	if got := limiter.Count(); got != 1 {
		t.Fatalf("窗口内请求数 = %d, want 1", got)
	}
}

func TestSlidingWindowZeroLimit(t *testing.T) {
	clock := &fakeClock{current: time.Unix(0, 0)}
	limiter := NewSlidingWindow(0, time.Second, clock.Now)
	if limiter.Allow() {
		t.Fatal("上限为 0 时应拒绝所有请求")
	}
}
