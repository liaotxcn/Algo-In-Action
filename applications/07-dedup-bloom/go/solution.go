// Package solution 提供布隆过滤器与布谷鸟过滤器的 Go 实现。
package solution

import (
	"encoding/binary"
	"hash/fnv"
	"math"
	"sync"
)

const (
	// defaultFalsePositiveRate 是非法误判率参数时使用的默认值。
	defaultFalsePositiveRate = 0.01
	// defaultBucketSize 是布谷鸟过滤器每个桶的槽位数。
	defaultBucketSize = 4
	// defaultMaxKicks 是布谷鸟过滤器插入时的最大踢出次数。
	defaultMaxKicks = 500
	// fingerprintMask 用于截取 16 位指纹。
	fingerprintMask = 0xFFFF
	// bitsPerWord 是位图底层 uint64 的字长。
	bitsPerWord = 64
	// wordMask 是位图字内下标的掩码。
	wordMask = bitsPerWord - 1
	// fingerprintSeed 是布谷鸟指纹的哈希种子。
	fingerprintSeed = 0x5151
	// primarySeed 是布谷鸟主桶的哈希种子。
	primarySeed = 0xA5A5
	// alternateSeed 是布谷鸟备用桶的哈希种子。
	alternateSeed = 0xBEEF
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

// BloomFilter 是线程安全的布隆过滤器。
type BloomFilter struct {
	mu        sync.RWMutex
	bits      []uint64
	hashCount int
	itemCount int
	bitCount  int
}

// NewBloomFilter 根据预期元素数与误判率计算位数组长度与哈希函数个数。
func NewBloomFilter(expectedItems int, falsePositiveRate float64) *BloomFilter {
	if expectedItems < 1 {
		expectedItems = 1
	}
	if falsePositiveRate <= 0 || falsePositiveRate >= 1 {
		falsePositiveRate = defaultFalsePositiveRate
	}

	bitCount := int(math.Ceil(
		-float64(expectedItems) * math.Log(falsePositiveRate) / (math.Ln2 * math.Ln2),
	))
	hashCount := int(math.Round(float64(bitCount) / float64(expectedItems) * math.Ln2))
	if hashCount < 1 {
		hashCount = 1
	}
	return &BloomFilter{
		bits:      make([]uint64, (bitCount+bitsPerWord-1)/bitsPerWord),
		hashCount: hashCount,
		bitCount:  bitCount,
	}
}

// positions 返回元素对应的 k 个位下标（双重哈希）。
func (b *BloomFilter) positions(data []byte) []uint64 {
	first := hashWithSeed(data, 0)
	second := hashWithSeed(data, 1) | 1
	width := uint64(b.bitCount)
	result := make([]uint64, b.hashCount)
	for i := 0; i < b.hashCount; i++ {
		result[i] = (first + uint64(i)*second) % width
	}
	return result
}

// Add 把元素加入过滤器。
func (b *BloomFilter) Add(data []byte) {
	b.mu.Lock()
	defer b.mu.Unlock()

	for _, position := range b.positions(data) {
		b.bits[position/bitsPerWord] |= 1 << (position & wordMask)
	}
	b.itemCount++
}

// MightContain 判断元素是否可能存在；false 表示一定不存在。
func (b *BloomFilter) MightContain(data []byte) bool {
	b.mu.RLock()
	defer b.mu.RUnlock()

	for _, position := range b.positions(data) {
		if b.bits[position/bitsPerWord]&(1<<(position&wordMask)) == 0 {
			return false
		}
	}
	return true
}

// Count 返回已添加的元素个数。
func (b *BloomFilter) Count() int {
	b.mu.RLock()
	defer b.mu.RUnlock()
	return b.itemCount
}

// BitCount 返回位数组长度。
func (b *BloomFilter) BitCount() int {
	return b.bitCount
}

// HashCount 返回哈希函数个数。
func (b *BloomFilter) HashCount() int {
	return b.hashCount
}

// CuckooFilter 是支持删除的布谷鸟过滤器，指纹存放在固定槽位数的桶中。
type CuckooFilter struct {
	mu         sync.Mutex
	buckets    [][]uint16
	mask       int
	bucketSize int
	maxKicks   int
	count      int
}

// NewCuckooFilter 创建容量至少为 capacity 的布谷鸟过滤器。
func NewCuckooFilter(capacity int) *CuckooFilter {
	if capacity < 1 {
		capacity = 1
	}
	bucketCount := 1
	for bucketCount*defaultBucketSize < capacity {
		bucketCount <<= 1
	}
	if bucketCount < 2 {
		bucketCount = 2
	}

	buckets := make([][]uint16, bucketCount)
	for i := range buckets {
		buckets[i] = make([]uint16, defaultBucketSize)
	}
	return &CuckooFilter{
		buckets:    buckets,
		mask:       bucketCount - 1,
		bucketSize: defaultBucketSize,
		maxKicks:   defaultMaxKicks,
	}
}

// fingerprint 返回元素对应的非零 16 位指纹。
func (c *CuckooFilter) fingerprint(data []byte) uint16 {
	fp := uint16(hashWithSeed(data, fingerprintSeed) & fingerprintMask)
	if fp == 0 {
		fp = 1
	}
	return fp
}

// primaryIndex 返回元素的主桶下标。
func (c *CuckooFilter) primaryIndex(data []byte) int {
	return int(hashWithSeed(data, primarySeed) & uint64(c.mask))
}

// altIndex 返回备用桶下标；对同一指纹再次异或可回到原下标。
func (c *CuckooFilter) altIndex(index int, fingerprint uint16) int {
	var buffer [2]byte
	binary.LittleEndian.PutUint16(buffer[:], fingerprint)
	return (index ^ int(hashWithSeed(buffer[:], alternateSeed))) & c.mask
}

// insertInto 尝试把指纹放入指定桶的空槽。
func (c *CuckooFilter) insertInto(index int, fingerprint uint16) bool {
	for slot := range c.buckets[index] {
		if c.buckets[index][slot] == 0 {
			c.buckets[index][slot] = fingerprint
			return true
		}
	}
	return false
}

// Add 插入元素；过滤器已满且踢出失败时返回 false。
func (c *CuckooFilter) Add(data []byte) bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	fingerprint := c.fingerprint(data)
	first := c.primaryIndex(data)
	second := c.altIndex(first, fingerprint)
	if c.insertInto(first, fingerprint) || c.insertInto(second, fingerprint) {
		c.count++
		return true
	}

	// 两个桶都满时，轮流踢出槽位中的指纹并寻找空位
	index := first
	current := fingerprint
	for kick := 0; kick < c.maxKicks; kick++ {
		slot := kick % c.bucketSize
		current, c.buckets[index][slot] = c.buckets[index][slot], current
		index = c.altIndex(index, current)
		if c.insertInto(index, current) {
			c.count++
			return true
		}
	}
	return false
}

// MightContain 判断元素是否存在；false 表示一定不存在。
func (c *CuckooFilter) MightContain(data []byte) bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	fingerprint := c.fingerprint(data)
	first := c.primaryIndex(data)
	second := c.altIndex(first, fingerprint)
	return c.containsIn(first, fingerprint) || c.containsIn(second, fingerprint)
}

// containsIn 判断指定桶是否包含指纹。
func (c *CuckooFilter) containsIn(index int, fingerprint uint16) bool {
	for _, stored := range c.buckets[index] {
		if stored == fingerprint {
			return true
		}
	}
	return false
}

// Delete 删除元素；不存在时返回 false。
func (c *CuckooFilter) Delete(data []byte) bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	fingerprint := c.fingerprint(data)
	first := c.primaryIndex(data)
	second := c.altIndex(first, fingerprint)
	for _, index := range []int{first, second} {
		for slot, stored := range c.buckets[index] {
			if stored != fingerprint {
				continue
			}
			c.buckets[index][slot] = 0
			c.count--
			return true
		}
	}
	return false
}

// Count 返回已插入的元素个数。
func (c *CuckooFilter) Count() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}
