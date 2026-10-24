package solution

import "testing"

func TestBinarySearch(t *testing.T) {
	nums := []int{1, 3, 5, 7, 9, 11, 13, 15}
	cases := []struct {
		target int
		want   int
	}{
		{1, 0},
		{13, 6},
		{15, 7},
		{8, -1},
		{0, -1},
		{16, -1},
	}
	for _, tc := range cases {
		if got := BinarySearch(nums, tc.target); got != tc.want {
			t.Fatalf("BinarySearch(%d) = %d, want %d", tc.target, got, tc.want)
		}
	}
	if got := BinarySearch([]int{}, 1); got != -1 {
		t.Fatalf("空数组 BinarySearch = %d, want -1", got)
	}
}

func TestBounds(t *testing.T) {
	nums := []int{1, 2, 2, 2, 4, 6}
	cases := []struct {
		target int
		lower  int
		upper  int
	}{
		{2, 1, 4},
		{1, 0, 1},
		{6, 5, 6},
		{3, 4, 4},
		{0, 0, 0},
		{7, 6, 6},
	}
	for _, tc := range cases {
		if got := LowerBound(nums, tc.target); got != tc.lower {
			t.Fatalf("LowerBound(%d) = %d, want %d", tc.target, got, tc.lower)
		}
		if got := UpperBound(nums, tc.target); got != tc.upper {
			t.Fatalf("UpperBound(%d) = %d, want %d", tc.target, got, tc.upper)
		}
	}
}

func TestSearchRange(t *testing.T) {
	nums := []int{5, 7, 7, 7, 8, 10}
	if got := SearchRange(nums, 7); got != [2]int{1, 3} {
		t.Fatalf("SearchRange(7) = %v, want [1 3]", got)
	}
	if got := SearchRange(nums, 6); got != [2]int{-1, -1} {
		t.Fatalf("SearchRange(6) = %v, want [-1 -1]", got)
	}
}

func TestInterpolationSearch(t *testing.T) {
	nums := make([]int, 0, 100)
	for i := 0; i < 100; i++ {
		nums = append(nums, i*2)
	}
	if got := InterpolationSearch(nums, 84); got != 42 {
		t.Fatalf("InterpolationSearch(84) = %d, want 42", got)
	}
	if got := InterpolationSearch(nums, 85); got != -1 {
		t.Fatalf("InterpolationSearch(85) = %d, want -1", got)
	}
}
