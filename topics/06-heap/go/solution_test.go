package solution

import (
	"reflect"
	"testing"
)

func TestMaxHeap(t *testing.T) {
	heap := NewMaxHeap()
	if _, ok := heap.Pop(); ok {
		t.Fatal("空堆 Pop 应返回 false")
	}

	for _, value := range []int{5, 2, 8, 1, 9, 3} {
		heap.Push(value)
	}
	if got, ok := heap.Peek(); !ok || got != 9 {
		t.Fatalf("Peek = %d, %v; want 9, true", got, ok)
	}

	got := make([]int, 0, 6)
	for {
		value, ok := heap.Pop()
		if !ok {
			break
		}
		got = append(got, value)
	}
	want := []int{9, 8, 5, 3, 2, 1}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Pop 顺序 = %v, want %v", got, want)
	}
}

func TestHeapSort(t *testing.T) {
	cases := []struct {
		input []int
		want  []int
	}{
		{[]int{}, []int{}},
		{[]int{7}, []int{7}},
		{[]int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{[]int{4, 2, 2, 8, 3, 3, 1}, []int{1, 2, 2, 3, 3, 4, 8}},
	}
	for _, tc := range cases {
		nums := make([]int, len(tc.input))
		copy(nums, tc.input)
		HeapSort(nums)
		if !reflect.DeepEqual(nums, tc.want) {
			t.Fatalf("HeapSort(%v) = %v, want %v", tc.input, nums, tc.want)
		}
	}
}

func TestTopK(t *testing.T) {
	cases := []struct {
		nums []int
		k    int
		want []int
	}{
		{[]int{3, 2, 1, 5, 6, 4}, 2, []int{6, 5}},
		{[]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4, []int{6, 5, 5, 4}},
		{[]int{1, 2}, 5, []int{2, 1}},
		{[]int{1, 2}, 0, []int{}},
		{[]int{}, 2, []int{}},
	}
	for _, tc := range cases {
		if got := TopK(tc.nums, tc.k); !reflect.DeepEqual(got, tc.want) {
			t.Fatalf("TopK(%v, %d) = %v, want %v", tc.nums, tc.k, got, tc.want)
		}
	}
}
