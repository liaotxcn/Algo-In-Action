package solution

import (
	"reflect"
	"testing"
)

func TestIsValid(t *testing.T) {
	cases := []struct {
		input string
		want  bool
	}{
		{"()", true},
		{"()[]{}", true},
		{"(]", false},
		{"([)]", false},
		{"{[]}", true},
		{"", true},
		{"(", false},
	}
	for _, tc := range cases {
		if got := IsValid(tc.input); got != tc.want {
			t.Fatalf("IsValid(%q) = %v, want %v", tc.input, got, tc.want)
		}
	}
}

func TestMinStack(t *testing.T) {
	stack := NewMinStack()
	stack.Push(-2)
	stack.Push(0)
	stack.Push(-3)

	if got := stack.GetMin(); got != -3 {
		t.Fatalf("GetMin = %d, want -3", got)
	}
	if got := stack.Pop(); got != -3 {
		t.Fatalf("Pop = %d, want -3", got)
	}
	if got := stack.Top(); got != 0 {
		t.Fatalf("Top = %d, want 0", got)
	}
	if got := stack.GetMin(); got != -2 {
		t.Fatalf("GetMin = %d, want -2", got)
	}
}

func TestMyQueue(t *testing.T) {
	queue := NewMyQueue()
	if !queue.Empty() {
		t.Fatal("新队列应为空")
	}
	queue.Push(1)
	queue.Push(2)
	if got := queue.Peek(); got != 1 {
		t.Fatalf("Peek = %d, want 1", got)
	}
	if got := queue.Pop(); got != 1 {
		t.Fatalf("Pop = %d, want 1", got)
	}
	if queue.Empty() {
		t.Fatal("队列不应为空")
	}
	queue.Push(3)
	if got := queue.Pop(); got != 2 {
		t.Fatalf("Pop = %d, want 2", got)
	}
	if got := queue.Pop(); got != 3 {
		t.Fatalf("Pop = %d, want 3", got)
	}
}

func TestMaxSlidingWindow(t *testing.T) {
	cases := []struct {
		nums []int
		k    int
		want []int
	}{
		{[]int{1, 3, -1, -3, 5, 3, 6, 7}, 3, []int{3, 3, 5, 5, 6, 7}},
		{[]int{1}, 1, []int{1}},
		{[]int{9, 8, 7}, 5, []int{9}},
		{[]int{}, 3, []int{}},
	}
	for _, tc := range cases {
		if got := MaxSlidingWindow(tc.nums, tc.k); !reflect.DeepEqual(got, tc.want) {
			t.Fatalf("MaxSlidingWindow(%v, %d) = %v, want %v", tc.nums, tc.k, got, tc.want)
		}
	}
}
