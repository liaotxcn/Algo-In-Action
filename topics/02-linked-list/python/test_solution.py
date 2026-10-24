"""02 链表的 pytest 测试。"""

from solution_loader import load_solution

solution = load_solution(__file__)


def build_list(*values: int):
    """根据传入值构建链表。"""
    dummy = solution.ListNode()
    tail = dummy
    for value in values:
        tail.next = solution.ListNode(value)
        tail = tail.next
    return dummy.next


def list_values(head) -> list[int]:
    """把链表转换成列表，便于断言。"""
    values = []
    while head is not None:
        values.append(head.val)
        head = head.next
    return values


def test_reverse_list() -> None:
    assert list_values(solution.reverse_list(build_list(1, 2, 3, 4, 5))) == [5, 4, 3, 2, 1]
    assert solution.reverse_list(None) is None


def test_has_cycle() -> None:
    head = build_list(3, 2, 0, -4)
    tail = head
    while tail.next is not None:
        tail = tail.next
    tail.next = head.next
    assert solution.has_cycle(head) is True
    assert solution.has_cycle(build_list(1, 2, 3)) is False


def test_get_intersection_node() -> None:
    shared = build_list(8, 4, 5)
    head_a = solution.ListNode(4, solution.ListNode(1, shared))
    head_b = solution.ListNode(5, solution.ListNode(6, solution.ListNode(1, shared)))
    assert solution.get_intersection_node(head_a, head_b) is shared
    assert solution.get_intersection_node(build_list(1), build_list(2)) is None


def test_middle_node() -> None:
    assert list_values(solution.middle_node(build_list(1, 2, 3, 4, 5))) == [3, 4, 5]
    assert list_values(solution.middle_node(build_list(1, 2, 3, 4))) == [3, 4]
    assert solution.middle_node(None) is None


def test_is_palindrome() -> None:
    assert solution.is_palindrome(build_list(1, 2, 2, 1)) is True
    assert solution.is_palindrome(build_list(1, 2, 3)) is False
    assert solution.is_palindrome(None) is True


def test_merge_two_lists() -> None:
    merged = solution.merge_two_lists(build_list(1, 2, 4), build_list(1, 3, 4))
    assert list_values(merged) == [1, 1, 2, 3, 4, 4]
