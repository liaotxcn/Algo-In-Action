// Package solution 提供线程安全的 LRU 缓存 Go 实现。
package solution

import (
	"container/list"
	"sync"
)

// cacheEntry 保存缓存的键值对，键用于淘汰时反查 map。
type cacheEntry[K comparable, V any] struct {
	key   K
	value V
}

// LRUCache 是基于哈希表 + 双向链表的最近最少使用缓存。
// 访问和写入均为 O(1)，所有操作由互斥锁保护。
type LRUCache[K comparable, V any] struct {
	mu       sync.Mutex
	capacity int
	order    *list.List
	items    map[K]*list.Element
	hits     int64
	misses   int64
}

// NewLRUCache 创建指定容量的 LRU 缓存；容量为 0 时不缓存任何数据。
func NewLRUCache[K comparable, V any](capacity int) *LRUCache[K, V] {
	if capacity < 0 {
		capacity = 0
	}
	return &LRUCache[K, V]{
		capacity: capacity,
		order:    list.New(),
		items:    make(map[K]*list.Element),
	}
}

// Get 读取键值；命中时把节点移到链表头部。
func (c *LRUCache[K, V]) Get(key K) (V, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	element, ok := c.items[key]
	if !ok {
		c.misses++
		var zero V
		return zero, false
	}
	c.hits++
	c.order.MoveToFront(element)
	return element.Value.(*cacheEntry[K, V]).value, true
}

// Put 写入键值；键已存在时更新并刷新热度，容量不足时淘汰链表尾部。
func (c *LRUCache[K, V]) Put(key K, value V) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.capacity == 0 {
		return
	}
	if element, ok := c.items[key]; ok {
		entry := element.Value.(*cacheEntry[K, V])
		entry.value = value
		c.order.MoveToFront(element)
		return
	}
	if c.order.Len() >= c.capacity {
		c.evictLocked()
	}
	entry := &cacheEntry[K, V]{key: key, value: value}
	c.items[key] = c.order.PushFront(entry)
}

// evictLocked 淘汰最久未使用的节点，调用方必须持有锁。
func (c *LRUCache[K, V]) evictLocked() {
	oldest := c.order.Back()
	if oldest == nil {
		return
	}
	c.order.Remove(oldest)
	delete(c.items, oldest.Value.(*cacheEntry[K, V]).key)
}

// Len 返回当前缓存条目数。
func (c *LRUCache[K, V]) Len() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.order.Len()
}

// Stats 返回命中与未命中次数。
func (c *LRUCache[K, V]) Stats() (int64, int64) {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.hits, c.misses
}
