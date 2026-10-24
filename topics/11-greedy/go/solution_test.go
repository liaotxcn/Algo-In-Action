package solution

import (
	"reflect"
	"testing"
)

func TestActivitySelection(t *testing.T) {
	start := []int{1, 3, 0, 5, 8, 5}
	end := []int{2, 4, 6, 7, 9, 9}
	got := ActivitySelection(start, end)
	want := []int{0, 1, 3, 4}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("ActivitySelection = %v, want %v", got, want)
	}
	if got := ActivitySelection(nil, nil); len(got) != 0 {
		t.Fatalf("空输入应返回空结果, got %v", got)
	}
}

func TestFractionalKnapsack(t *testing.T) {
	// 单位价值：60/10=6, 100/20=5, 120/30=4
	got := FractionalKnapsack([]int{10, 20, 30}, []int{60, 100, 120}, 50)
	want := 240.0
	if got != want {
		t.Fatalf("FractionalKnapsack = %v, want %v", got, want)
	}
	if got := FractionalKnapsack(nil, nil, 10); got != 0 {
		t.Fatalf("空输入 = %v, want 0", got)
	}
}

func TestJumpGame(t *testing.T) {
	cases := []struct {
		nums []int
		want bool
	}{
		{[]int{2, 3, 1, 1, 4}, true},
		{[]int{3, 2, 1, 0, 4}, false},
		{[]int{0}, true},
	}
	for _, tc := range cases {
		if got := JumpGame(tc.nums); got != tc.want {
			t.Fatalf("JumpGame(%v) = %v, want %v", tc.nums, got, tc.want)
		}
	}
}

func TestMinJumps(t *testing.T) {
	cases := []struct {
		nums []int
		want int
	}{
		{[]int{2, 3, 1, 1, 4}, 2},
		{[]int{2, 3, 0, 1, 4}, 2},
		{[]int{0}, 0},
		{[]int{1, 1, 1, 1}, 3},
		{[]int{0, 1}, -1},
	}
	for _, tc := range cases {
		if got := MinJumps(tc.nums); got != tc.want {
			t.Fatalf("MinJumps(%v) = %d, want %d", tc.nums, got, tc.want)
		}
	}
}

func TestFindContentChildren(t *testing.T) {
	if got := FindContentChildren([]int{1, 2, 3}, []int{1, 1}); got != 1 {
		t.Fatalf("FindContentChildren = %d, want 1", got)
	}
	if got := FindContentChildren([]int{1, 2}, []int{1, 2, 3}); got != 2 {
		t.Fatalf("FindContentChildren = %d, want 2", got)
	}
}
