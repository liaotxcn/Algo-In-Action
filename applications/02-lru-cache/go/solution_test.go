package solution

import "testing"

func TestLRUGetPut(t *testing.T) {
	cache := NewLRUCache[int, string](2)
	cache.Put(1, "one")
	cache.Put(2, "two")

	if value, ok := cache.Get(1); !ok || value != "one" {
		t.Fatalf("Get(1) = %q, %v, want one, true", value, ok)
	}

	cache.Put(3, "three") // 淘汰最久未使用的 2
	if _, ok := cache.Get(2); ok {
		t.Fatal("键 2 应已被淘汰")
	}
	if value, ok := cache.Get(3); !ok || value != "three" {
		t.Fatalf("Get(3) = %q, %v, want three, true", value, ok)
	}

	hits, misses := cache.Stats()
	if hits != 2 {
		t.Fatalf("命中次数 = %d, want 2", hits)
	}
	if misses != 1 {
		t.Fatalf("未命中次数 = %d, want 1", misses)
	}
}

func TestLRUUpdateRefreshesOrder(t *testing.T) {
	cache := NewLRUCache[int, int](2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	cache.Put(1, 10) // 更新后 1 变为最近使用
	cache.Put(3, 3)  // 应淘汰 2

	if value, ok := cache.Get(1); !ok || value != 10 {
		t.Fatalf("Get(1) = %d, %v, want 10, true", value, ok)
	}
	if _, ok := cache.Get(2); ok {
		t.Fatal("键 2 应已被淘汰")
	}
	if cache.Len() != 2 {
		t.Fatalf("Len = %d, want 2", cache.Len())
	}
}

func TestLRUZeroCapacity(t *testing.T) {
	cache := NewLRUCache[int, int](0)
	cache.Put(1, 1)
	if _, ok := cache.Get(1); ok {
		t.Fatal("零容量缓存不应保留数据")
	}
}
