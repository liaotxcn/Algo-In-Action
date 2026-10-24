// Package solution 提供分布式 ID 生成与任务调度的 Go 实现。
package solution

import (
	"container/heap"
	"errors"
	"sync"
	"time"
)

const (
	// nodeBits 是机器号占用的位数。
	nodeBits = 10
	// sequenceBits 是同一毫秒内序列号占用的位数。
	sequenceBits = 12
	// maxNodeID 是机器号最大值。
	maxNodeID = (1 << nodeBits) - 1
	// maxSequence 是序列号最大值。
	maxSequence = (1 << sequenceBits) - 1
	// timestampShift 是时间戳左移位数。
	timestampShift = nodeBits + sequenceBits
	// nodeShift 是机器号左移位数。
	nodeShift = sequenceBits
)

var (
	// ErrInvalidNodeID 表示机器号超出 0-1023 范围。
	ErrInvalidNodeID = errors.New("节点 ID 超出范围")
	// ErrClockBackwards 表示检测到时钟回拨。
	ErrClockBackwards = errors.New("检测到时钟回拨")
	// ErrSequenceExhausted 表示当前毫秒内序列号已用尽。
	ErrSequenceExhausted = errors.New("当前毫秒内序列号已用尽")
	// ErrNoWorkers 表示调度器没有可用的执行节点。
	ErrNoWorkers = errors.New("没有可用的执行节点")
)

// Snowflake 生成 64 位趋势递增 ID：1 位符号 + 41 位时间戳 + 10 位机器号 + 12 位序列号。
type Snowflake struct {
	mu          sync.Mutex
	epochMillis int64
	nodeID      int64
	lastMillis  int64
	sequence    int64
	now         func() time.Time
}

// NewSnowflake 创建 ID 生成器；now 为 nil 时使用 time.Now。
func NewSnowflake(nodeID int64, now func() time.Time) (*Snowflake, error) {
	if nodeID < 0 || nodeID > maxNodeID {
		return nil, ErrInvalidNodeID
	}
	if now == nil {
		now = time.Now
	}
	// 自定义纪元固定为 2020-01-01 UTC，保证 41 位时间戳可用约 69 年
	epoch := time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC)
	return &Snowflake{
		epochMillis: epoch.UnixMilli(),
		nodeID:      nodeID,
		now:         now,
	}, nil
}

// NextID 生成下一个 ID；时钟回拨或序列号耗尽时返回错误。
func (s *Snowflake) NextID() (int64, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	current := s.now().UnixMilli()
	if current < s.lastMillis {
		return 0, ErrClockBackwards
	}
	if current == s.lastMillis {
		if s.sequence >= maxSequence {
			return 0, ErrSequenceExhausted
		}
		s.sequence++
	} else {
		s.sequence = 0
	}
	s.lastMillis = current

	id := ((current - s.epochMillis) << timestampShift) |
		(s.nodeID << nodeShift) |
		s.sequence
	return id, nil
}

// Parse 拆解 ID，返回生成时间、机器号与序列号。
func (s *Snowflake) Parse(id int64) (time.Time, int64, int64) {
	millis := (id >> timestampShift) + s.epochMillis
	nodeID := (id >> nodeShift) & maxNodeID
	sequence := id & maxSequence
	return time.UnixMilli(millis), nodeID, sequence
}

// Scheduler 按轮询方式把任务分配到执行节点。
type Scheduler struct {
	mu      sync.Mutex
	workers []string
	next    int
}

// NewScheduler 创建轮询调度器，workers 为空时所有分配返回错误。
func NewScheduler(workers []string) *Scheduler {
	copied := make([]string, len(workers))
	copy(copied, workers)
	return &Scheduler{workers: copied}
}

// Assign 返回下一个执行节点。
func (s *Scheduler) Assign() (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if len(s.workers) == 0 {
		return "", ErrNoWorkers
	}
	worker := s.workers[s.next]
	s.next = (s.next + 1) % len(s.workers)
	return worker, nil
}

// Task 是待调度的任务。
type Task struct {
	Name     string
	Priority int
	seq      int64
}

// taskQueue 是高优先级优先、同优先级先进先出的堆。
type taskQueue []Task

func (q taskQueue) Len() int { return len(q) }
func (q taskQueue) Less(i, j int) bool {
	if q[i].Priority != q[j].Priority {
		return q[i].Priority > q[j].Priority
	}
	return q[i].seq < q[j].seq
}
func (q taskQueue) Swap(i, j int) { q[i], q[j] = q[j], q[i] }
func (q *taskQueue) Push(x any)   { *q = append(*q, x.(Task)) }
func (q *taskQueue) Pop() any {
	old := *q
	last := len(old) - 1
	value := old[last]
	*q = old[:last]
	return value
}

// PriorityScheduler 按优先级调度任务，相同优先级保持提交顺序。
type PriorityScheduler struct {
	mu      sync.Mutex
	queue   taskQueue
	nextSeq int64
}

// NewPriorityScheduler 创建优先级调度器。
func NewPriorityScheduler() *PriorityScheduler {
	return &PriorityScheduler{}
}

// Submit 提交任务，priority 越大越先执行。
func (s *PriorityScheduler) Submit(name string, priority int) {
	s.mu.Lock()
	defer s.mu.Unlock()

	task := Task{Name: name, Priority: priority, seq: s.nextSeq}
	s.nextSeq++
	heap.Push(&s.queue, task)
}

// Pop 取出下一个任务；队列为空时返回 false。
func (s *PriorityScheduler) Pop() (Task, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.queue.Len() == 0 {
		return Task{}, false
	}
	return heap.Pop(&s.queue).(Task), true
}
