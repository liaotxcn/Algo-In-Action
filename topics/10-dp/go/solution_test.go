package solution

import "testing"

func TestFib(t *testing.T) {
	want := []int{0, 1, 1, 2, 3, 5, 8, 13, 21}
	for n, expected := range want {
		if got := Fib(n); got != expected {
			t.Fatalf("Fib(%d) = %d, want %d", n, got, expected)
		}
	}
	if got := Fib(-1); got != 0 {
		t.Fatalf("Fib(-1) = %d, want 0", got)
	}
}

func TestClimbStairs(t *testing.T) {
	cases := map[int]int{1: 1, 2: 2, 3: 3, 4: 5, 5: 8}
	for n, want := range cases {
		if got := ClimbStairs(n); got != want {
			t.Fatalf("ClimbStairs(%d) = %d, want %d", n, got, want)
		}
	}
}

func TestKnapsack01(t *testing.T) {
	weights := []int{1, 3, 4}
	values := []int{15, 20, 30}
	if got := Knapsack01(weights, values, 4); got != 35 {
		t.Fatalf("Knapsack01 = %d, want 35", got)
	}
	if got := Knapsack01(weights, values, 0); got != 0 {
		t.Fatalf("容量为 0 时 = %d, want 0", got)
	}
}

func TestLIS(t *testing.T) {
	cases := []struct {
		nums []int
		want int
	}{
		{[]int{10, 9, 2, 5, 3, 7, 101, 18}, 4},
		{[]int{0, 1, 0, 3, 2, 3}, 4},
		{[]int{7, 7, 7}, 1},
		{[]int{}, 0},
	}
	for _, tc := range cases {
		if got := LIS(tc.nums); got != tc.want {
			t.Fatalf("LIS(%v) = %d, want %d", tc.nums, got, tc.want)
		}
	}
}

func TestEditDistance(t *testing.T) {
	cases := []struct {
		a, b string
		want int
	}{
		{"horse", "ros", 3},
		{"intention", "execution", 5},
		{"", "abc", 3},
		{"same", "same", 0},
	}
	for _, tc := range cases {
		if got := EditDistance(tc.a, tc.b); got != tc.want {
			t.Fatalf("EditDistance(%q, %q) = %d, want %d", tc.a, tc.b, got, tc.want)
		}
	}
}

func TestCoinChange(t *testing.T) {
	cases := []struct {
		coins  []int
		amount int
		want   int
	}{
		{[]int{1, 2, 5}, 11, 3},
		{[]int{2}, 3, -1},
		{[]int{1}, 0, 0},
		{[]int{1, 2, 5}, 0, 0},
	}
	for _, tc := range cases {
		if got := CoinChange(tc.coins, tc.amount); got != tc.want {
			t.Fatalf("CoinChange(%v, %d) = %d, want %d", tc.coins, tc.amount, got, tc.want)
		}
	}
}
