package solution

import (
	"fmt"
	"testing"
)

func TestBloomFilterNoFalseNegative(t *testing.T) {
	filter := NewBloomFilter(100, 0.01)
	items := make([]string, 0, 100)
	for i := 0; i < 100; i++ {
		item := fmt.Sprintf("item-%d", i)
		items = append(items, item)
		filter.Add([]byte(item))
	}

	for _, item := range items {
		if !filter.MightContain([]byte(item)) {
			t.Fatalf("已添加元素 %q 不应被判定为不存在", item)
		}
	}
	if filter.Count() != len(items) {
		t.Fatalf("Count = %d, want %d", filter.Count(), len(items))
	}
}

func TestBloomFilterFalsePositiveRate(t *testing.T) {
	filter := NewBloomFilter(100, 0.01)
	for i := 0; i < 100; i++ {
		filter.Add([]byte(fmt.Sprintf("item-%d", i)))
	}

	falsePositives := 0
	const queries = 1000
	for i := 0; i < queries; i++ {
		if filter.MightContain([]byte(fmt.Sprintf("absent-%d", i))) {
			falsePositives++
		}
	}
	// 理论误判率约 1%，放宽到 10% 以避免测试不稳定
	if falsePositives > queries/10 {
		t.Fatalf("误判数 = %d，超过 %d", falsePositives, queries/10)
	}
}

func TestCuckooFilter(t *testing.T) {
	filter := NewCuckooFilter(256)
	items := make([]string, 0, 100)
	for i := 0; i < 100; i++ {
		item := fmt.Sprintf("item-%d", i)
		items = append(items, item)
		if !filter.Add([]byte(item)) {
			t.Fatalf("插入 %q 失败", item)
		}
	}

	for _, item := range items {
		if !filter.MightContain([]byte(item)) {
			t.Fatalf("已插入元素 %q 不应被判定为不存在", item)
		}
	}
	if filter.Count() != len(items) {
		t.Fatalf("Count = %d, want %d", filter.Count(), len(items))
	}
}

func TestCuckooFilterDelete(t *testing.T) {
	filter := NewCuckooFilter(64)
	filter.Add([]byte("alpha"))
	filter.Add([]byte("beta"))

	if !filter.Delete([]byte("alpha")) {
		t.Fatal("删除已存在元素应返回 true")
	}
	if filter.MightContain([]byte("alpha")) {
		t.Fatal("删除后不应再命中")
	}
	if !filter.MightContain([]byte("beta")) {
		t.Fatal("其他元素不应受影响")
	}
	if filter.Delete([]byte("alpha")) {
		t.Fatal("重复删除应返回 false")
	}
	if filter.Count() != 1 {
		t.Fatalf("Count = %d, want 1", filter.Count())
	}
}
