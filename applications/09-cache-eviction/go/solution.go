// Package solution 提供 FIFO 与 LFU 缓存淘汰策略的 Go 实现。
package solution

import (
	"container/list"
	"sync"
)

// fifoEntry 是 FIFO 缓存的键值对。
type fifoEntry[K comparable, V any] struct {
	key   K
	value V
}

// FIFOCache 按写入顺序淘汰数据，读取不会改变顺序。
type FIFOCache[K comparable, V any] struct {
	mu        sync.Mutex
	capacity  int
	order     *list.List
	items     map[K]*list.Element
	evictions int
}

// NewFIFOCache 创建指定容量的 FIFO 缓存。
func NewFIFOCache[K comparable, V any](capacity int) *FIFOCache[K, V] {
	if capacity < 0 {
		capacity = 0
	}
	return &FIFOCache[K, V]{
		capacity: capacity,
		order:    list.New(),
		items:    make(map[K]*list.Element),
	}
}

// Get 读取键值；不改变淘汰顺序。
func (c *FIFOCache[K, V]) Get(key K) (V, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	element, ok := c.items[key]
	if !ok {
		var zero V
		return zero, false
	}
	return element.Value.(*fifoEntry[K, V]).value, true
}

// Put 写入键值；容量不足时淘汰最早写入的键。
func (c *FIFOCache[K, V]) Put(key K, value V) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.capacity == 0 {
		return
	}
	if element, ok := c.items[key]; ok {
		element.Value.(*fifoEntry[K, V]).value = value
		return
	}
	if c.order.Len() >= c.capacity {
		oldest := c.order.Front()
		if oldest != nil {
			c.order.Remove(oldest)
			delete(c.items, oldest.Value.(*fifoEntry[K, V]).key)
			c.evictions++
		}
	}
	entry := &fifoEntry[K, V]{key: key, value: value}
	c.items[key] = c.order.PushBack(entry)
}

// Len 返回当前缓存条目数。
func (c *FIFOCache[K, V]) Len() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.order.Len()
}

// Evictions 返回累计淘汰次数。
func (c *FIFOCache[K, V]) Evictions() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.evictions
}

// lfuEntry 是 LFU 缓存的键值对及访问频次。
type lfuEntry[K comparable, V any] struct {
	key     K
	value   V
	freq    int
	element *list.Element
}

// LFUCache 按访问频次淘汰数据，频次相同时淘汰最久未使用的数据。
type LFUCache[K comparable, V any] struct {
	mu        sync.Mutex
	capacity  int
	items     map[K]*lfuEntry[K, V]
	buckets   map[int]*list.List
	minFreq   int
	evictions int
}

// NewLFUCache 创建指定容量的 LFU 缓存。
func NewLFUCache[K comparable, V any](capacity int) *LFUCache[K, V] {
	if capacity < 0 {
		capacity = 0
	}
	return &LFUCache[K, V]{
		capacity: capacity,
		items:    make(map[K]*lfuEntry[K, V]),
		buckets:  make(map[int]*list.List),
	}
}

// Get 读取键值；命中时提升访问频次。
func (c *LFUCache[K, V]) Get(key K) (V, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, ok := c.items[key]
	if !ok {
		var zero V
		return zero, false
	}
	c.touchLocked(entry)
	return entry.value, true
}

// Put 写入键值；键已存在时更新并提升频次，容量不足时淘汰频次最低的键。
func (c *LFUCache[K, V]) Put(key K, value V) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.capacity == 0 {
		return
	}
	if entry, ok := c.items[key]; ok {
		entry.value = value
		c.touchLocked(entry)
		return
	}
	if len(c.items) >= c.capacity {
		c.evictLocked()
	}

	entry := &lfuEntry[K, V]{key: key, value: value, freq: 1}
	c.items[key] = entry
	entry.element = c.bucketLocked(1).PushFront(entry)
	c.minFreq = 1
}

// touchLocked 把条目移动到频次 +1 的链表头部。
func (c *LFUCache[K, V]) touchLocked(entry *lfuEntry[K, V]) {
	current := c.buckets[entry.freq]
	current.Remove(entry.element)
	if current.Len() == 0 {
		delete(c.buckets, entry.freq)
		if c.minFreq == entry.freq {
			c.minFreq++
		}
	}

	entry.freq++
	entry.element = c.bucketLocked(entry.freq).PushFront(entry)
}

// bucketLocked 返回指定频次的链表，不存在时创建。
func (c *LFUCache[K, V]) bucketLocked(freq int) *list.List {
	bucket, ok := c.buckets[freq]
	if !ok {
		bucket = list.New()
		c.buckets[freq] = bucket
	}
	return bucket
}

// evictLocked 从最低频次链表的尾部淘汰一个条目。
func (c *LFUCache[K, V]) evictLocked() {
	bucket := c.buckets[c.minFreq]
	if bucket == nil {
		return
	}
	victim := bucket.Back()
	if victim == nil {
		return
	}
	bucket.Remove(victim)
	delete(c.items, victim.Value.(*lfuEntry[K, V]).key)
	if bucket.Len() == 0 {
		delete(c.buckets, c.minFreq)
	}
	c.evictions++
}

// Len 返回当前缓存条目数。
func (c *LFUCache[K, V]) Len() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return len(c.items)
}

// Evictions 返回累计淘汰次数。
func (c *LFUCache[K, V]) Evictions() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.evictions
}
