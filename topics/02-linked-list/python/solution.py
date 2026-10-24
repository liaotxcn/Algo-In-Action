"""02 链表 —— Python 实现。"""

from __future__ import annotations


class ListNode:
    """单链表节点。"""

    def __init__(self, val: int = 0, next: "ListNode | None" = None) -> None:
        self.val = val
        self.next = next


def reverse_list(head: ListNode | None) -> ListNode | None:
    """反转单链表，时间复杂度 O(n)，空间复杂度 O(1)。"""
    prev = None
    while head is not None:
        nxt = head.next
        head.next = prev
        prev = head
        head = nxt
    return prev


def has_cycle(head: ListNode | None) -> bool:
    """快慢指针判断链表是否有环。"""
    slow = fast = head
    while fast is not None and fast.next is not None:
        slow = slow.next
        fast = fast.next.next
        if slow is fast:
            return True
    return False


def get_intersection_node(head_a: ListNode | None, head_b: ListNode | None) -> ListNode | None:
    """返回两个链表的第一个公共节点。"""
    p, q = head_a, head_b
    while p is not q:
        p = p.next if p is not None else head_b
        q = q.next if q is not None else head_a
    return p


def middle_node(head: ListNode | None) -> ListNode | None:
    """返回中间节点；偶数个节点时返回后一个。"""
    slow = fast = head
    while fast is not None and fast.next is not None:
        slow = slow.next
        fast = fast.next.next
    return slow


def is_palindrome(head: ListNode | None) -> bool:
    """判断链表是否为回文，反转后半段实现 O(1) 额外空间。"""
    if head is None or head.next is None:
        return True

    second_half = reverse_list(middle_node(head))
    while second_half is not None:
        if head.val != second_half.val:
            return False
        head = head.next
        second_half = second_half.next
    return True


def merge_two_lists(l1: ListNode | None, l2: ListNode | None) -> ListNode | None:
    """合并两个升序链表。"""
    dummy = ListNode()
    tail = dummy
    while l1 is not None and l2 is not None:
        if l1.val <= l2.val:
            tail.next, l1 = l1, l1.next
        else:
            tail.next, l2 = l2, l2.next
        tail = tail.next

    tail.next = l1 if l1 is not None else l2
    return dummy.next
