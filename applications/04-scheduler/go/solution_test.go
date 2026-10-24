package solution

import (
	"errors"
	"testing"
	"time"
)

func TestSnowflakeIDs(t *testing.T) {
	clock := &stepClock{current: time.Now()}
	generator, err := NewSnowflake(7, clock.Now)
	if err != nil {
		t.Fatalf("NewSnowflake 返回错误: %v", err)
	}

	first, err := generator.NextID()
	if err != nil {
		t.Fatalf("NextID 返回错误: %v", err)
	}
	second, err := generator.NextID()
	if err != nil {
		t.Fatalf("NextID 返回错误: %v", err)
	}
	if second <= first {
		t.Fatalf("同一毫秒内 ID 应递增: %d, %d", first, second)
	}

	generatedAt, nodeID, sequence := generator.Parse(second)
	if nodeID != 7 {
		t.Fatalf("机器号 = %d, want 7", nodeID)
	}
	if sequence != 1 {
		t.Fatalf("序列号 = %d, want 1", sequence)
	}
	if generatedAt.UnixMilli() != clock.Now().UnixMilli() {
		t.Fatalf("解析时间 = %v, want %v", generatedAt, clock.Now())
	}

	clock.Advance(time.Millisecond)
	third, err := generator.NextID()
	if err != nil {
		t.Fatalf("NextID 返回错误: %v", err)
	}
	if _, _, sequence := generator.Parse(third); sequence != 0 {
		t.Fatalf("新毫秒序列号应重置为 0, got %d", sequence)
	}
}

func TestSnowflakeValidation(t *testing.T) {
	if _, err := NewSnowflake(-1, nil); !errors.Is(err, ErrInvalidNodeID) {
		t.Fatalf("非法机器号错误 = %v, want ErrInvalidNodeID", err)
	}

	clock := &stepClock{current: time.Now()}
	generator, _ := NewSnowflake(1, clock.Now)
	if _, err := generator.NextID(); err != nil {
		t.Fatalf("首次生成失败: %v", err)
	}
	clock.Advance(-time.Second)
	if _, err := generator.NextID(); !errors.Is(err, ErrClockBackwards) {
		t.Fatalf("时钟回拨错误 = %v, want ErrClockBackwards", err)
	}
}

func TestRoundRobinScheduler(t *testing.T) {
	scheduler := NewScheduler([]string{"a", "b", "c"})
	want := []string{"a", "b", "c", "a"}
	for index, expected := range want {
		got, err := scheduler.Assign()
		if err != nil {
			t.Fatalf("第 %d 次分配失败: %v", index, err)
		}
		if got != expected {
			t.Fatalf("第 %d 次分配 = %q, want %q", index, got, expected)
		}
	}

	empty := NewScheduler(nil)
	if _, err := empty.Assign(); !errors.Is(err, ErrNoWorkers) {
		t.Fatalf("空调度器错误 = %v, want ErrNoWorkers", err)
	}
}

func TestPriorityScheduler(t *testing.T) {
	scheduler := NewPriorityScheduler()
	scheduler.Submit("low", 1)
	scheduler.Submit("high-1", 10)
	scheduler.Submit("mid", 5)
	scheduler.Submit("high-2", 10)

	want := []string{"high-1", "high-2", "mid", "low"}
	for _, expected := range want {
		task, ok := scheduler.Pop()
		if !ok {
			t.Fatalf("应取出任务 %q", expected)
		}
		if task.Name != expected {
			t.Fatalf("取出任务 = %q, want %q", task.Name, expected)
		}
	}
	if _, ok := scheduler.Pop(); ok {
		t.Fatal("队列为空时应返回 false")
	}
}

// stepClock 是可手动推进或回拨的测试时钟。
type stepClock struct {
	current time.Time
}

// Now 返回测试时钟当前时间。
func (c *stepClock) Now() time.Time {
	return c.current
}

// Advance 推进测试时钟，负值表示回拨。
func (c *stepClock) Advance(duration time.Duration) {
	c.current = c.current.Add(duration)
}
