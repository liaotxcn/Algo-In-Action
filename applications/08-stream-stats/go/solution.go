// Package solution 提供流式统计组件的 Go 实现。
package solution

import (
	"encoding/binary"
	"hash/fnv"
	"math/rand"
	"sync"
)

const (
	// defaultSketchWidth 是 Count-Min Sketch 每行的默认计数器数量。
	defaultSketchWidth = 256
	// defaultSketchDepth 是 Count-Min Sketch 的默认哈希行数。
	defaultSketchDepth = 4
	// hashSeedStride 用于为每一行生成不同的哈希种子。
	hashSeedStride = 0x9E3779B97F4A7C15
)

// hashWithSeed 计算带种子的 FNV-1a 64 位哈希。
func hashWithSeed(data []byte, seed uint64) uint64 {
	hasher := fnv.New64a()
	var buffer [8]byte
	binary.LittleEndian.PutUint64(buffer[:], seed)
	hasher.Write(buffer[:])
	hasher.Write(data)
	return hasher.Sum64()
}

// CountMinSketch 用固定内存估计元素出现次数，只会高估不会低估。
type CountMinSketch struct {
	mu     sync.Mutex
	width  int
	depth  int
	counts [][]int
}

// NewCountMinSketch 创建计数器矩阵；参数非正时使用默认值。
func NewCountMinSketch(width, depth int) *CountMinSketch {
	if width <= 0 {
		width = defaultSketchWidth
	}
	if depth <= 0 {
		depth = defaultSketchDepth
	}
	counts := make([][]int, depth)
	for row := range counts {
		counts[row] = make([]int, width)
	}
	return &CountMinSketch{width: width, depth: depth, counts: counts}
}

// Add 增加元素的计数。
func (s *CountMinSketch) Add(data []byte, delta int) {
	if delta <= 0 {
		return
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	for row := 0; row < s.depth; row++ {
		column := int(hashWithSeed(data, uint64(row)*hashSeedStride+1) % uint64(s.width))
		s.counts[row][column] += delta
	}
}

// Estimate 返回元素出现次数的估计值。
func (s *CountMinSketch) Estimate(data []byte) int {
	s.mu.Lock()
	defer s.mu.Unlock()

	estimate := 0
	for row := 0; row < s.depth; row++ {
		column := int(hashWithSeed(data, uint64(row)*hashSeedStride+1) % uint64(s.width))
		if row == 0 || s.counts[row][column] < estimate {
			estimate = s.counts[row][column]
		}
	}
	return estimate
}

// MovingAverage 维护最近 capacity 个数值的滑动平均。
type MovingAverage struct {
	mu       sync.Mutex
	capacity int
	values   []float64
	index    int
	count    int
	sum      float64
}

// NewMovingAverage 创建窗口大小为 capacity 的滑动平均；参数非正时按 1 处理。
func NewMovingAverage(capacity int) *MovingAverage {
	if capacity < 1 {
		capacity = 1
	}
	return &MovingAverage{capacity: capacity, values: make([]float64, capacity)}
}

// Add 加入新数值，窗口已满时替换最旧的数值。
func (m *MovingAverage) Add(value float64) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if m.count == m.capacity {
		m.sum -= m.values[m.index]
	} else {
		m.count++
	}
	m.values[m.index] = value
	m.sum += value
	m.index = (m.index + 1) % m.capacity
}

// Average 返回当前窗口内的平均值；窗口为空时返回 0。
func (m *MovingAverage) Average() float64 {
	m.mu.Lock()
	defer m.mu.Unlock()

	if m.count == 0 {
		return 0
	}
	return m.sum / float64(m.count)
}

// Len 返回窗口内有效数值个数。
func (m *MovingAverage) Len() int {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.count
}

// ReservoirSampler 使用蓄水池抽样从数据流中等概率抽取固定数量样本。
type ReservoirSampler struct {
	mu       sync.Mutex
	capacity int
	random   *rand.Rand
	seen     int
	items    []int
}

// NewReservoirSampler 创建抽样器；相同 seed 与相同输入会得到相同结果。
func NewReservoirSampler(capacity int, seed int64) *ReservoirSampler {
	if capacity < 0 {
		capacity = 0
	}
	return &ReservoirSampler{
		capacity: capacity,
		random:   rand.New(rand.NewSource(seed)),
		items:    make([]int, 0, capacity),
	}
}

// Add 向数据流中加入一个元素。
func (s *ReservoirSampler) Add(value int) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.seen++
	if len(s.items) < s.capacity {
		s.items = append(s.items, value)
		return
	}
	if s.capacity == 0 {
		return
	}
	// 第 seen 个元素以 capacity/seen 的概率进入蓄水池
	if slot := s.random.Intn(s.seen); slot < s.capacity {
		s.items[slot] = value
	}
}

// Sample 返回当前样本的副本。
func (s *ReservoirSampler) Sample() []int {
	s.mu.Lock()
	defer s.mu.Unlock()

	result := make([]int, len(s.items))
	copy(result, s.items)
	return result
}

// Seen 返回已观察到的元素总数。
func (s *ReservoirSampler) Seen() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.seen
}
