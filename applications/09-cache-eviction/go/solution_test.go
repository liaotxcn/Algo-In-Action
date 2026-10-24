package solution

import "testing"

func TestFIFOEviction(t *testing.T) {
	cache := NewFIFOCache[int, string](2)
	cache.Put(1, "one")
	cache.Put(2, "two")

	if _, ok := cache.Get(1); !ok {
		t.Fatal("键 1 应存在")
	}
	cache.Put(3, "three") // FIFO 只看写入顺序，淘汰 1

	if _, ok := cache.Get(1); ok {
		t.Fatal("键 1 应已被淘汰")
	}
	if value, ok := cache.Get(2); !ok || value != "two" {
		t.Fatalf("Get(2) = %q, %v, want two, true", value, ok)
	}
	if value, ok := cache.Get(3); !ok || value != "three" {
		t.Fatalf("Get(3) = %q, %v, want three, true", value, ok)
	}
	if cache.Evictions() != 1 {
		t.Fatalf("Evictions = %d, want 1", cache.Evictions())
	}
}

func TestFIFOUpdateDoesNotReorder(t *testing.T) {
	cache := NewFIFOCache[int, int](2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	cache.Put(1, 10) // 更新不改变写入顺序
	cache.Put(3, 3)  // 仍淘汰 1

	if _, ok := cache.Get(1); ok {
		t.Fatal("键 1 应已被淘汰")
	}
	if value, ok := cache.Get(2); !ok || value != 2 {
		t.Fatalf("Get(2) = %d, %v, want 2, true", value, ok)
	}
}

func TestLFUEviction(t *testing.T) {
	cache := NewLFUCache[int, string](2)
	cache.Put(1, "one")
	cache.Put(2, "two")
	cache.Get(1)          // 1 的频次变为 2
	cache.Put(3, "three") // 淘汰频次最低的 2

	if _, ok := cache.Get(2); ok {
		t.Fatal("键 2 应已被淘汰")
	}
	if value, ok := cache.Get(1); !ok || value != "one" {
		t.Fatalf("Get(1) = %q, %v, want one, true", value, ok)
	}
	if value, ok := cache.Get(3); !ok || value != "three" {
		t.Fatalf("Get(3) = %q, %v, want three, true", value, ok)
	}
	if cache.Evictions() != 1 {
		t.Fatalf("Evictions = %d, want 1", cache.Evictions())
	}
}

func TestLFUTieBreakLRU(t *testing.T) {
	cache := NewLFUCache[string, int](2)
	cache.Put("a", 1)
	cache.Put("b", 2)
	cache.Get("a")
	cache.Get("b")
	cache.Put("c", 3) // a、b 频次相同，淘汰更久未使用的 a

	if _, ok := cache.Get("a"); ok {
		t.Fatal("同频次下应淘汰最久未使用的 a")
	}
	if value, ok := cache.Get("b"); !ok || value != 2 {
		t.Fatalf("Get(b) = %d, %v, want 2, true", value, ok)
	}
}

func TestZeroCapacity(t *testing.T) {
	fifo := NewFIFOCache[int, int](0)
	fifo.Put(1, 1)
	if fifo.Len() != 0 {
		t.Fatalf("FIFO 零容量 Len = %d, want 0", fifo.Len())
	}

	lfu := NewLFUCache[int, int](0)
	lfu.Put(1, 1)
	if lfu.Len() != 0 {
		t.Fatalf("LFU 零容量 Len = %d, want 0", lfu.Len())
	}
}
