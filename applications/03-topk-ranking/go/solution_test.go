package solution

import (
	"reflect"
	"testing"
)

func TestTopK(t *testing.T) {
	got := TopK([]int{3, 1, 5, 12, 2, 11}, 3)
	want := []int{12, 11, 5}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("TopK = %v, want %v", got, want)
	}

	got = TopK([]int{1, 1, 2, 2, 3}, 3)
	want = []int{3, 2, 2}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("含重复元素的 TopK = %v, want %v", got, want)
	}

	if got := TopK(nil, 3); len(got) != 0 {
		t.Fatalf("空数组应返回空结果, got %v", got)
	}
	if got := TopK([]int{1, 2}, 0); len(got) != 0 {
		t.Fatalf("k=0 应返回空结果, got %v", got)
	}
}

func TestTopKFrequent(t *testing.T) {
	got := TopKFrequent([]string{"i", "love", "leetcode", "i", "love", "coding"}, 2)
	want := []string{"i", "love"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("TopKFrequent = %v, want %v", got, want)
	}
}

func TestTopKFrequentTieBreak(t *testing.T) {
	got := TopKFrequent([]string{"b", "a", "b", "a", "c"}, 2)
	want := []string{"a", "b"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("同频次排行 = %v, want %v", got, want)
	}
}

func TestTopKClampsK(t *testing.T) {
	got := TopK([]int{4, 2, 9}, 10)
	want := []int{9, 4, 2}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("k 超过长度时 = %v, want %v", got, want)
	}
}
