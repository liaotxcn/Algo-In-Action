// Package solution 提供单链表相关算法的 Go 实现。
package solution

// ListNode 是单链表节点。
type ListNode struct {
	Val  int
	Next *ListNode
}

// ReverseList 反转单链表并返回新的头节点。
// 时间复杂度 O(n)，空间复杂度 O(1)。
func ReverseList(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}
	return prev
}

// HasCycle 判断链表是否存在环（快慢指针）。
// 时间复杂度 O(n)，空间复杂度 O(1)。
func HasCycle(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

// GetIntersectionNode 返回两个链表的第一个公共节点，没有则返回 nil。
// 时间复杂度 O(m+n)，空间复杂度 O(1)。
func GetIntersectionNode(headA, headB *ListNode) *ListNode {
	// 两个指针走完各自链表后切换到另一条链表，最终会在交点或 nil 相遇
	p, q := headA, headB
	for p != q {
		if p != nil {
			p = p.Next
		} else {
			p = headB
		}
		if q != nil {
			q = q.Next
		} else {
			q = headA
		}
	}
	return p
}

// MiddleNode 返回链表的中间节点；偶数个节点时返回后一个中间节点。
// 时间复杂度 O(n)，空间复杂度 O(1)。
func MiddleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

// IsPalindrome 判断链表是否为回文。
// 实现会反转后半段链表，时间复杂度 O(n)，空间复杂度 O(1)。
func IsPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	mid := MiddleNode(head)
	secondHalf := ReverseList(mid)
	for secondHalf != nil {
		if head.Val != secondHalf.Val {
			return false
		}
		head = head.Next
		secondHalf = secondHalf.Next
	}
	return true
}

// MergeTwoLists 合并两个升序链表，返回新的升序链表（复用原节点）。
// 时间复杂度 O(m+n)，空间复杂度 O(1)。
func MergeTwoLists(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	tail := dummy

	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			tail.Next = l1
			l1 = l1.Next
		} else {
			tail.Next = l2
			l2 = l2.Next
		}
		tail = tail.Next
	}

	if l1 != nil {
		tail.Next = l1
	} else {
		tail.Next = l2
	}
	return dummy.Next
}
