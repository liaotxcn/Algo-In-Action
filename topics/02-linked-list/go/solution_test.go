package solution

import (
	"reflect"
	"testing"
)

// buildList 根据切片构建链表，返回头节点。
func buildList(values ...int) *ListNode {
	dummy := &ListNode{}
	tail := dummy
	for _, value := range values {
		tail.Next = &ListNode{Val: value}
		tail = tail.Next
	}
	return dummy.Next
}

// listValues 将链表转换为切片，便于断言。
func listValues(head *ListNode) []int {
	values := make([]int, 0)
	for node := head; node != nil; node = node.Next {
		values = append(values, node.Val)
	}
	return values
}

func TestReverseList(t *testing.T) {
	got := listValues(ReverseList(buildList(1, 2, 3, 4, 5)))
	want := []int{5, 4, 3, 2, 1}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("ReverseList = %v, want %v", got, want)
	}
	if ReverseList(nil) != nil {
		t.Fatal("空链表反转后应为 nil")
	}
}

func TestHasCycle(t *testing.T) {
	head := buildList(3, 2, 0, -4)
	tail := head
	for tail.Next != nil {
		tail = tail.Next
	}
	tail.Next = head.Next // 尾部连到第二个节点

	if !HasCycle(head) {
		t.Fatal("应检测到环")
	}
	if HasCycle(buildList(1, 2, 3)) {
		t.Fatal("无环链表应返回 false")
	}
}

func TestGetIntersectionNode(t *testing.T) {
	shared := buildList(8, 4, 5)
	headA := &ListNode{Val: 4, Next: &ListNode{Val: 1, Next: shared}}
	headB := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 1, Next: shared}}}

	if got := GetIntersectionNode(headA, headB); got != shared {
		t.Fatalf("GetIntersectionNode = %p, want %p", got, shared)
	}

	if got := GetIntersectionNode(buildList(1), buildList(2)); got != nil {
		t.Fatalf("无交点应返回 nil, got %v", listValues(got))
	}
}

func TestMiddleNode(t *testing.T) {
	cases := []struct {
		values []int
		want   []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{3, 4, 5}},
		{[]int{1, 2, 3, 4}, []int{3, 4}},
		{[]int{}, []int{}},
	}
	for _, tc := range cases {
		var head *ListNode
		if len(tc.values) > 0 {
			head = buildList(tc.values...)
		}
		if got := listValues(MiddleNode(head)); !reflect.DeepEqual(got, tc.want) {
			t.Fatalf("MiddleNode(%v) = %v, want %v", tc.values, got, tc.want)
		}
	}
}

func TestIsPalindrome(t *testing.T) {
	if !IsPalindrome(buildList(1, 2, 2, 1)) {
		t.Fatal("1->2->2->1 应为回文")
	}
	if IsPalindrome(buildList(1, 2, 3)) {
		t.Fatal("1->2->3 不应为回文")
	}
	if !IsPalindrome(nil) {
		t.Fatal("空链表应为回文")
	}
}

func TestMergeTwoLists(t *testing.T) {
	got := listValues(MergeTwoLists(buildList(1, 2, 4), buildList(1, 3, 4)))
	want := []int{1, 1, 2, 3, 4, 4}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("MergeTwoLists = %v, want %v", got, want)
	}
}
